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

type HospitalDoctorRepository interface {
	AssignDoctor(rel HospitalDoctorRelationship) error
	RemoveDoctor(rel HospitalDoctorRelationship) error
	ListDoctorsByHospital(hospitalID int64) ([]Doctor, error)
	DoctorExists(doctorID int64) (bool, error)
}
