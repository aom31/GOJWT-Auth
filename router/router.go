package router

import (
	"net/http"
	"project-for-portfolioDEV/GOJWT-Auth/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(authHandler *handler.AuthHandler) *gin.Engine {

	route := gin.Default()

	route.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	router := route.Group("/api")
	authenRouter := router.Group("/auth")

	authenRouter.POST("/register", authHandler.Register)
	authenRouter.POST("/login", authHandler.Login)

	return route
}
