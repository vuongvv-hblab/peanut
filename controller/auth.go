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

// Login godoc
// @Summary      Create an user
// @Description  Create an user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        body  body object  true  "Body"
// @Created      200  {object}  domain.Auth
// @Failure      400  {object}  domain.ErrorResponse
// @Failure      404  {object}  domain.ErrorResponse
// @Failure      500  {object}  domain.ErrorResponse
// @Router      /login [post]
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
