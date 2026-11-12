package httpmiddlewares

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

func isAllowedOrigin(origin string) bool {
	u, err := url.Parse(origin)
	if err != nil || u.Scheme != "https" {
		return false
	}

	host := u.Hostname()

	// Dominio estable de producción
	if host == "mabtrosoftweb-programming10.vercel.app" {
		return true
	}

	// Dominios Preview del proyecto
	return strings.HasPrefix(host, "mabtrosoftweb-") &&
		strings.HasSuffix(host, "-programming10.vercel.app")
}

func EnableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		origin := r.Header.Get("Origin")

		log.Printf("CORS DEBUG | Method: %s | Origin: %s", r.Method, origin)

		if isAllowedOrigin(origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
		}

		w.Header().Set("Access-Control-Allow-Methods",
			"GET, POST, PUT, PATCH, DELETE, OPTIONS")

		w.Header().Set("Access-Control-Allow-Headers",
			"Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
