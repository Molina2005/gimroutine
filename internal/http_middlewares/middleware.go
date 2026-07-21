package httpmiddlewares

import (
	"context"
	"modulo/internal/models"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// Middleware que valida que el usuario envíe un token JWT válido
// antes de permitir el acceso a la ruta protegida.
func MiddlewareAccessUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			http.Error(w, "token requerido", http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimPrefix(authorization, "Bearer ")
		claims := &models.DataToken{}
		token, err := jwt.ParseWithClaims(
			tokenString,
			claims,
			func(t *jwt.Token) (any, error) {
				return []byte("JWTKEY"), nil
			},
		)
		if err != nil || !token.Valid {
			http.Error(w, "token invalido", http.StatusUnauthorized)
		}
		ctx := context.WithValue(r.Context(), "user", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
