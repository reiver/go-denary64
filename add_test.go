package denary64

import (
	"fmt"

	"testing"
)

func TestAddError(t *testing.T) {

	const maxUint64 uint64 = 1<<64 - 1

	tests := []struct{
		Left           uint64
		Right          uint64
		ExpectedResult uint64
		ExpectedErr    error
	}{
		{
			Left:           maxUint64,
			Right:          1,
			ExpectedResult: 0,
			ExpectedErr:    fmt.Errorf("denary64: overflow; %d + %d", maxUint64, 1),
		},
		{
			Left:           maxUint64,
			Right:          2,
			ExpectedResult: 1,
			ExpectedErr:    fmt.Errorf("denary64: overflow; %d + %d", maxUint64, 2),
		},
		{
			Left:           maxUint64,
			Right:          3,
			ExpectedResult: 2,
			ExpectedErr:    fmt.Errorf("denary64: overflow; %d + %d", maxUint64, 3),
		},

		{
			Left:           maxUint64,
			Right:          maxUint64-1,
			ExpectedResult: maxUint64-2,
			ExpectedErr:    fmt.Errorf("denary64: overflow; %d + %d", maxUint64, maxUint64-1),
		},
		{
			Left:           maxUint64,
			Right:          maxUint64,
			ExpectedResult: maxUint64-1,
			ExpectedErr:    fmt.Errorf("denary64: overflow; %d + %d", maxUint64, maxUint64),
		},









		{
			Left:           1,
			Right:          maxUint64,
			ExpectedResult: 0,
			ExpectedErr:    fmt.Errorf("denary64: overflow; %d + %d", 1, maxUint64),
		},
		{
			Left:           2,
			Right:          maxUint64,
			ExpectedResult: 1,
			ExpectedErr:    fmt.Errorf("denary64: overflow; %d + %d", 2, maxUint64),
		},
		{
			Left:           3,
			Right:          maxUint64,
			ExpectedResult: 2,
			ExpectedErr:    fmt.Errorf("denary64: overflow; %d + %d", 3, maxUint64),
		},

		{
			Left:           maxUint64-1,
			Right:          maxUint64,
			ExpectedResult: maxUint64-2,
			ExpectedErr:    fmt.Errorf("denary64: overflow; %d + %d", maxUint64-1, maxUint64),
		},
		{
			Left:           maxUint64,
			Right:          maxUint64,
			ExpectedResult: maxUint64-1,
			ExpectedErr:    fmt.Errorf("denary64: overflow; %d + %d", maxUint64, maxUint64),
		},









		{
			Left:           maxUint64-1,
			Right:          2,
			ExpectedResult: 0,
			ExpectedErr:    fmt.Errorf("denary64: overflow; %d + %d", maxUint64-1, 2),
		},
		{
			Left:           maxUint64-1,
			Right:          3,
			ExpectedResult: 1,
			ExpectedErr:    fmt.Errorf("denary64: overflow; %d + %d", maxUint64-1, 3),
		},
		{
			Left:           maxUint64-1,
			Right:          4,
			ExpectedResult: 2,
			ExpectedErr:    fmt.Errorf("denary64: overflow; %d + %d", maxUint64-1, 4),
		},

		{
			Left:           maxUint64-1,
			Right:          maxUint64-1,
			ExpectedResult: maxUint64-3,
			ExpectedErr:    fmt.Errorf("denary64: overflow; %d + %d", maxUint64-1, maxUint64-1),
		},
		{
			Left:           maxUint64-1,
			Right:          maxUint64,
			ExpectedResult: maxUint64-2,
			ExpectedErr:    fmt.Errorf("denary64: overflow; %d + %d", maxUint64-1, maxUint64),
		},









		{
			Left:           2,
			Right:          maxUint64-1,
			ExpectedResult: 0,
			ExpectedErr:    fmt.Errorf("denary64: overflow; %d + %d", 2, maxUint64-1),
		},
		{
			Left:           3,
			Right:          maxUint64-1,
			ExpectedResult: 1,
			ExpectedErr:    fmt.Errorf("denary64: overflow; %d + %d", 3, maxUint64-1),
		},
		{
			Left:           4,
			Right:          maxUint64-1,
			ExpectedResult: 2,
			ExpectedErr:    fmt.Errorf("denary64: overflow; %d + %d", 4, maxUint64-1),
		},

		{
			Left:           maxUint64-1,
			Right:          maxUint64-1,
			ExpectedResult: maxUint64-3,
			ExpectedErr:    fmt.Errorf("denary64: overflow; %d + %d", maxUint64-1, maxUint64-1),
		},
		{
			Left:           maxUint64,
			Right:          maxUint64-1,
			ExpectedResult: maxUint64-2,
			ExpectedErr:    fmt.Errorf("denary64: overflow; %d + %d", maxUint64, maxUint64-1),
		},
	}

	for testNumber, test := range tests {
		actualResult, actualErr := add(test.Left, test.Right)
		if nil == actualErr {
			t.Errorf("For test #%d, expected an error, but did not actually get one: %#v; %d + %d = %d", testNumber, actualErr, test.Left, test.Right, actualResult)
			continue
		}
		if expected, actual := fmt.Sprintf("%T", test.ExpectedErr), fmt.Sprintf("%T", actualErr); expected != actual {
			t.Errorf("For test #%d (%d + %d = %d)...", testNumber, test.Left, test.Right, actualResult)
			t.Errorf("\tEXPECTED error type: %s", expected)
			t.Errorf("\tACTUAL   error type: %s", actual)
			continue
		}
		if expected, actual := test.ExpectedErr.Error(), actualErr.Error(); expected != actual {
			t.Errorf("For test #%d (%d + %d = %d)...", testNumber, test.Left, test.Right, actualResult)
			t.Errorf("\tEXPECTED error message: %q", expected)
			t.Errorf("\tACTUAL   error message: %q", actual)
			continue
		}

		if expected, actual := test.ExpectedResult, actualResult; expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			continue
		}
	}
}

func TestAdd(t *testing.T) {

	const maxUint64 uint64 = 1<<64 - 1

	tests := []struct{
		Left     uint64
		Right    uint64
		Expected uint64
	}{
		{
			Left:     0,
			Right:    0,
			Expected: 0,
		},
		{
			Left:     0,
			Right:    1,
			Expected: 1,
		},
		{
			Left:     0,
			Right:    2,
			Expected: 2,
		},
		{
			Left:     0,
			Right:    3,
			Expected: 3,
		},
		{
			Left:     0,
			Right:    4,
			Expected: 4,
		},
		{
			Left:     0,
			Right:    5,
			Expected: 5,
		},

		{
			Left:     0,
			Right:    maxUint64-1,
			Expected: maxUint64-1,
		},
		{
			Left:     0,
			Right:    maxUint64,
			Expected: maxUint64,
		},









		{
			Left:     0,
			Right:    0,
			Expected: 0,
		},
		{
			Left:     1,
			Right:    0,
			Expected: 1,
		},
		{
			Left:     2,
			Right:    0,
			Expected: 2,
		},
		{
			Left:     3,
			Right:    0,
			Expected: 3,
		},
		{
			Left:     4,
			Right:    0,
			Expected: 4,
		},
		{
			Left:     5,
			Right:    0,
			Expected: 5,
		},

		{
			Left:     maxUint64-1,
			Right:    0,
			Expected: maxUint64-1,
		},
		{
			Left:     maxUint64,
			Right:    0,
			Expected: maxUint64,
		},









		{
			Left:     1,
			Right:    0,
			Expected: 1,
		},
		{
			Left:     1,
			Right:    1,
			Expected: 2,
		},
		{
			Left:     1,
			Right:    2,
			Expected: 3,
		},
		{
			Left:     1,
			Right:    3,
			Expected: 4,
		},
		{
			Left:     1,
			Right:    4,
			Expected: 5,
		},
		{
			Left:     0,
			Right:    5,
			Expected: 5,
		},

		{
			Left:     1,
			Right:    maxUint64-2,
			Expected: maxUint64-1,
		},
		{
			Left:     1,
			Right:    maxUint64-1,
			Expected: maxUint64,
		},









		{
			Left:     0,
			Right:    1,
			Expected: 1,
		},
		{
			Left:     1,
			Right:    1,
			Expected: 2,
		},
		{
			Left:     2,
			Right:    1,
			Expected: 3,
		},
		{
			Left:     3,
			Right:    1,
			Expected: 4,
		},
		{
			Left:     4,
			Right:    1,
			Expected: 5,
		},
		{
			Left:     5,
			Right:    1,
			Expected: 6,
		},

		{
			Left:     maxUint64-2,
			Right:    1,
			Expected: maxUint64-1,
		},
		{
			Left:     maxUint64-1,
			Right:    1,
			Expected: maxUint64,
		},
	}

	for testNumber, test := range tests {
		actual, err := add(test.Left, test.Right)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			continue
		}

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			continue
		}
	}
}
