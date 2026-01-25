package usecase

import (
	"errors"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/domain"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/repository"
)

type DoctorUsecase struct {
	doctorRepo *repository.DoctorRepository
}

// constructor
func NewDoctorUsecase(d *repository.DoctorRepository) *DoctorUsecase {
	return &DoctorUsecase{
		doctorRepo: d,
	}
}

// create doctor
func (u *DoctorUsecase) Create(user domain.User, doctor *domain.Doctor) error {
	if user.Role != domain.HospitalAdmin {
		return errors.New("only hospital admin can create a doctor")
	}

	return u.doctorRepo.Create(doctor)
}

// update doctor info
func (u *DoctorUsecase) Update(user domain.User, doctor *domain.Doctor) error {
	if user.Role != domain.HospitalAdmin {
		return errors.New("only hospital admin can update a doctor")
	}
	return u.doctorRepo.Update(doctor)
}

// delete doctor
func (u *DoctorUsecase) Delete(user domain.User, doctorID int64) error {
	if user.Role != domain.HospitalAdmin {
		return errors.New("only hospital admin can delete a doctor")
	}
	return u.doctorRepo.Delete(doctorID)
}

// get doctor by id
func (u *DoctorUsecase) GetByID(doctorID int64) (*domain.Doctor, error) {
	return u.doctorRepo.GetByID(doctorID)
}

// list doctors with pagination and search
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
