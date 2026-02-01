package hospital

import (
	"encoding/json"
	"net/http"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/domain"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/middleware"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/usecase"
	"github.com/Hamiduzzaman96/Hospital-Management.git/pkg/helper"
)

type HospitalHandler struct {
	hh *usecase.HospitalUsecase
}

func NewHospitalHandler(hc *usecase.HospitalUsecase) *HospitalHandler {
	return &HospitalHandler{hh: hc}
}

func (h *HospitalHandler) Create(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only Post Method Allowed to create", http.StatusMethodNotAllowed)
	}

	var hospital domain.Hospital

	if err := json.NewDecoder(r.Body).Decode(&hospital); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	user := r.Context().Value(middleware.UserContextKey).(domain.User)

	if err := h.hh.Create(user, &hospital); err != nil {
		helper.Error(w, http.StatusForbidden, err.Error())
	}

	helper.Success(w, 200, "Hospital created successfully", nil)
}
