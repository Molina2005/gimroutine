package config

import (
	"modulo/internal/webhook"
	webkooks "modulo/internal/webhook"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PromiseWebHook struct {
	WebHooks *webhook.HandlerWebHook
}

func LayersOfModulesWebHook(connect *pgxpool.Pool) *PromiseWebHook {
	repoWebHook := webkooks.NewRepository(connect)
	serviceWebHook := webkooks.NewService(repoWebHook)
	handlerWebHook := webkooks.NewHandler(serviceWebHook)

	return &PromiseWebHook{
		WebHooks: handlerWebHook,
	}
}
