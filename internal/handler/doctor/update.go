package doctor

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/domain"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/middleware"
	"github.com/Hamiduzzaman96/Hospital-Management.git/pkg/helper"
)

func (h *DoctorHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "doctor id is required", http.StatusBadRequest)
		return
	}

	did, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid doctor id", http.StatusBadRequest)
		return
	}

	var doctor domain.Doctor
	if err := json.NewDecoder(r.Body).Decode(&doctor); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	user := r.Context().Value(middleware.UserContextKey).(domain.User)

	doctor.DocID = did
	doctor.HospitalID = user.HospitalID

	if err := h.du.Update(user, &doctor); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	helper.Success(w, 200, "Doctor updated successfully", nil)
}
