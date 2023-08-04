package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	//init router gin
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	server := &http.Server{
		Addr:    ":8888",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
