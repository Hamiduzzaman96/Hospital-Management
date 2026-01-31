package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/domain"
	"github.com/Hamiduzzaman96/Hospital-Management.git/pkg/helper"
	"github.com/golang-jwt/jwt/v5"
)

const UserContextKey string = "user"

func NewJwtMiddleware(jwtsecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				helper.Error(w, http.StatusUnauthorized, "No authorization header found")
				return
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenStr == authHeader {
				helper.Error(w, http.StatusUnauthorized, "Invalid authorization format")
				return
			}

			token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
				return []byte(jwtsecret), nil
			})
			if err != nil || !token.Valid {
				helper.Error(w, http.StatusUnauthorized, "Expired or invalid token")
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				helper.Error(w, http.StatusUnauthorized, "Invalid token claims")
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
