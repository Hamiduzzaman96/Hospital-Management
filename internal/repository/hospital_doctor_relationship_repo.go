package repository

import (
	"database/sql"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/domain"
)

type HospitalDoctorRelationship struct {
	db *sql.DB
}

func NewHospitalDoctorRelationship(db *sql.DB) *HospitalDoctorRelationship {
	return &HospitalDoctorRelationship{db: db}
}

func (r *HospitalDoctorRelationship) DoctorAssign(h *domain.HospitalDoctorRelationship) error {
	_, err := r.db.Exec(
		`INSERT INTO hospital_doctor_rel (hospital_id, doctor_id) VALUES ($1, $2)`,
		h.HospitalID, h.DoctorID,
	)
	return err
}

func (r *HospitalDoctorRelationship) RemoveDoctor(hospitalID, doctorID int64) error {
	_, err := r.db.Exec(
		`DELETE FROM hospital_doctor_rel WHERE hospital_id = $1 AND doctor_id = $2`,
		hospitalID, doctorID,
	)
	return err
}

func (r *HospitalDoctorRelationship) ListDoctorsByHospital(hospitalID int64) ([]domain.Doctor, error) {
	rows, err := r.db.Query(
		`SELECT d.doc_id, d.name, d.email
		 FROM doctors d
		 JOIN hospital_doctor_rel hdr
		   ON d.doc_id = hdr.doctor_id
		 WHERE hdr.hospital_id = $1`,
		hospitalID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var doctors []domain.Doctor

	for rows.Next() {
		var d domain.Doctor
		if err := rows.Scan(&d.DocID, &d.Name, &d.Email); err != nil {
			return nil, err
		}
		doctors = append(doctors, d)
	}

	return doctors, nil
}
