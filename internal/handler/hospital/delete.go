package hospital

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/domain"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/middleware"
)

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
