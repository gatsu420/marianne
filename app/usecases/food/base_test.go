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
	dummyPGTimestamptz pgtype.Timestamptz
}

func (s *testSuite) SetupSuite() {
	s.dummyPGText = pgtype.Text{
		String: "dummy",
		Valid:  true,
	}
	s.dummyPGTimestamptz = pgtype.Timestamptz{
		Time:  time.Now(),
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
