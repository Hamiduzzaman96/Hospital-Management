package hospitaldoctor

import (
	"net/http"
	"strconv"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/domain"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/middleware"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/usecase"
	"github.com/Hamiduzzaman96/Hospital-Management.git/pkg/helper"
)

type HospitalDoctorHandler struct {
	hosdoc usecase.HospitalDoctorUsecase
}

func NewHospitalDoctorHandler(uc usecase.HospitalDoctorUsecase) *HospitalDoctorHandler {
	return &HospitalDoctorHandler{hosdoc: uc}
}

func (h *HospitalDoctorHandler) AssignDoctor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only Post Method Allowed to Assaign a Doctor", http.StatusMethodNotAllowed)
	}

	docIdStr := r.URL.Query().Get("doctor_id")
	if docIdStr == "" {
		http.Error(w, "doctor id is required", http.StatusBadRequest)
		return
	}

	doctorID, err := strconv.ParseInt(docIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid doctor id", http.StatusBadRequest)
		return
	}

	user := r.Context().Value(middleware.UserContextKey).(domain.User)
	if err := h.hosdoc.AssignDoctor(user, doctorID); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	helper.Success(w, 200, "Doctor assigned to Hospital Successfully", nil)
}

func (h *HospitalDoctorHandler) RemoveDoctor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Only DELETE Method Allowed to Delete", http.StatusMethodNotAllowed)
	}

	docIdStr := r.URL.Query().Get("doctor_id")
	if docIdStr == "" {
		http.Error(w, "doctor id is required", http.StatusBadRequest)
		return
	}

	docId, err := strconv.ParseInt(docIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid doctor id", http.StatusBadRequest)
		return
	}
	user := r.Context().Value(middleware.UserContextKey).(domain.User)

	if err := h.hosdoc.RemoveDoctor(user, docId); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}

	helper.Success(w, 200, "Doctor removed successfully", nil)
}

func (h *HospitalDoctorHandler) ListByDoctor(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserContextKey).(domain.User)

	doctors, err := h.hosdoc.ListDoctors(user)
	if err != nil {
		helper.Error(w, http.StatusForbidden, err.Error())
		return
	}

	helper.Success(w, http.StatusOK, "Assigned doctors list", doctors)
}
