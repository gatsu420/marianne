package tests_test

import (
	"errors"
	"testing"

	"github.com/gatsu420/marianne/common/tests"
)

// This tests the test lmao
func Test_AssertEqual(t *testing.T) {
	t.Run("values are equal", func(t *testing.T) {
		cases := []struct {
			actual   interface{}
			expected interface{}
		}{
			{99, 99},
			{"something", "something"},
			{errors.New("some error"), errors.New("some error")},
			{[]int{1, 2, 3}, []int{1, 2, 3}},
		}

		for _, c := range cases {
			tests.AssertEqual(t, c.actual, c.expected)
		}
	})

	t.Run("values are not equal", func(t *testing.T) {
		cases := []struct {
			actual   interface{}
			expected interface{}
		}{
			{99, 98},
			{"something", "somethinc"},
			{errors.New("some error"), nil},
			{[]int{1, 2, 3}, nil},
		}

		for _, c := range cases {
			mockT := &testing.T{}
			tests.AssertEqual(mockT, c.actual, c.expected)
			if !mockT.Failed() {
				t.Error("non equal values must fail the test")
			}
		}
	})
}
