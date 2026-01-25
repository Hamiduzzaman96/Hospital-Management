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

func (h *HospitalHandler) List(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	search := query.Get("search")

	page, _ := strconv.ParseInt(query.Get("page"), 10, 64)
	size, _ := strconv.ParseInt(query.Get("size"), 10, 64)

	hospitals, err := h.hh.List(search, page, size)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(hospitals)
}

func (h *HospitalHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Hospital ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid hospital Id", http.StatusBadRequest)
		return
	}

	user, ok := r.Context().Value(middleware.UserContextKey).(domain.User)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if err := h.hh.Delete(user, id); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]string{
		"meassge": "Hospital deleted successfully",
	})
}
