package usecase

import (
	"context"
	"peanut/domain"
	"peanut/pkg/apierrors"
	"peanut/repository"
	"time"
)

type ContentUsecase interface {
	GetContents(ctx context.Context, userId int) ([]domain.Content, error)
	CreateContent(ctx context.Context, c domain.CreateContentReq, userId int, contentPath string, thumbnailPath string) error
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

func (cc *contentUsecase) CreateContent(ctx context.Context, c domain.CreateContentReq, userId int, contentPath string, thumbnailPath string) (err error) {
	//var maxBytes int64 = 1024 * 1024 * 5 // 5MB
	//
	//var w http.ResponseWriter = ctx.Writer
	//ctx.Request.Body = http.MaxBytesReader(w, ctx.Request.Body, maxBytes)
	//if err := ctx.Request.ParseMultipartForm(maxBytes); err != nil {
	//	http.Error(w, "The uploaded file is too big. Please choose an file that's less than 1MB in size", http.StatusBadRequest)
	//	return
	//}
	date, errDate := time.Parse("2006-01-02 15:04:05", c.Playtime)
	if errDate != nil {
		err = apierrors.NewErrorf(apierrors.InternalError, errDate.Error())
		return
	}
	content := domain.Content{
		Name:        c.Name,
		Thumbnail:   thumbnailPath,
		Content:     contentPath,
		Description: c.Description,
		Playtime:    date,
		Resolution:  c.Resolution,
		Aspect:      c.Aspect,
		Tag:         c.Tag,
		Category:    c.Category,
		UserId:      userId,
	}
	_, err = cc.ContentRepo.CreateContent(content)
	if err != nil {
		return err
	}

	return
}
