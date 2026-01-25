package domain

type Doctor struct {
	DocID      int64  `json:"doc_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	HospitalID int64  `json:"hospital_id"`
}
