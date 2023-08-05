package service

import (
	"errors"
	"project-for-portfolioDEV/GOJWT-Auth/config"
	"project-for-portfolioDEV/GOJWT-Auth/model"
	"project-for-portfolioDEV/GOJWT-Auth/repository"
	"project-for-portfolioDEV/GOJWT-Auth/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
)

type IAuthenticationService interface {
	Login(users model.LoginRequest) (string, error)
	Register(user model.UserRequest)
}

type authenticationService struct {
	userRepo repository.IUserRepository
	validate *validator.Validate
}

func NewAuthenticationService(userRepo repository.IUserRepository, validate *validator.Validate) IAuthenticationService {
	return &authenticationService{
		userRepo: userRepo,
		validate: validate,
	}
}

func (service *authenticationService) Login(usersRequest model.LoginRequest) (string, error) {
	//1. Find username in db
	userLogin, err := service.userRepo.FindUserByUsername(usersRequest.UserStandard.Username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	config, _ := config.LoadConfig(".")

	//2. Verify password
	err = utils.VerifyPassword(userLogin.Password, usersRequest.UserStandard.Password)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	//3. Generate Token
	tokenGen, err := utils.GenerateToken(config.TokenExpiresIn, userLogin.Id, config.TokenSecret)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return tokenGen, nil
}

func (service *authenticationService) Register(user model.UserRequest) {

	//1. hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Errorf("failed register with password %v", err.Error())
	}

	//2. instead new password hash
	regisUser := model.Users{
		Username: user.Username,
		Email:    user.Email,
		Password: hashedPassword,
	}

	//3. save new user to register
	service.userRepo.SaveUser(regisUser)
}
