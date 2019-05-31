package denary64

import (
	"testing"
)

func TestNothing(t *testing.T) {

	var actual   Result
	var expected Result = Nothing()

	if expected != actual {
		t.Errorf("Expected %#v, but actually got %#v.", expected, actual)
		return
	}
}
