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

// GetBooks godoc
//	@Summary		Get books
//	@Description	Get books
//	@Tags			book
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]domain.Book
//	@Failure		400	{object}	domain.ErrorResponse
//	@Failure		404	{object}	domain.ErrorResponse
//	@Failure		500	{object}	domain.ErrorResponse
//	@Security		Bearer
//	@Router			/api/v1/books [get]
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

// GetBook godoc
//	@Summary		Get book
//	@Description	Get book
//	@Tags			book
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Book ID"
//	@Success		200	{object}	domain.Book
//	@Failure		400	{object}	domain.ErrorResponse
//	@Failure		404	{object}	domain.ErrorResponse
//	@Failure		500	{object}	domain.ErrorResponse
//	@Security		Bearer
//	@Router			/api/v1/books/{id} [get]
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

// CreateBook godoc
//	@Summary		Create an book
//	@Description	Create an book
//	@Tags			book
//	@Accept			json
//	@Produce		json
//	@Param			body	body		object	true	"Body"
//	@Success		200		{object}	domain.Book
//	@Failure		400		{object}	domain.ErrorResponse
//	@Failure		404		{object}	domain.ErrorResponse
//	@Failure		500		{object}	domain.ErrorResponse
//	@Security		Bearer
//	@Router			/api/v1/books [post]
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

// EditBook godoc
//	@Summary		Edit an book
//	@Description	Edit an book
//	@Tags			book
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int		true	"Book ID"
//	@Param			body	body		object	true	"Body"
//	@Success		200		{object}	domain.Book
//	@Failure		400		{object}	domain.ErrorResponse
//	@Failure		404		{object}	domain.ErrorResponse
//	@Failure		500		{object}	domain.ErrorResponse
//	@Security		Bearer
//	@Router			/api/v1/books/{id} [put]
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

// DeleteBook godoc
//	@Summary		Delete an book
//	@Description	Delete an book
//	@Tags			book
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Book ID"
//	@Success		200	{object}	domain.Book
//	@Failure		400	{object}	domain.ErrorResponse
//	@Failure		404	{object}	domain.ErrorResponse
//	@Failure		500	{object}	domain.ErrorResponse
//	@Security		Bearer
//	@Router			/api/v1/books/{id} [Delete]
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
