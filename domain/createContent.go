package domain

import (
	"mime/multipart"
)

type CreateContentReq struct {
	Name        string                `form:"name" binding:"required"`
	Thumbnail   *multipart.FileHeader `form:"thumbnail" binding:"required"`
	Content     *multipart.FileHeader `form:"content" binding:"required"`
	Description string                `form:"description" binding:"required"`
	Playtime    string                `form:"playtime" binding:"required"`
	Resolution  string                `form:"resolution" binding:"required"`
	Aspect      string                `form:"aspect" binding:"required"`
	Tag         bool                  `form:"tag" binding:"required"`
	Category    string                `form:"category" binding:"required"`
}
