package tests_test

import (
	"testing"
	"time"

	"github.com/gatsu420/marianne/common/tests"
)

func Test_MockPGText(t *testing.T) {
	tests.AssertEqual(t, tests.MockPGText().String, "mock")
	tests.AssertEqual(t, tests.MockPGText().Valid, true)
}

func Test_MockPGTimestamptz(t *testing.T) {
	tests.AssertEqual(t, tests.MockPGTimestamptz().Time, time.Date(2025, time.July, 4, 20, 47, 0, 0, time.UTC))
	tests.AssertEqual(t, tests.MockPGTimestamptz().Valid, true)
}
