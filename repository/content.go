package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"peanut/domain"
	"time"
)

type ContentRepo interface {
	GetContents(ctx context.Context, userId int) ([]domain.Content, error)
	CreateContent(ctx context.Context, u domain.CreateContentReq, userId int, thumbnail string, contentPath string) (*domain.Content, error)
}
type contentRepo struct {
	DB *gorm.DB
}

func NewContentRepo(db *gorm.DB) ContentRepo {
	return &contentRepo{DB: db}
}

func (r *contentRepo) GetContents(ctx context.Context, userId int) (contents []domain.Content, err error) {
	result := r.DB.Where("user_id= ?", userId).Find(&contents)

	if result.Error != nil {
		err = fmt.Errorf("[repo.Book.GetBooks] failed: %w", result.Error)
		return nil, err
	}
	return contents, nil
}

func (r *contentRepo) CreateContent(ctx context.Context, c domain.CreateContentReq, userId int, thumbnail string, contentPath string) (content *domain.Content, err error) {
	date, _ := time.Parse("2006-01-02", c.Playtime)
	content = &domain.Content{
		Name:        c.Name,
		Thumbnail:   thumbnail,
		Content:     contentPath,
		Description: c.Description,
		Playtime:    date,
		Resolution:  c.Resolution,
		Aspect:      c.Aspect,
		Tag:         c.Tag,
		Category:    c.Category,
		UserId:      userId,
	}
	result := r.DB.Create(content)

	if result.Error != nil {
		err = fmt.Errorf("[repo.Book.CreateBook] failed: %w", result.Error)
		return nil, err
	}
	return
}
