package httpmiddlewares

import (
	"modulo/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Funcion obtener token JWT
func GenerateJWT(idUser int, role string) (string, error) {
	// Se agregan models.DataToken parametros y informacion de JWT
	data := models.DataToken{
		Id:   idUser,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			// Momento en que va a expirar el token
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			// Momento en que se creo el token
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	// Se obtiene el token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	// Se pasa el token a string, el token queda protegido con clave secreta
	return token.SignedString([]byte("JWTSECRETKEY"))
}
