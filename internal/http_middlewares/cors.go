package httpmiddlewares

import "net/http"

// Middleware para controlar los permisos CORS
func EnableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		allowedOrigins := []string{
			"https://mabtrosoftweb-programming10.vercel.app",
			"https://mabtrosoftweb-eux38lmbe-programming10.vercel.app",
		}

		origin := r.Header.Get("Origin")

		for _, allowed := range allowedOrigins {
			if origin == allowed {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Vary", "Origin")
				break
			}
		}

		w.Header().Set(
			"Access-Control-Allow-Methods",
			"GET, POST, PUT, PATCH, DELETE, OPTIONS",
		)

		w.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Authorization",
		)

		// Resolver solicitudes preflight
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
