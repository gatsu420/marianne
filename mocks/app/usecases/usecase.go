package mockusecases

import (
	"context"

	"github.com/gatsu420/marianne/app/usecases/food"
	"github.com/gatsu420/marianne/common/tests"
)

type mockUsecase struct {
	expectedErr error
}

func NewMockUsecase(opts ...func(*mockUsecase)) *mockUsecase {
	mock := &mockUsecase{}
	for _, o := range opts {
		o(mock)
	}
	return mock
}

func WithExpectedErr(err error) func(*mockUsecase) {
	return func(mu *mockUsecase) {
		mu.expectedErr = err
	}
}

func (m *mockUsecase) GetFood(ctx context.Context, id int) (*food.GetFoodRow, error) {
	if m.expectedErr != nil {
		return nil, m.expectedErr
	}
	return &food.GetFoodRow{
		ID:           99,
		Name:         "mock",
		Type:         tests.MockPGText().String,
		IntakeStatus: tests.MockPGText().String,
		Feeder:       tests.MockPGText().String,
		Location:     tests.MockPGText().String,
		Remarks:      tests.MockPGText().String,
		CreatedAt:    tests.MockPGTimestamptz().Time,
		UpdatedAt:    tests.MockPGTimestamptz().Time,
	}, nil
}

func (m *mockUsecase) ListFood(ctx context.Context, args *food.ListFoodArgs) ([]food.ListFoodRow, error) {
	if m.expectedErr != nil {
		return nil, m.expectedErr
	}
	return []food.ListFoodRow{
		{
			ID:           99,
			Name:         tests.MockPGText().String,
			Type:         tests.MockPGText().String,
			IntakeStatus: tests.MockPGText().String,
			Feeder:       tests.MockPGText().String,
			Location:     tests.MockPGText().String,
			Remarks:      tests.MockPGText().String,
			CreatedAt:    tests.MockPGTimestamptz().Time,
			UpdatedAt:    tests.MockPGTimestamptz().Time,
		},
	}, nil
}
