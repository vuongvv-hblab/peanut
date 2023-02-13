package usecase

import (
	"context"
	"peanut/domain"
	"peanut/repository"
)

type BookUsecase interface {
	GetBooks(ctx context.Context, userId int) ([]domain.Book, error)
	GetBook(ctx context.Context, userId int, id int) (*domain.Book, error)
	CreateBook(ctx context.Context, u domain.CreateBookReq, userId int) error
	EditBook(ctx context.Context, u domain.CreateBookReq, userId int, id int) error
	DeleteBook(ctx context.Context, userId int, id int) error
}

type bookUsecase struct {
	BookRepo repository.BookRepo
}

func NewBookUsecase(r repository.BookRepo) BookUsecase {
	return &bookUsecase{
		BookRepo: r,
	}
}
func (uc *bookUsecase) GetBooks(ctx context.Context, userId int) (books []domain.Book, err error) {

	books, err = uc.BookRepo.GetBooks(ctx, userId)
	if err != nil {
		return nil, err
	}

	return books, nil
}
func (uc *bookUsecase) GetBook(ctx context.Context, userId int, id int) (book *domain.Book, err error) {

	book, err = uc.BookRepo.GetBook(ctx, userId, id)
	if err != nil {
		return nil, err
	}

	return book, nil
}
func (uc *bookUsecase) CreateBook(ctx context.Context, u domain.CreateBookReq, userId int) (err error) {

	_, err = uc.BookRepo.CreateBook(ctx, u, userId)
	if err != nil {
		return err
	}

	return
}
func (uc *bookUsecase) EditBook(ctx context.Context, u domain.CreateBookReq, userId int, id int) (err error) {

	_, err = uc.BookRepo.EditBook(ctx, u, userId, id)
	if err != nil {
		return err
	}

	return
}
func (uc *bookUsecase) DeleteBook(ctx context.Context, userId int, id int) (err error) {

	_, err = uc.BookRepo.DeleteBook(ctx, userId, id)
	if err != nil {
		return err
	}

	return
}
