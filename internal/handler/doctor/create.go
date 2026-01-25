package doctor

import (
	"encoding/json"
	"net/http"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/domain"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/middleware"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/usecase"
)

type DoctorHandler struct {
	du *usecase.DoctorUsecase
}

func NewDoctorHandler(dh *usecase.DoctorUsecase) *DoctorHandler {
	return &DoctorHandler{du: dh}
}

func (h *DoctorHandler) Create(w http.ResponseWriter, r *http.Request) {
	var doctor domain.Doctor

	if err := json.NewDecoder(r.Body).Decode(&doctor); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, ok := r.Context().Value(middleware.UserContextKey).(domain.User)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	doctor.HospitalID = user.HospitalID

	if err := h.du.Create(user, user.HospitalID, &doctor); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Doctor created successfully",
	})
}
