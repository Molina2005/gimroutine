package httpmiddlewares

import (
	"net/http"
	"os"
)

// Funcion para utilizacion de CORS - Manejo de permisos de dominios Backend - Frontend
func EnableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		domainLocal := os.Getenv("DOMAIN_LOCAL")
		domainTesting := os.Getenv("DOMAIN_TESTING")
		domainProduction := os.Getenv("DOMAIN_PRODUCTION")
		allowedOrigins := []string{
			// ENTORNO LOCAL
			domainLocal,
			// ENTORNO PRUEBAS
			domainTesting,
			// ENTORNO PRODUCCION
			domainProduction,
		}

		origin := r.Header.Get("Origin")
		for _, allowed := range allowedOrigins {
			if origin == allowed {
				w.Header().Set("Access-Control-Allow-Origin", origin) // Especifica que origenes pueden acceder al recurso
				w.Header().Set("Vary", "Origin")
				break
			}
		}
		(w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Enumera los metodos HTTP permitidos para las solicitudes de origen cruzados
		(w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")     // Especifica que encabezados se pueden usar en la solicitud de origen cruzado
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})

}
