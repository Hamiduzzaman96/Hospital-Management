package repository

import (
	"database/sql"
	"errors"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(u *domain.User) error {
	return r.db.QueryRow(
		`INSERT INTO users (name, email, password, role, hospital_id) 
		 VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		u.Name, u.Email, u.Password, u.Role, u.HospitalID,
	).Scan(&u.ID)
}

func (r *UserRepository) GetByEmail(email string) (*domain.User, error) {
	var u domain.User

	err := r.db.QueryRow(
		`SELECT id, name, email, password, role, hospital_id 
		 FROM users 
		 WHERE email = $1`,
		email,
	).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Role, &u.HospitalID)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, err
	}

	return &u, nil
}
