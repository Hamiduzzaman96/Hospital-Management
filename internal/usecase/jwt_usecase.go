package usecase

import (
	"errors"
	"time"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type JwtUsecase struct {
	userRepo *repository.UserRepository
	secret   string
}

func NewJwtUsecase(r *repository.UserRepository, secret string) *JwtUsecase {
	return &JwtUsecase{
		userRepo: r,
		secret:   secret,
	}
}

func (u *JwtUsecase) Login(email, password string) (string, error) {
	user, err := u.userRepo.GetbyEmail(email)
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
