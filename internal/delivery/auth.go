package delivery

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

var jwtSecret = []byte("Directorsad@12")

func AuthMiddleware(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			claims := jwt.MapClaims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return jwtSecret, nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			role := claims["role"].(string)
			for _, allowed := range allowedRoles {
				if role == allowed {
					ctx := context.WithValue(r.Context(), "role", role)
					next.ServeHTTP(w, r.WithContext(ctx))
					return
				}
			}

			http.Error(w, "Forbidden", http.StatusForbidden)
		})
	}
}
