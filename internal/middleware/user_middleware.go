package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/domain"
	"github.com/golang-jwt/jwt/v5"
)

const UserContextKey string = "user"

func NewAuthMiddleware(jwtsecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "No authorization header found", 401)
				return
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

			token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
				return []byte(jwtsecret), nil
			})
			if err != nil || !token.Valid {
				http.Error(w, "expired token", 401)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Invalid token", 401)
				return
			}

			user := domain.User{
				ID:         int64(claims["userID"].(float64)),
				Role:       claims["role"].(string),
				HospitalID: int64(claims["hospitalID"].(float64)),
			}
			ctx := context.WithValue(r.Context(), UserContextKey, user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
