package httpmiddlewares

import (
	"net/http"
)

// Funcion para utilizacion de CORS - Manejo de permisos de dominios Backend - Frontend
func EnableCors(cors http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedOrigins := []string{
			"https://mabtrosoftweb.vercel.app",
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
		cors.ServeHTTP(w, r)
	})

}
