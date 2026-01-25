package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/domain"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/usecase"
)

type UserHandler struct {
	uh *usecase.UserUsecase
}

func NewUserHandler(uh *usecase.UserUsecase) *UserHandler {
	return &UserHandler{uh: uh}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var userList domain.User

	if err := json.NewDecoder(r.Body).Decode(&userList); err != nil {
		http.Error(w, "Invalid JSON format", 400)
		return
	}

	err := h.uh.Register(
		userList.Name,
		userList.Email,
		userList.Password,
		userList.Role,
		userList.HospitalID,
	)
	if err != nil {
		http.Error(w, "Please provide correct information", 400)
	}

	w.WriteHeader(200)

}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var userList domain.User

	if err := json.NewDecoder(r.Body).Decode(&userList); err != nil {
		http.Error(w, "Invalid JSON format", 400)
		return
	}

	token, err := h.uh.Login(userList.Email, userList.Password)
	if err != nil {
		http.Error(w, "Please provide correct email and password", 401)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message":      "Login Successfully",
		"access_token": token,
	})
}
