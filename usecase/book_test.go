package usecase_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"peanut/domain"
)

var _ = Describe("Books", func() {
	var u, createdB domain.CreateBookReq
	userId := 10
	BeforeEach(func() {
		u = domain.CreateBookReq{
			Name: "book 1",
		}
		createdB = domain.CreateBookReq{}
	})

	Describe("API Create", func() {
		Context("with existed book", func() {
			It("should be error", func() {
				// TODO: fill in your test in this case
			})
		})
		Context("with new book", func() {
			It("should be success", func() {
				// prepare
				bookRepo.EXPECT().CreateBook(ctx, u, userId).Return(&u, nil)
				// do
				err := bookUc.CreateBook(ctx, u, userId)
				// check
				Expect(err).To(BeNil())
			})
		})
		Context("with database error response", func() {
			It("should be err", func() {
				// prepare
				bookRepo.EXPECT().CreateBook(ctx, u, userId).
					Return(nil, fmt.Errorf("database error"))
				// do
				err := bookUc.CreateBook(ctx, u, userId)
				// check
				Expect(err).NotTo(BeNil())
			})
		})
	})

	Describe("API Get user", func() {
		Context("with existed id", func() {
			It("should be return user", func() {
				// Do something
				Expect(u).NotTo(Equal(createdB))
			})
		})
	})
})
