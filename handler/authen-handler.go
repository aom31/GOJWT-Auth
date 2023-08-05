package handler

import (
	"net/http"
	"project-for-portfolioDEV/GOJWT-Auth/model"
	"project-for-portfolioDEV/GOJWT-Auth/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.IAuthenticationService
}

func NewAuthHandler(authService service.IAuthenticationService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (handler *AuthHandler) Login(ctx *gin.Context) {
	loginRequest := model.LoginRequest{}

	//1. Bind struct to json
	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		panic(err)
	}

	token, err := handler.authService.Login(loginRequest)
	if err != nil {
		response := model.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid username or password",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	tokenResponse := model.LoginResponse{
		TokenType: "Bearer",
		Token:     token,
	}

	response := model.Response{
		Code:    http.StatusOK,
		Status:  "Success",
		Message: "Successfully login",
		Data:    tokenResponse,
	}

	ctx.JSON(http.StatusOK, response)

}

func (handler *AuthHandler) Register(ctx *gin.Context) {
	var user model.UserRequest

	//1. Bind struct to json
	if err := ctx.ShouldBindJSON(&user); err != nil {
		panic(err)
	}

	handler.authService.Register(user)

	response := model.Response{
		Code:    http.StatusAccepted,
		Status:  "Success",
		Message: "Successfully created user",
		Data:    nil,
	}

	ctx.JSON(http.StatusAccepted, response)

}
