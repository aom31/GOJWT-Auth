package main

import (
	"log"
	"net/http"
	"project-for-portfolioDEV/GOJWT-Auth/config"
	"project-for-portfolioDEV/GOJWT-Auth/handler"
	"project-for-portfolioDEV/GOJWT-Auth/model"
	"project-for-portfolioDEV/GOJWT-Auth/repository"
	"project-for-portfolioDEV/GOJWT-Auth/router"
	"project-for-portfolioDEV/GOJWT-Auth/service"

	"github.com/go-playground/validator/v10"
)

func main() {

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Could not load environment variables with %v", err)
	}

	//database
	db := config.ConnectionDB(&loadConfig)

	validate := validator.New()

	db.Table("users").AutoMigrate(&model.Users{})

	//Init Repository
	userRepository := repository.NewUserRepository(db)

	//Init service
	authenticationService := service.NewAuthenticationService(userRepository, validate)

	//Init handler
	authenticationHandler := handler.NewAuthHandler(authenticationService)

	routes := router.NewRouter(authenticationHandler)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
