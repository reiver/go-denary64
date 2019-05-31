package denary64

import (
	"testing"
)

func TestNone(t *testing.T) {

	var actual   Result
	var expected Result = None()

	if expected != actual {
		t.Errorf("Expected %#v, but actually got %#v.", expected, actual)
		return
	}
}
