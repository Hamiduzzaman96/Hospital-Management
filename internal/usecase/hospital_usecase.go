package usecase

import (
	"errors"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/domain"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/repository"
)

type HospitalUsecase struct {
	hospitalRepo *repository.HospitalRepository
}

func NewHospitalUsecase(r *repository.HospitalRepository) *HospitalUsecase {
	return &HospitalUsecase{hospitalRepo: r}
}

func (u *HospitalUsecase) Create(user domain.User, hospital *domain.Hospital) error {
	if user.Role != domain.SuperAdmin {
		return errors.New("Only Super Admin Allowed to Create the Hospital")
	}
	return u.hospitalRepo.Create(hospital)
}

func (u *HospitalUsecase) Delete(user domain.User, id int64) error {
	if user.Role != domain.SuperAdmin {
		return errors.New("Only Super Admin Allowed to delete the hospital")
	}
	return u.hospitalRepo.Delete(id)
}

func (u *HospitalUsecase) GetByID(id int64) (*domain.Hospital, error) {
	return u.hospitalRepo.GetByID(id)
}

func (u *HospitalUsecase) ListHospitals(search string, page, size int64) ([]domain.Hospital, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	offset := (page - 1) * size
	return u.hospitalRepo.List(search, size, offset)
}
