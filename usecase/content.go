package usecase

import (
	"context"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"peanut/domain"
	"peanut/repository"
	"time"
)

type ContentUsecase interface {
	GetContents(ctx context.Context, userId int) ([]domain.Content, error)
	CreateContent(ctx *gin.Context, c domain.CreateContentReq, userId int) error
}

type contentUsecase struct {
	ContentRepo repository.ContentRepo
}

func NewContentUsecase(r repository.ContentRepo) ContentUsecase {
	return &contentUsecase{
		ContentRepo: r,
	}
}

func (cc *contentUsecase) GetContents(ctx context.Context, userId int) (contents []domain.Content, err error) {
	contents, err = cc.ContentRepo.GetContents(ctx, userId)
	if err != nil {
		return nil, err
	}

	return contents, nil
}

func (cc *contentUsecase) CreateContent(ctx *gin.Context, c domain.CreateContentReq, userId int) (err error) {
	// Save uploaded file
	file, _ := ctx.FormFile("content")
	dst := time.Now().Format("20060102150405") + filepath.Ext(file.Filename)
	contentPath := "public/content/" + dst
	err = ctx.SaveUploadedFile(file, contentPath)
	if err != nil {
		return err
	}

	// Save upload thumbnail
	thumbnail, _ := ctx.FormFile("thumbnail")
	thumbnailDts := time.Now().Format("20060102150405") + filepath.Ext(thumbnail.Filename)
	thumbnailPath := "public/thumbnail/" + thumbnailDts
	err = ctx.SaveUploadedFile(file, thumbnailPath)
	if err != nil {
		return err
	}

	_, err = cc.ContentRepo.CreateContent(ctx, c, userId, thumbnailPath, contentPath)
	if err != nil {
		return err
	}

	return
}
