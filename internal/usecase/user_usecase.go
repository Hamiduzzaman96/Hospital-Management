package usecase

import (
	"errors"
	"time"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/domain"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	userRepo *repository.UserRepository
	secret   string
}

func NewUserUsecase(r *repository.UserRepository, secret string) *UserUsecase {
	return &UserUsecase{
		userRepo: r,
		secret:   secret,
	}
}

func (u *UserUsecase) Register(name, email, password, role string, hospitalID int64) error {
	user, _ := u.userRepo.GetByEmail(email)
	if user != nil {
		return errors.New("user already registered")
	}

	pass, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	if role == domain.HospitalAdmin && hospitalID <= 0 {
		return errors.New("hospital_id is required for hospital admin")
	}

	if role == domain.SuperAdmin {
		hospitalID = 0
	}

	newUser := &domain.User{
		Name:       name,
		Email:      email,
		Password:   string(pass),
		Role:       role,
		HospitalID: hospitalID,
	}

	return u.userRepo.Create(newUser)
}

func (u *UserUsecase) Login(email, password string) (string, error) {
	user, err := u.userRepo.GetByEmail(email)
	if err != nil {
		return "", errors.New("You entered wrong credentials")
	}

	if bcrypt.CompareHashAndPassword(
		[]byte(user.Password), []byte(password),
	) != nil {
		return "", errors.New("Your email and password are wrong, please provide correct email and password")
	}

	claims := jwt.MapClaims{
		"userID":     user.ID,
		"role":       string(user.Role),
		"hospitalID": user.HospitalID,
		"expiry":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(u.secret))
}
