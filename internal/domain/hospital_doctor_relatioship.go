package domain

type HospitalDoctorRelationship struct {
	HospitalID int64 `json:"hospital_id"`
	DoctorID   int64 `json:"doctor_id"`
}

type HospitalDoctorUsecase interface {
	AssignDoctor(user User, doctorID int64) error
	RemoveDoctor(user User, doctorID int64) error
	ListDoctors(user User) ([]Doctor, error)
}
