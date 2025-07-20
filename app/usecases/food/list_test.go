package food_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/gatsu420/marianne/app/usecases/food"
	commonerr "github.com/gatsu420/marianne/common/errors"
	"github.com/gatsu420/marianne/common/tests"
	mockrepository "github.com/gatsu420/marianne/mocks/app/repository"
)

func Test_ListFood(t *testing.T) {
	testCases := []struct {
		testName     string
		args         *food.ListFoodArgs
		repoErr      error
		expectedRows []food.ListFoodRow
		expectedErr  error
	}{
		{
			testName: "unable to get rows from repository",
			args: &food.ListFoodArgs{
				StartTimestamp: time.Now(),
				EndTimestamp:   time.Now(),
				Type:           "mock",
				IntakeStatus:   "mock",
				Feeder:         "mock",
				Location:       "mock",
			},
			repoErr:      errors.New("some repo error"),
			expectedRows: nil,
			expectedErr:  commonerr.New(commonerr.ErrMsgInternal, commonerr.ErrInternal),
		},
		{
			testName: "success",
			args: &food.ListFoodArgs{
				StartTimestamp: time.Now(),
				EndTimestamp:   time.Now(),
				Type:           "mock",
				IntakeStatus:   "mock",
				Feeder:         "mock",
				Location:       "mock",
			},
			repoErr: nil,
			expectedRows: []food.ListFoodRow{
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

			rows, err := usecase.ListFood(context.Background(), tc.args)
			tests.AssertEqual(t, rows, tc.expectedRows)
			tests.AssertEqual(t, err, tc.expectedErr)
		})
	}
}
