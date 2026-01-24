package repository

import (
	"database/sql"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/domain"
)

type DoctorRepository struct {
	db *sql.DB
}

func NewDoctorRepository(db *sql.DB) *DoctorRepository {
	return &DoctorRepository{db: db}
}

func (r *DoctorRepository) Create(doc *domain.Doctor) error {
	return r.db.QueryRow(
		`INSERT INTO doctors (name, email) VALUES ($1, $2) RETURNING id`,
		doc.Name, doc.Email,
	).Scan(&doc.DocID)
}

func (r *DoctorRepository) Update(doc *domain.Doctor) error {
	_, err := r.db.Exec(
		`UPDATE doctors SET name = $1, email = $2 WHERE id = $3`,
		doc.Name, doc.Email, doc.DocID,
	)
	return err
}

func (r *DoctorRepository) Delete(id int64) error {
	result, err := r.db.Exec(`DELETE FROM WHERE id = $1`, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *DoctorRepository) GetByID(id int64) (*domain.Doctor, error) {
	var doc domain.Doctor
	err := r.db.QueryRow(
		`SELECT id, name, email FROM doctors WHERE id = $1`,
		id,
	).Scan(&doc.DocID, &doc.Name, &doc.Email)
	if err != nil {
		return nil, err
	}
	return &doc, nil
}
func (r *DoctorRepository) List(search string, limit, offset int64) ([]domain.Doctor, error) {
	rows, err := r.db.Query(
		`SELECT id, name, email FROM doctors WHERE name ILIKE '%'||$1||'%'
		ORDER BY id DESC LIMIT &2 OFFSET $3`,
		search, limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var doctors []domain.Doctor
	for rows.Next() {
		var doc domain.Doctor
		rows.Scan(&doc.DocID, &doc.Name, &doc.Email)
		doctors = append(doctors, doc)
	}
	return doctors, nil
}
