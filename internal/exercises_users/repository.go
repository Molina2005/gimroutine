package exercisesusers

import "github.com/jackc/pgx/v5/pgxpool"

type RepositoryExercisesUsers struct {
	db *pgxpool.Pool
}

func NewRepository(r *pgxpool.Pool) *RepositoryExercisesUsers {
	return &RepositoryExercisesUsers{db: r}
}
