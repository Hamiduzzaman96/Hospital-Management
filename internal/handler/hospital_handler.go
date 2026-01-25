package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/domain"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/middleware"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/usecase"
)

type HospitalHandler struct {
	hh *usecase.HospitalUsecase
}

func NewHospitalHandler(hc *usecase.HospitalUsecase) *HospitalHandler {
	return &HospitalHandler{hh: hc}
}

func (h *HospitalHandler) Create(w http.ResponseWriter, r *http.Request) {
	var hospital domain.Hospital

	if err := json.NewDecoder(r.Body).Decode(&hospital); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	user := r.Context().Value(middleware.UserContextKey).(domain.User)

	if err := h.hh.Create(user, &hospital); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hospital created successfully",
	})
}

func (h *HospitalHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Hospital ID is required", 400)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid hospital ID", http.StatusBadRequest)
		return
	}

	hospital, err := h.hh.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hospital)
}
