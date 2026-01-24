package usecase

import (
	"errors"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/domain"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/repository"
)

type DoctorUsecase struct {
	doctorRepo            *repository.DoctorRepository
	hospitalDoctorRelRepo *repository.HospitalDoctorRelationship
}

func NewDoctorUsecase(d *repository.DoctorRepository, h *repository.HospitalDoctorRelationship) *DoctorUsecase {
	return &DoctorUsecase{doctorRepo: d, hospitalDoctorRelRepo: h}
}

func (u *DoctorUsecase) Create(role domain.User, hospitalID int64, d *domain.Doctor) error {
	if role.Role != domain.HospitalAdmin {
		return errors.New("Only Hospital Admin Allowed to Create a Doctor")
	}

	err := u.doctorRepo.Create(d)
	if err != nil {
		return err
	}

	return u.hospitalDoctorRelRepo.DoctorAssign(&domain.HospitalDoctorRelationship{
		HospitalID: hospitalID,
		DoctorID:   d.DocID,
	})
}

func (u *DoctorUsecase) Update(role domain.User, d *domain.Doctor) error {
	if role.Role != domain.HospitalAdmin {
		return errors.New("Only Only Hospital Admin Allowed to Update a Doctor")
	}
	return u.doctorRepo.Update(d)
}

func (u *DoctorUsecase) Delete(role domain.User, docID int64) error {
	if role.Role != domain.HospitalAdmin {
		return errors.New("Only Only Hospital Admin Allowed to Delete a Doctor")
	}
	return u.doctorRepo.Delete(docID)
}

func (u *DoctorUsecase) GetByID(docID int64) (*domain.Doctor, error) {
	return u.doctorRepo.GetByID(docID)
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
