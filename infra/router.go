package infra

import (
	"fmt"
	"net/http"

	"peanut/config"
	"peanut/controller"
	"time"

	"peanut/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	_ "peanut/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	Router *gin.Engine
	Store  *gorm.DB
}

func SetupServer(s *gorm.DB) Server {
	// Init router
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// Custom middleware
	r.Use(middleware.HandleError)
	r.NoRoute(middleware.HandleNoRoute)
	r.NoMethod(middleware.HandleNoMethod)

	// CORS config
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PATCH", "PUT", "DELETE"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	// Config route
	r.GET("", func(context *gin.Context) {
		fmt.Println(123)
		context.JSON(200, "abcd")
	})
	r.POST("/login", controller.NewAuthController(s).Login)
	v1 := r.Group("api/v1")
	{
		userCtrl := controller.NewUserController(s)
		users := v1.Group("/users")
		{
			users.GET("", userCtrl.GetUsers)
			users.POST("", userCtrl.CreateUser)
			users.GET("/:id", userCtrl.GetUser)
			// users.PATCH("/:id", userCtrl.UpdateUser)
			// users.DELETE("/:id", userCtrl.DeleteUserByID)
		}
		bookCtrl := controller.NewBookController(s)
		books := v1.Group("/books").Use(middleware.JWTAuthMiddleware())
		{
			books.GET("", bookCtrl.GetBooks)
			books.POST("", bookCtrl.CreateBook)
			books.GET("/:id", bookCtrl.GetBook)
			books.PUT("/:id", bookCtrl.EditBook)
			books.DELETE("/:id", bookCtrl.DeleteBook)
		}
		contentCtrl := controller.NewContentController(s)
		contents := v1.Group("/contents").Use(middleware.JWTAuthMiddleware())
		{
			contents.GET("", contentCtrl.GetContents)
			contents.POST("", contentCtrl.CreateContent)
		}
	}

	// health check
	r.GET("api/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	if config.IsDevelopment() {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return Server{
		Store:  s,
		Router: r,
	}
}
