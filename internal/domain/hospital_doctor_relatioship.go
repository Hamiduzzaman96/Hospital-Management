package domain

type HospitalDoctorRelationship struct {
	HospitalID int64 `json:"hospital_id"`
	DoctorID   int64 `json:"doctor_id"`
}
