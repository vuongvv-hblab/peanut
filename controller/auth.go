package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"peanut/domain"
	"peanut/pkg/response"
	"peanut/usecase"
)

type AuthController struct {
	Usecase usecase.AuthUsecase
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		Usecase: usecase.NewAuthUsecase(db),
	}
}

func (c *AuthController) Login(ctx *gin.Context) {
	user := domain.Auth{}

	if !bindJSON(ctx, &user) {
		return
	}

	jwt, _ := c.Usecase.Login(ctx, user)
	if jwt != nil {
		response.OK(ctx, jwt)
	} else {
		response.WithStatusCode(ctx, http.StatusUnauthorized, nil)
	}
}
