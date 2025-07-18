package testassert

import (
	"reflect"
	"testing"
)

func Equal(t *testing.T, actual interface{}, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("had %v, expected %v", actual, expected)
	}
}
