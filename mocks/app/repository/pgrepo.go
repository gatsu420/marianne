package mockrepository

import (
	"context"

	"github.com/gatsu420/marianne/app/repository"
	"github.com/gatsu420/marianne/common/tests"
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
		Type:         tests.MockPGText(),
		IntakeStatus: tests.MockPGText(),
		Feeder:       tests.MockPGText(),
		Location:     tests.MockPGText(),
		Remarks:      tests.MockPGText(),
		CreatedAt:    tests.MockPGTimestamptz(),
		UpdatedAt:    tests.MockPGTimestamptz(),
	}, nil
}

func (m *mockPGRepo) ListFood(ctx context.Context, args repository.ListFoodArgs) ([]repository.ListFoodRow, error) {
	if m.expectedErr != nil {
		return []repository.ListFoodRow{}, m.expectedErr
	}
	return []repository.ListFoodRow{
		{
			ID:           99,
			Name:         "mock",
			Type:         tests.MockPGText(),
			IntakeStatus: tests.MockPGText(),
			Feeder:       tests.MockPGText(),
			Location:     tests.MockPGText(),
			Remarks:      tests.MockPGText(),
			CreatedAt:    tests.MockPGTimestamptz(),
			UpdatedAt:    tests.MockPGTimestamptz(),
		},
	}, nil
}

func (m *mockPGRepo) CreateFood(ctx context.Context, args repository.CreateFoodArgs) error {
	if m.expectedErr != nil {
		return m.expectedErr
	}
	return nil
}
