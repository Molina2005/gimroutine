package routes

import (
	"modulo/internal/config"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Funcion para manejar handlers de cada uno de susu modulos, con su URL de acceso
func RoutesWebHook(connect *pgxpool.Pool, r *chi.Mux) {
	// Llamado funcion para obtener informacion de handlers de modulos
	promises := config.LayersOfModulesWebHook(connect)

	r.HandleFunc("/NotifyChangesVitals", promises.WebHooks.NotifyChangesVitals)
}
