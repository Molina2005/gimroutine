package webhook

import "github.com/jackc/pgx/v5/pgxpool"

type RepoWebHook struct {
	db *pgxpool.Pool
}

func NewRepository(r *pgxpool.Pool) *RepoWebHook {
	return &RepoWebHook{db: r}
}
