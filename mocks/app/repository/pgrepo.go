package mockrepository

import (
	"context"
	"time"

	"github.com/gatsu420/marianne/app/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type mockPGRepo struct {
	expectedErr error
}

func NewMockPGRepo(opts ...func(*mockPGRepo)) *mockPGRepo {
	mock := &mockPGRepo{}
	for _, o := range opts {
		o(mock)
	}
	return mock
}

func WithExpectedErr(err error) func(*mockPGRepo) {
	return func(mp *mockPGRepo) {
		mp.expectedErr = err
	}
}

func (m *mockPGRepo) GetFood(ctx context.Context, id int) (repository.GetFoodRow, error) {
	if m.expectedErr != nil {
		return repository.GetFoodRow{}, m.expectedErr
	}
	return repository.GetFoodRow{
		ID:           99,
		Name:         "mock",
		Type:         pgtype.Text{String: "mock", Valid: true},
		IntakeStatus: pgtype.Text{String: "mock", Valid: true},
		Feeder:       pgtype.Text{String: "mock", Valid: true},
		Location:     pgtype.Text{String: "mock", Valid: true},
		Remarks:      pgtype.Text{String: "mock", Valid: true},
		CreatedAt:    pgtype.Timestamptz{Time: time.Date(2025, time.July, 4, 20, 47, 0, 0, time.UTC), Valid: true},
		UpdatedAt:    pgtype.Timestamptz{Time: time.Date(2025, time.July, 4, 20, 47, 0, 0, time.UTC), Valid: true},
	}, nil
}
