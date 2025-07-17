package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PGRepo interface {
	GetFood(ctx context.Context, id int) (GetFoodRow, error)
	ListFood(ctx context.Context, args ListFoodArgs) ([]ListFoodRow, error)
}

type pgRepoImpl struct {
	pgxPool *pgxpool.Pool
}

func NewPGRepo(pgxPool *pgxpool.Pool) PGRepo {
	return &pgRepoImpl{
		pgxPool: pgxPool,
	}
}

type QueryLogImpl struct{}

func (q QueryLogImpl) TraceQueryStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	// fmt.Println("start log")
	// fmt.Printf("%v \n", data.SQL)
	// // fmt.Printf("%v\n", data.Args...)
	return ctx
}

func (q QueryLogImpl) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
	// fmt.Println("end log")
	// fmt.Printf("%v \n", data.CommandTag)
	// fmt.Printf("%v \n", data.CommandTag.String())
}
