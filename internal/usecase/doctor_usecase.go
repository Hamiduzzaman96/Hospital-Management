package usecase

import (
	"errors"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/domain"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/repository"
)

type DoctorUsecase struct {
	doctorRepo *repository.DoctorRepository
}

func NewDoctorUsecase(d *repository.DoctorRepository) *DoctorUsecase {
	return &DoctorUsecase{
		doctorRepo: d,
	}
}

func (u *DoctorUsecase) Create(user domain.User, doctor *domain.Doctor) error {
	if user.Role != domain.HospitalAdmin {
		return errors.New("only hospital admin can create a doctor")
	}

	return u.doctorRepo.Create(doctor)
}

func (u *DoctorUsecase) Update(user domain.User, doctor *domain.Doctor) error {
	if user.Role != domain.HospitalAdmin {
		return errors.New("only hospital admin can update a doctor")
	}
	return u.doctorRepo.Update(doctor)
}

func (u *DoctorUsecase) Delete(user domain.User, doctorID int64) error {
	if user.Role != domain.HospitalAdmin {
		return errors.New("only hospital admin can delete a doctor")
	}
	return u.doctorRepo.Delete(doctorID)
}

func (u *DoctorUsecase) GetByID(doctorID int64) (*domain.Doctor, error) {
	return u.doctorRepo.GetByID(doctorID)
}

func (u *DoctorUsecase) List(search string, page, size int64) ([]domain.Doctor, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 1
	}
	offset := (page - 1) * size
	return u.doctorRepo.List(search, page, offset)
}
