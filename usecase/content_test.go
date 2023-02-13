package usecase_test

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"peanut/domain"
	"time"
)

var _ = Describe("Content", func() {
	var c domain.CreateContentReq
	var content domain.Content
	var thumbnailPath, contentPath string
	userId := 10
	date, _ := time.Parse("2006-01-02 15:04:05", "2022-12-15 16:07:54")
	BeforeEach(func() {
		c = domain.CreateContentReq{
			Name:        "1",
			Description: "1",
			Playtime:    "2022-12-15 16:07:54",
			Resolution:  "1",
			Aspect:      "1",
			Tag:         true,
			Category:    "1",
		}
		contentPath = "public/content/20230109115750.jpg"
		thumbnailPath = "public/thumbnail/20230109115750.jpg"
		content = domain.Content{
			Thumbnail:   "public/thumbnail/20230109115750.jpg",
			Content:     "public/content/20230109115750.jpg",
			Name:        "1",
			Description: "1",
			Playtime:    date,
			Resolution:  "1",
			Aspect:      "1",
			Tag:         true,
			Category:    "1",
			UserId:      userId,
		}
	})

	Describe("API Create", func() {
		Context("with new content", func() {
			It("should be success", func() {
				// prepare
				contentRepo.EXPECT().CreateContent(content).Return(&content, nil)
				// do
				err := contentUc.CreateContent(ctx, c, userId, contentPath, thumbnailPath)
				// check
				Expect(err).To(BeNil())
			})
		})
		Context("with database error response", func() {
			It("should be err", func() {
				// prepare
				contentRepo.EXPECT().CreateContent(content).
					Return(nil, fmt.Errorf("database error"))
				// do
				err := contentUc.CreateContent(ctx, c, userId, contentPath, thumbnailPath)
				// check
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
