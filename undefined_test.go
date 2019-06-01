package denary64_test

import (
	"github.com/reiver/go-denary64"

	"testing"
)

func TestUndefined(t *testing.T) {

	var actual denary64.Type

	if expected := denary64.Undefined(); expected != actual {
		t.Errorf("Expected %#v but actually got %#v", expected, actual)
		return
	}
}
