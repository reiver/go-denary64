package denary64

import (
	"fmt"

	"testing"
)

func TestErrNil(t *testing.T) {

	var result Result = Err(nil)

	if Nothing() != result {
		t.Errorf("Expected Nothing(), but actually got %#v", result)
		return
	}
}

func TestErr(t *testing.T) {

	tests := []struct{
		Err      error
		Expected error
	}{}

	for i:=0; i<100; i++ {
		err := fmt.Errorf("Error #%d", i)

		test := struct{
			Err      error
			Expected error
		}{
			Err:      err,
			Expected: err,
		}

		tests = append(tests, test)
	}

	for testNumber, test := range tests {

		var actual Result = Err(test.Err)

		if Nothing() == actual {
			t.Errorf("For test #%d, did not expect denary.Nothing(), but actually got %#v", testNumber, actual)
			continue
		}

		if nil == actual.err {
			t.Errorf("For test #%d, did not expect internal err to be nil, but actually is %#v", testNumber, actual.err)
			continue
		}

		if expected := Err(test.Err); expected != actual {
			t.Errorf("For test #%d, expected #%v, but actually got %#v", testNumber, expected, actual)
			continue
		}
	}
}
