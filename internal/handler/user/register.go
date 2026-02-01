package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/domain"
)

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only Post Method Allowed to Register", http.StatusMethodNotAllowed)
	}

	var user domain.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("JSON decode error:", err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := h.uh.Register(user.Name, user.Email, user.Password, user.Role, user.HospitalID); err != nil {
		log.Println("Register error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "user registered successfully"})
}
