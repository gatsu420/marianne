package food_test

import (
	"context"
	"errors"
	"testing"

	"github.com/gatsu420/marianne/app/usecases/food"
	commonerr "github.com/gatsu420/marianne/common/errors"
	"github.com/gatsu420/marianne/common/tests"
	mockrepository "github.com/gatsu420/marianne/mocks/app/repository"
	"github.com/jackc/pgx/v5"
)

func Test_GetFood(t *testing.T) {
	testCases := []struct {
		testName    string
		id          int
		repoErr     error
		expectedRow *food.GetFoodRow
		expectedErr error
	}{
		{
			testName:    "repo error",
			id:          99,
			repoErr:     errors.New("some repo error"),
			expectedRow: nil,
			expectedErr: commonerr.New(commonerr.ErrMsgInternal, commonerr.ErrInternal),
		},
		{
			testName:    "food is not found",
			id:          99,
			repoErr:     pgx.ErrNoRows,
			expectedRow: nil,
			expectedErr: commonerr.New(commonerr.ErrMsgFoodNotFound, commonerr.ErrFoodNotFound),
		},
		{
			testName: "success",
			id:       99,
			repoErr:  nil,
			expectedRow: &food.GetFoodRow{
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
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			mockPGRepo := mockrepository.NewMockPGRepo(
				mockrepository.WithExpectedErr(tc.repoErr),
			)
			usecase := food.NewUsecase(mockPGRepo)

			row, err := usecase.GetFood(context.Background(), tc.id)
			tests.AssertEqual(t, row, tc.expectedRow)
			tests.AssertEqual(t, err, tc.expectedErr)
		})
	}
}
