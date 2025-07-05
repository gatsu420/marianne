package food_test

import (
	"context"
	"testing"
	"time"

	"github.com/gatsu420/marianne/app/usecases/food"
	mockrepository "github.com/gatsu420/marianne/mocks/app/repository"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
	mockPGRepo *mockrepository.MockPGRepo
	usecase    food.Usecase
	ctx        context.Context

	dummyPGText        pgtype.Text
	dummyTimestamp     time.Time
	dummyPGTimestamptz pgtype.Timestamptz
}

func (s *testSuite) SetupSuite() {
	s.dummyPGText = pgtype.Text{
		String: "dummy",
		Valid:  true,
	}
	s.dummyTimestamp = time.Date(2025, time.July, 4, 20, 47, 0, 0, time.UTC)
	s.dummyPGTimestamptz = pgtype.Timestamptz{
		Time:  s.dummyTimestamp,
		Valid: true,
	}
}

func (s *testSuite) SetupTest() {
	s.mockPGRepo = mockrepository.NewMockPGRepo(s.T())
	s.usecase = food.NewUsecase(s.mockPGRepo)
	s.ctx = context.Background()
}

func Test(t *testing.T) {
	suite.Run(t, &testSuite{})
}
