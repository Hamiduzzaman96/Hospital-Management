package doctor

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/domain"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/middleware"
)

func (h *DoctorHandler) Update(w http.ResponseWriter, r *http.Request) {
	// 1. get id from query
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "doctor id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid doctor id", http.StatusBadRequest)
		return
	}

	// 2. decode body
	var doctor domain.Doctor
	if err := json.NewDecoder(r.Body).Decode(&doctor); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// 3. get user from context
	user := r.Context().Value(middleware.UserContextKey).(domain.User)

	// 4. enforce ownership
	doctor.DocID = id
	doctor.HospitalID = user.HospitalID

	// 5. call usecase
	if err := h.du.Update(user, &doctor); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Doctor updated successfully",
	})
}
