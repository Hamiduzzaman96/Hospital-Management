package hospital

import (
	"net/http"
	"strconv"

	"github.com/Hamiduzzaman96/Hospital-Management.git/pkg/helper"
)

func (h *HospitalHandler) List(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	search := query.Get("search")

	page, _ := strconv.ParseInt(query.Get("page"), 10, 64)
	size, _ := strconv.ParseInt(query.Get("size"), 10, 64)

	hospitals, err := h.hh.List(search, page, size)
	if err != nil {
		helper.Error(w, http.StatusForbidden, err.Error())
	}

	helper.Success(w, 200, "", hospitals)
}
