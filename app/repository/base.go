package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PGRepo interface {
	GetFood(ctx context.Context, id int) (GetFoodRow, error)
	// ListFood(ctx context.Context, args ListFoodArgs) ([]ListFoodRow, error)
}

type pgRepoImpl struct {
	pgxPool *pgxpool.Pool
}

func NewPGRepo(pgxPool *pgxpool.Pool) PGRepo {
	return &pgRepoImpl{
		pgxPool: pgxPool,
	}
}
