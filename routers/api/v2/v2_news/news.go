package v2_news

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"landing-api/models/dto"
	"landing-api/pkg/app"
	"landing-api/pkg/status_codes"
	"landing-api/pkg/validator"
	"landing-api/service"
	"net/http"
)

func Get(c *gin.Context) {
	context := app.NewContext(c)

	rules := validator.MapData{
		"id": []string{"required", "uuid_v4"},
	}

	validationErrors := context.ValidateParams(&rules)

	if len(validationErrors) > 0 {
		context.ResponseValidationError(validationErrors)
		return
	}
	newsID := context.GinContext.Param("id")

	news, err := service.GetNewsByID(newsID)
	if err != nil {
		context.ErrorResponse(http.StatusInternalServerError, status_codes.Error, fmt.Sprintf("%s", err))
	}

	if news == nil {
		context.NotFound()
		return
	}

	dtoNews := dto.NewsRead{
		ID:          news.ID,
		Title:       news.Title,
		Description: news.Description,
		Image:       news.Image,
		PostedAt:    news.PostedAt,
	}

	context.JSONResponse(200, dtoNews)
}

func GetAll(c *gin.Context) {
	context := app.NewContext(c)

	allNews, err := service.GetAllNews()
	if err != nil {
		context.ErrorResponse(http.StatusInternalServerError, status_codes.Error, fmt.Sprintf("%s", err))
	}

	if allNews == nil {
		context.JSONResponse(200, make([]interface{}, 0))
		return
	}

	dtoNews := make([]dto.NewsRead, len(allNews))
	for idx, news := range allNews {
		dtoNews[idx] = dto.NewsRead{
			ID:          news.ID,
			Title:       news.Title,
			Description: news.Description,
			Image:       news.Image,
			PostedAt:    news.PostedAt,
		}
	}

	context.JSONResponse(200, dtoNews)
}

func Create(c *gin.Context) {
	context := app.NewContext(c)

	var newsCreate dto.NewsCreate
	err := context.GinContext.Bind(&newsCreate)
	if err != nil {
		context.ErrorResponse(http.StatusInternalServerError, status_codes.Error, fmt.Sprintf("%s", err))
		return
	}

	if newsCreate.Title == "" {
		context.ResponseValidationError(map[string][]string{
			"title": {
				"Title is required",
			},
		})
		return
	}

	news, err := service.CreateNews(&newsCreate)
	if err != nil {
		context.ErrorResponse(http.StatusInternalServerError, status_codes.Error, fmt.Sprintf("%s", err))
		return
	}

	created := dto.NewsRead{
		ID:          news.ID,
		Title:       news.Title,
		Description: news.Description,
		Image:       news.Image,
		PostedAt:    news.PostedAt,
	}

	context.JSONResponse(200, created)
}
