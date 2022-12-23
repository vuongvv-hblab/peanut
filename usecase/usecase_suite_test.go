package usecase_test

import (
	"context"
	"testing"

	"peanut/repository/mock"
	"peanut/usecase"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ctx context.Context
var db *gorm.DB
var bookRepo *mock.MockBookRepo
var bookUc usecase.BookUsecase
var userRepo *mock.MockUserRepo
var userUc usecase.UserUsecase

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Books Usecase Suite")
}

var _ = BeforeSuite(func() {
	sqlDB, smock, _ := sqlmock.New()

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	Expect(err).To(BeNil())
	Expect(smock).NotTo(BeNil())
	Expect(db).NotTo(BeNil())

	ctrl := gomock.NewController(GinkgoT())
	defer ctrl.Finish()

	userRepo = mock.NewMockUserRepo(ctrl)
	userUc = usecase.NewUserUsecase(userRepo)

	bookRepo = mock.NewMockBookRepo(ctrl)
	bookUc = usecase.NewBookUsecase(bookRepo)
})
