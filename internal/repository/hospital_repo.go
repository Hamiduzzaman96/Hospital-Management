package repository

import (
	"database/sql"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/domain"
)

type HospitalRepository struct {
	db *sql.DB
}

func NewHospitalRepository(db *sql.DB) *HospitalRepository {
	return &HospitalRepository{db: db}
}

func (r *HospitalRepository) Create(hospital *domain.Hospital) error {
	err := r.db.QueryRow(
		`INSERT INTO hospitals (name, address) VALUES ($1,$2) RETURNING id`,
		hospital.Name, hospital.Address,
	).Scan(&hospital.ID)
	return err
}

func (r *HospitalRepository) Update(hospital *domain.Hospital) error {
	_, err := r.db.Exec(
		`UPDATE hospitals SET name = $1, address = $2 WHERE id = $3`,
		hospital.Name, hospital.Address, hospital.ID,
	)
	return err
}

func (r *HospitalRepository) Delete(id int64) error {
	result, err := r.db.Exec(
		`DELETE FROM hospitals WHERE id = $1`,
		id,
	)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
func (r *HospitalRepository) GetByID(id int64) (*domain.Hospital, error) {
	var hospital domain.Hospital

	err := r.db.QueryRow(
		`SELECT id, name, address FROM hospitals WHERE id = $1`,
		id,
	).Scan(&hospital.ID, &hospital.Name, &hospital.Address)
	if err != nil {
		return nil, err
	}
	return &hospital, nil
}

func (r *HospitalRepository) List(search string, limit, offset int64) ([]domain.Hospital, error) {
	rows, err := r.db.Query(
		`SELECT id,name,address FROM hospitals WHERE name ILIKE '%'||$1||'%'
		ORDER BY id DESC
		LIMIT $2 OFFSET $3`,
		search, limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hospitals []domain.Hospital

	for rows.Next() {
		var h domain.Hospital
		if err := rows.Scan(&h.ID, &h.Name, &h.Address); err != nil {
			return nil, err
		}
		hospitals = append(hospitals, h)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return hospitals, nil
}

func (r *HospitalRepository) Count(search string) (int64, error) {
	var count int64
	err := r.db.QueryRow(
		`SELECT COUNT(*) FROM hospitals WHERE name ILIKE '%'||$1||'%'`,
		search,
	).Scan(&count)
	return count, err
}
