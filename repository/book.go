package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"peanut/domain"
)

type BookRepo interface {
	GetBooks(ctx context.Context, userId int) ([]domain.Book, error)
	GetBook(ctx context.Context, userId int, id int) (*domain.Book, error)
	CreateBook(ctx context.Context, u domain.CreateBookReq, userId int) (*domain.Book, error)
	EditBook(ctx context.Context, u domain.CreateBookReq, userId int, id int) (*domain.Book, error)
	DeleteBook(ctx context.Context, userId int, id int) (*domain.Book, error)
}

type bookRepo struct {
	DB *gorm.DB
}

func NewBookRepo(db *gorm.DB) BookRepo {
	return &bookRepo{DB: db}
}
func (r *bookRepo) GetBooks(ctx context.Context, userId int) (books []domain.Book, err error) {
	result := r.DB.Where("user_id= ?", userId).Find(&books)

	if result.Error != nil {
		err = fmt.Errorf("[repo.Book.GetBooks] failed: %w", result.Error)
		return nil, err
	}
	return books, nil
}
func (r *bookRepo) GetBook(ctx context.Context, userId int, id int) (book *domain.Book, err error) {
	result := r.DB.Where("user_id= ? and id=?", userId, id).First(&book)

	if result.Error != nil {
		err = fmt.Errorf("[repo.Book.GetBook] failed: %w", result.Error)
		return nil, err
	}
	return book, nil
}
func (r *bookRepo) CreateBook(ctx context.Context, u domain.CreateBookReq, userId int) (book *domain.Book, err error) {
	book = &domain.Book{
		Name:   u.Name,
		UserId: userId,
	}
	result := r.DB.Create(book)

	if result.Error != nil {
		err = fmt.Errorf("[repo.Book.CreateBook] failed: %w", result.Error)
		return nil, err
	}
	return
}
func (r *bookRepo) EditBook(ctx context.Context, u domain.CreateBookReq, userId int, id int) (book *domain.Book, err error) {
	book = &domain.Book{
		Name: u.Name,
	}
	result := r.DB.Where("user_id= ? and id=?", userId, id).Updates(book)
	if book.ID == 0 {
		err = fmt.Errorf("[repo.Book.EditBook] failed: not exist book")
		return nil, err
	}
	if result.Error != nil {
		err = fmt.Errorf("[repo.Book.EditBook] failed: %w", result.Error)
		return nil, err
	}
	return
}
func (r *bookRepo) DeleteBook(ctx context.Context, userId int, id int) (book *domain.Book, err error) {
	book = &domain.Book{}
	result := r.DB.Where("user_id= ? and id=?", userId, id).Delete(book)

	if result.Error != nil {
		err = fmt.Errorf("[repo.Book.GetBook] failed: %w", result.Error)
		return nil, err
	}
	return
}
