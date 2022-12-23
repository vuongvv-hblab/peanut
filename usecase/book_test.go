package usecase_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"peanut/domain"
)

var _ = Describe("Books", func() {
	var b domain.CreateBookReq
	var book domain.Book
	userId, bookId := 10, 6

	BeforeEach(func() {
		b = domain.CreateBookReq{
			Name: "book 1",
		}
		//createdB = domain.CreateBookReq{}
	})

	Describe("API Create", func() {
		Context("with new book", func() {
			It("should be success", func() {
				// prepare
				bookRepo.EXPECT().CreateBook(ctx, b, userId).Return(&book, nil)
				// do
				err := bookUc.CreateBook(ctx, b, userId)
				// check
				Expect(err).To(BeNil())
			})
		})
		Context("with database error response", func() {
			It("should be err", func() {
				// prepare
				bookRepo.EXPECT().CreateBook(ctx, b, userId).
					Return(nil, fmt.Errorf("database error"))
				// do
				err := bookUc.CreateBook(ctx, b, userId)
				// check
				Expect(err).NotTo(BeNil())
			})
		})
	})
	Describe("API edit", func() {
		Context("with existed book", func() {
			It("should be success", func() {
				// prepare
				bookRepo.EXPECT().EditBook(ctx, b, userId, bookId).Return(&book, nil)
				// do
				err := bookUc.EditBook(ctx, b, userId, bookId)
				// check
				Expect(err).To(BeNil())
			})
		})
		Context("with database error response", func() {
			It("should be err", func() {
				// prepare
				bookRepo.EXPECT().EditBook(ctx, b, userId, bookId).
					Return(nil, fmt.Errorf("database error"))
				// do
				err := bookUc.EditBook(ctx, b, userId, bookId)
				// check
				Expect(err).NotTo(BeNil())
			})
		})
	})
	Describe("API Get", func() {
		Context("with existed id", func() {
			It("should be return book", func() {
				bookRepo.EXPECT().GetBook(ctx, userId, bookId).Return(&book, nil)

				_, err := bookUc.GetBook(ctx, userId, bookId)
				// Do something
				Expect(err).To(BeNil())
			})
		})
	})
	Describe("API delete", func() {
		Context("with existed id", func() {
			It("should be return book", func() {
				bookRepo.EXPECT().DeleteBook(ctx, userId, bookId).Return(&book, nil)

				err := bookUc.DeleteBook(ctx, userId, bookId)
				// Do something
				Expect(err).To(BeNil())
			})
		})
	})
})
