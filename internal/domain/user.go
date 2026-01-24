package domain

const (
	SuperAdmin    string = "super_admin"
	HospitalAdmin string = "hospital_admin"
)

type User struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	HospitalID int64  `json:"hospital_id"`
}
