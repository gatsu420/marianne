package repository

import "github.com/jackc/pgx/v5/pgxpool"

type PGRepo interface {
	GetFood(id int) (GetFoodRow, error)
}

type pgRepoImpl struct {
	pgxPool *pgxpool.Pool
}

func NewPGRepo(pgxPool *pgxpool.Pool) PGRepo {
	return &pgRepoImpl{
		pgxPool: pgxPool,
	}
}
