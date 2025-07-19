package tests

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func MockPGText() pgtype.Text {
	return pgtype.Text{
		String: "mock",
		Valid:  true,
	}
}

func MockPGTimestamptz() pgtype.Timestamptz {
	return pgtype.Timestamptz{
		Time:  time.Date(2025, time.July, 4, 20, 47, 0, 0, time.UTC),
		Valid: true,
	}
}
