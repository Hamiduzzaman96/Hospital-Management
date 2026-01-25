package usecase

import (
	"errors"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/domain"
)

type HospitalDoctorUsecase struct {
	repo domain.HospitalDoctorRepository
}

func NewHospitalDoctorUsecase(repo domain.HospitalDoctorRepository) *HospitalDoctorUsecase {
	return &HospitalDoctorUsecase{repo: repo}
}

// AssignDoctor assigns a doctor to the admin's hospital
func (u *HospitalDoctorUsecase) AssignDoctor(user domain.User, doctorID int64) error {
	if user.Role != domain.HospitalAdmin {
		return errors.New("only hospital admin can assign doctor")
	}

	exists, err := u.repo.DoctorExists(doctorID)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("doctor not found")
	}

	rel := domain.HospitalDoctorRelationship{
		HospitalID: user.HospitalID,
		DoctorID:   doctorID,
	}

	return u.repo.AssignDoctor(rel)
}

// RemoveDoctor removes a doctor from the admin's hospital
func (u *HospitalDoctorUsecase) RemoveDoctor(user domain.User, doctorID int64) error {
	if user.Role != domain.HospitalAdmin {
		return errors.New("only hospital admin can remove doctor")
	}

	rel := domain.HospitalDoctorRelationship{
		HospitalID: user.HospitalID,
		DoctorID:   doctorID,
	}

	return u.repo.RemoveDoctor(rel)
}

// ListDoctors lists all doctors assigned to the admin's hospital
func (u *HospitalDoctorUsecase) ListDoctors(user domain.User) ([]domain.Doctor, error) {
	if user.Role != domain.HospitalAdmin {
		return nil, errors.New("only hospital admin can view doctors")
	}

	return u.repo.ListDoctorsByHospital(user.HospitalID)
}
