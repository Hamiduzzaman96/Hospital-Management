package user

import "github.com/Hamiduzzaman96/Hospital-Management.git/internal/usecase"

type UserHandler struct {
	uh *usecase.UserUsecase
}

func NewUserHandler(uh *usecase.UserUsecase) *UserHandler {
	return &UserHandler{uh: uh}
}
