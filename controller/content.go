package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"net/http"
	"peanut/config"
	"peanut/domain"
	"peanut/pkg/jwt"
	"peanut/pkg/response"
	"peanut/repository"
	"peanut/usecase"
)

type ContentController struct {
	Usecase usecase.ContentUsecase
}

func NewContentController(db *gorm.DB) *ContentController {
	fmt.Println(config.IsDevelopment())
	return &ContentController{
		Usecase: usecase.NewContentUsecase(repository.NewContentRepo(db)),
	}
}

// GetContents godoc
//
//	@Summary		Get contents
//	@Description	Get contents
//	@Tags			content
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]domain.Content
//	@Failure		400	{object}	domain.ErrorResponse
//	@Failure		404	{object}	domain.ErrorResponse
//	@Failure		500	{object}	domain.ErrorResponse
//	@Security		Bearer
//	@Router			/api/v1/contents [get]
func (c *ContentController) GetContents(ctx *gin.Context) {
	token, _ := jwt.GetToken(ctx)
	claims := token.Claims.(jwt2.MapClaims)

	userId := int(claims["id"].(float64))

	contents, err := c.Usecase.GetContents(ctx, userId)
	if checkError(ctx, err) {
		return
	} else {
		response.OK(ctx, contents)
	}
}

// CreateContent godoc
//
//	@Summary		Create an content
//	@Description	Create an content
//	@Tags			content
//	@Accept			json
//	@Produce		json
//	@Param			name		formData	string	true	"name"
//	@Param			thumbnail	formData	file	true	"thumbnail"
//	@Param			content		formData	file	true	"content"
//	@Param			description	formData	string	true	"description"
//	@Param			playtime	formData	string	true	"playtime"
//	@Param			resolution	formData	string	true	"resolution"
//	@Param			aspect		formData	string	true	"aspect"
//	@Param			tag			formData	bool	true	"tag"
//	@Param			category	formData	string	true	"category"
//	@Created		201  {object}  domain.Content
//	@Failure		400	{object}	domain.ErrorResponse
//	@Failure		404	{object}	domain.ErrorResponse
//	@Failure		500	{object}	domain.ErrorResponse
//	@Security		Bearer
//	@Router			/api/v1/contents [post]
func (c *ContentController) CreateContent(ctx *gin.Context) {
	token, _ := jwt.GetToken(ctx)
	claims := token.Claims.(jwt2.MapClaims)

	userId := int(claims["id"].(float64))
	content := domain.CreateContentReq{}
	if !bindForm(ctx, &content) {
		return
	}
	contentPath, _ := saveUploadedFile(ctx, "content", config.ContentPath())
	thumbnailPath, _ := saveUploadedFile(ctx, "thumbnail", config.ThumbnailPath())

	err := c.Usecase.CreateContent(ctx, content, userId, contentPath, thumbnailPath)
	if checkError(ctx, err) {
		return
	}

	response.WithStatusCode(ctx, http.StatusCreated, nil)
}
