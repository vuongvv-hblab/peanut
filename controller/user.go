package controller

import (
	"fmt"
	"net/http"
	"peanut/config"
	"peanut/domain"
	"peanut/pkg/response"
	"peanut/repository"
	"peanut/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	Usecase usecase.UserUsecase
}

func NewUserController(db *gorm.DB) *UserController {
	fmt.Println(config.IsDevelopment())
	return &UserController{
		Usecase: usecase.NewUserUsecase(repository.NewUserRepo(db)),
	}
}

func (c *UserController) GetUsers(ctx *gin.Context) {

}

// GetUser godoc
// @Summary      Get user
// @Description  Get user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  domain.User
// @Failure      400  {object}  domain.ErrorResponse
// @Failure      404  {object}  domain.ErrorResponse
// @Failure      500  {object}  domain.ErrorResponse
// @Router       /api/v1/users/{id} [get]
func (c *UserController) GetUser(ctx *gin.Context) {

}

// CreateUser godoc
// @Summary      Create an user
// @Description  Create an user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        body  body object  true  "Body"
// @Created      201  {object}  domain.User
// @Failure      400  {object}  domain.ErrorResponse
// @Failure      404  {object}  domain.ErrorResponse
// @Failure      500  {object}  domain.ErrorResponse
// @Router       /api/v1/users [post]
func (c *UserController) CreateUser(ctx *gin.Context) {
	user := domain.CreateUserReq{}
	if !bindJSON(ctx, &user) {
		return
	}

	err := c.Usecase.CreateUser(ctx, user)
	if checkError(ctx, err) {
		return
	}

	response.WithStatusCode(ctx, http.StatusCreated, nil)
}
