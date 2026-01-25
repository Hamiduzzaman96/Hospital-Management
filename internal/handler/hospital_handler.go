package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/domain"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/usecase"
)

type HospitalHandler struct {
	hu *usecase.HospitalUsecase
}

func NewHospitalHandler(huc *usecase.HospitalUsecase) *HospitalHandler {
	return &HospitalHandler{hu: huc}
}

func (h *HospitalHandler) Create(w http.ResponseWriter, r *http.Request) {
	var hospitalList domain.Hospital

	if err := json.NewDecoder(r.Body).Decode(&hospitalList); err != nil {
		http.Error(w, "Invalid JSON format", 400)
		return
	}

	role := domain.User{
		Role: domain.SuperAdmin,
	}

	h.hu.Create(role, &hospitalList)
}
