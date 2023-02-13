package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"peanut/domain"
)

type UserRepo interface {
	GetUsers(ctx context.Context) ([]domain.User, error)
	GetUser(ctx context.Context, id int) (*domain.User, error)
	GetUserByGmail(ctx context.Context, email string) (*domain.User, error)
	CreateUser(ctx context.Context, u domain.CreateUserReq) (*domain.User, error)
}

type userRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{DB: db}
}

func (r *userRepo) GetUsers(ctx context.Context) (users []domain.User, err error) {
	return
}

func (r *userRepo) GetUser(ctx context.Context, id int) (user *domain.User, err error) {
	return
}
func (r *userRepo) GetUserByGmail(ctx context.Context, email string) (user *domain.User, err error) {
	user = &domain.User{}

	result := r.DB.Where("email = ?", email).First(user)

	if result.Error != nil {
		err = fmt.Errorf("[repo.auth.Login] failed: %w", result.Error)
		return nil, err
	}
	return user, nil
}

func (r *userRepo) CreateUser(ctx context.Context, u domain.CreateUserReq) (user *domain.User, err error) {
	user = &domain.User{
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}
	result := r.DB.Create(user)

	if result.Error != nil {
		err = fmt.Errorf("[repo.User.CreateUser] failed: %w", result.Error)
		return nil, err
	}
	return
}
