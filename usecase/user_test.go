package usecase_test

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"peanut/pkg/crypto"

	"peanut/domain"
)

var _ = Describe("User", func() {
	var u domain.CreateUserReq
	var user domain.User
	BeforeEach(func() {
		p := crypto.HashString("12345678")
		u = domain.CreateUserReq{
			Username: "vuong",
			Email:    "vuong@example.com",
			Password: p,
		}

	})

	Describe("API Create", func() {
		Context("with existed user", func() {
			It("should be error", func() {
				// TODO: fill in your test in this case
			})
		})
		Context("with new user", func() {
			It("should be success", func() {
				// prepare
				userRepo.EXPECT().CreateUser(ctx, u).Return(&user, nil)
				// do
				err := userUc.CreateUser(ctx, u)
				fmt.Println(err, user)
				// check
				Expect(err).To(BeNil())
			})
		})
		Context("with database error response", func() {
			It("should be err", func() {
				// prepare
				userRepo.EXPECT().CreateUser(ctx, u).
					Return(nil, fmt.Errorf("database error"))
				// do
				err := userUc.CreateUser(ctx, u)
				// check
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
