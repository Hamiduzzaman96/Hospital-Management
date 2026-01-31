package middleware

import (
	"net/http"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/domain"
	"github.com/Hamiduzzaman96/Hospital-Management.git/pkg/helper"
)

func NewRoleMiddleware(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			user, ok := r.Context().Value(UserContextKey).(domain.User)
			if !ok {
				helper.Error(w, http.StatusUnauthorized, "Unauthorized user")
				return
			}

			for _, role := range allowedRoles {
				if user.Role == role {
					next.ServeHTTP(w, r)
					return
				}
			}

			helper.Error(w, http.StatusForbidden, "Forbidden")
		})
	}
}
