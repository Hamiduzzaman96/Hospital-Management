package usecase

import (
	"errors"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/domain"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/repository"
)

type HospitalDoctorUsecase struct {
	repo repository.HospitalDoctorRelationship
}

func NewHospitalDoctorUsecase(repo repository.HospitalDoctorRelationship) *HospitalDoctorUsecase {
	return &HospitalDoctorUsecase{repo: repo}
}

func (u *HospitalDoctorUsecase) AssignDoctor(user domain.User, doctorID int64) error {
	if user.Role != domain.HospitalAdmin {
		return errors.New("only hospital admin can assign doctor")
	}

	rel := domain.HospitalDoctorRelationship{
		HospitalID: user.HospitalID,
		DoctorID:   doctorID,
	}

	return u.repo.DoctorAssign(&rel)
}

func (u *HospitalDoctorUsecase) RemoveDoctor(user domain.User, doctorID int64) error {
	if user.Role != domain.HospitalAdmin {
		return errors.New("only hospital admin can remove doctor")
	}

	rel := domain.HospitalDoctorRelationship{
		HospitalID: user.HospitalID,
		DoctorID:   doctorID,
	}

	return u.repo.RemoveDoctor(rel.HospitalID, rel.DoctorID)
}

func (u *HospitalDoctorUsecase) ListDoctors(user domain.User) ([]domain.Doctor, error) {
	if user.Role != domain.HospitalAdmin {
		return nil, errors.New("only hospital admin can view doctors")
	}

	return u.repo.ListDoctorsByHospital(user.HospitalID)
}
