package food_test

import (
	"context"
	"errors"
	"testing"

	commonerr "github.com/gatsu420/marianne/common/errors"

	"github.com/gatsu420/marianne/app/usecases/food"
	"github.com/gatsu420/marianne/common/tests"
	mockrepository "github.com/gatsu420/marianne/mocks/app/repository"
)

func Test_CreateFood(t *testing.T) {
	testCases := []struct {
		caseName    string
		args        *food.CreateFoodArgs
		repoErr     error
		expectedErr error
	}{
		{
			caseName: "repo error",
			args: &food.CreateFoodArgs{
				Name:           "mock",
				TypeID:         99,
				IntakeStatusID: 99,
				FeederID:       99,
				LocationID:     99,
				Remarks:        "mock",
			},
			repoErr:     errors.New("some repo error"),
			expectedErr: commonerr.New(commonerr.ErrMsgInternal, commonerr.ErrInternal),
		},
		{
			caseName: "successfully created food",
			args: &food.CreateFoodArgs{
				Name:           "mock",
				TypeID:         99,
				IntakeStatusID: 99,
				FeederID:       99,
				LocationID:     99,
				Remarks:        "mock",
			},
			repoErr:     nil,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.caseName, func(t *testing.T) {
			mockPGRepo := mockrepository.NewMockPGRepo(
				mockrepository.WithExpectedErr(tc.repoErr),
			)
			usecase := food.NewUsecase(mockPGRepo)

			err := usecase.CreateFood(context.Background(), tc.args)
			tests.AssertEqual(t, err, tc.expectedErr)
		})
	}
}
