package doctor

import (
	"net/http"
	"strconv"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/domain"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/middleware"
	"github.com/Hamiduzzaman96/Hospital-Management.git/pkg/helper"
)

func (h *DoctorHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

	// 2. get user from context
	user := r.Context().Value(middleware.UserContextKey).(domain.User)

	// 3. call usecase
	if err := h.du.Delete(user, id); err != nil {
		helper.Error(w, http.StatusForbidden, err.Error())
	}

	helper.Success(w, 200, "Doctor deleted successfully", nil)
}
