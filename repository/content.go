package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"peanut/domain"
)

type ContentRepo interface {
	GetContents(ctx context.Context, userId int) ([]domain.Content, error)
	CreateContent(c domain.Content) (*domain.Content, error)
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

func (r *contentRepo) CreateContent(c domain.Content) (content *domain.Content, err error) {
	result := r.DB.Create(&c)

	if result.Error != nil {
		err = fmt.Errorf("[repo.Content.CreateContent] failed: %w", result.Error)
		return nil, err
	}
	return
}
