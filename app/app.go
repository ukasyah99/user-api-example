package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gitlab.com/ukasyah99/user-api/handler"
	"gitlab.com/ukasyah99/user-api/middleware"
	"os"
)

type App struct {
	DB *sqlx.DB
	Router *gin.Engine
}

func (a *App) Run() {
	godotenv.Load()

	a.DB = sqlx.MustConnect("mysql", os.Getenv("MYSQL_URI"))
	a.Router = gin.Default()

	a.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	a.Router.POST("/auth/login", a.handle(handler.AuthLogin))
	a.Router.POST("/auth/refresh", a.handle(handler.AuthRefresh))

	a.Router.Use(middleware.Authorize(a.DB))
	a.Router.POST("/auth/logout", a.handle(handler.AuthLogout))
	a.Router.GET("/auth/me", a.handle(handler.AuthMe))
	a.Router.GET("/users", a.handle(handler.UserFindAll))
	a.Router.GET("/users/:id", a.handle(handler.UserFindOne))
	a.Router.POST("/users", a.handle(handler.UserCreate))
	a.Router.PUT("/users/:id", a.handle(handler.UserUpdate))
	a.Router.DELETE("/users/:id", a.handle(handler.UserDelete))
	
	a.Router.Run()
}

func (a *App) handle(handler handler.HandlerFunction) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(a.DB, c)
	}
}
