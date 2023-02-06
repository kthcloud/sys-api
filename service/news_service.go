package service

import (
	"landing-api/models/dto"
	"landing-api/models/news"
)

func CreateNews(newsCreate *dto.NewsCreate) (*news.News, error) {
	return news.Create(newsCreate)
}

func GetNewsByID(id string) (*news.News, error) {
	return news.Get(id)
}

func GetAllNews() ([]news.News, error) {
	return news.GetAll()
}

func DeleteNewsByID(id string) error {
	return news.Delete(id)
}
