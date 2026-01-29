package domain

type Hospital struct {
	HospitalID int64  `json:"hospital_id"`
	Name       string `json:"name"`
	Address    string `json:"address"`
}
