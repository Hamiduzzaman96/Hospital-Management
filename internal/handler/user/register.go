package user

import (
	"encoding/json"
	"net/http"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/domain"
)

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
		return
	}

	w.WriteHeader(200)

}
