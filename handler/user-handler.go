package handler

import "project-for-portfolioDEV/GOJWT-Auth/repository"

type UserHandler struct {
	userRepo repository.IUserRepository
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}
