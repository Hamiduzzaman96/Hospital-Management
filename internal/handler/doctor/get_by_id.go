package doctor

import (
	"net/http"
	"strconv"

	"github.com/Hamiduzzaman96/Hospital-Management.git/pkg/helper"
)

func (h *DoctorHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Doctor ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid Doctor ID", http.StatusBadRequest)
		return
	}

	doctor, err := h.du.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	helper.Success(w, 200, "id found", doctor)
}
