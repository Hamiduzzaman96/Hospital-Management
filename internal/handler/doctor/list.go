package doctor

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *DoctorHandler) List(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET Method Allowed to Show List", http.StatusMethodNotAllowed)
	}

	query := r.URL.Query()

	search := query.Get("search")

	page, _ := strconv.ParseInt(query.Get("page"), 10, 64)
	size, _ := strconv.ParseInt(query.Get("size"), 10, 64)

	doctors, err := h.du.List(search, page, size)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(doctors)
}
