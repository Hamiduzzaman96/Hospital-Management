package hospital

import (
	"encoding/json"
	"net/http"
	"strconv"
)

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
