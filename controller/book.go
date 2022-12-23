package controller

import (
	"github.com/gin-gonic/gin"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"net/http"
	"peanut/domain"
	"peanut/pkg/jwt"
	"peanut/pkg/response"
	"peanut/repository"
	"peanut/usecase"
	"strconv"
)

type BookController struct {
	Usecase usecase.BookUsecase
}

func NewBookController(db *gorm.DB) *BookController {
	return &BookController{
		Usecase: usecase.NewBookUsecase(repository.NewBookRepo(db)),
	}
}
func (c *BookController) GetBooks(ctx *gin.Context) {
	token, _ := jwt.GetToken(ctx)
	claims := token.Claims.(jwt2.MapClaims)

	userId := int(claims["id"].(float64))

	books, err := c.Usecase.GetBooks(ctx, userId)
	if checkError(ctx, err) {
		return
	} else {
		response.OK(ctx, books)
	}
}
func (c *BookController) GetBook(ctx *gin.Context) {
	token, _ := jwt.GetToken(ctx)
	claims := token.Claims.(jwt2.MapClaims)

	userId := int(claims["id"].(float64))

	id, err2 := strconv.Atoi(ctx.Param("id"))

	if err2 != nil {
		return
	}
	book, err := c.Usecase.GetBook(ctx, userId, id)
	if checkError(ctx, err) {
		return
	} else {
		response.OK(ctx, book)
	}
}
func (c *BookController) CreateBook(ctx *gin.Context) {
	token, _ := jwt.GetToken(ctx)
	claims := token.Claims.(jwt2.MapClaims)

	userId := int(claims["id"].(float64))
	book := domain.CreateBookReq{}
	if !bindJSON(ctx, &book) {
		return
	}

	err := c.Usecase.CreateBook(ctx, book, userId)
	if checkError(ctx, err) {
		return
	}

	response.WithStatusCode(ctx, http.StatusCreated, nil)
}
func (c *BookController) EditBook(ctx *gin.Context) {
	token, _ := jwt.GetToken(ctx)
	claims := token.Claims.(jwt2.MapClaims)

	userId := int(claims["id"].(float64))
	id, err2 := strconv.Atoi(ctx.Param("id"))

	if err2 != nil {
		return
	}
	book := domain.CreateBookReq{}
	if !bindJSON(ctx, &book) {
		return
	}

	err := c.Usecase.EditBook(ctx, book, userId, id)
	if checkError(ctx, err) {
		return
	}

	response.OK(ctx, nil)
}
func (c *BookController) DeleteBook(ctx *gin.Context) {
	token, _ := jwt.GetToken(ctx)
	claims := token.Claims.(jwt2.MapClaims)

	userId := int(claims["id"].(float64))
	id, err2 := strconv.Atoi(ctx.Param("id"))

	if err2 != nil {
		return
	}

	err := c.Usecase.DeleteBook(ctx, userId, id)
	if checkError(ctx, err) {
		return
	}

	response.OK(ctx, nil)
}
