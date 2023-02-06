package dto

import "mime/multipart"

type NewsCreate struct {
	Title       string                `form:"title"`
	Description string                `form:"description"`
	Image       *multipart.FileHeader `form:"image"`
}
