package denary64_test

import (
	"github.com/reiver/go-denary64"

	"errors"

	"testing"
)

func TestResultGoString(t *testing.T) {

	tests := []struct{
		Value    denary64.Result
		Expected string
	}{
		{
			Value:     denary64.Nothing(),
			Expected: "denary64.Nothing()",
		},









		{
			Value:     denary64.Err(nil),
			Expected: "denary64.Nothing()",
		},



		{
			Value: denary64.Err(errors.New("Apple BANANA cherry")),
			Expected:      `denary64.Error("Apple BANANA cherry")`,
		},









		{
			Value:     denary64.Error("Apple BANANA cherry"),
			Expected: `denary64.Error("Apple BANANA cherry")`,
		},
		{
			Value:     denary64.Error(""),
			Expected: `denary64.Error("")`,
		},
		{
			Value:     denary64.Error("ONE two Three fOuR FiVe"),
			Expected: `denary64.Error("ONE two Three fOuR FiVe")`,
		},









		{
			Value:     denary64.Some(denary64.Uint64(0)),
			Expected: "denary64.Some(denary64.Uint64(0))",
		},
		{
			Value:     denary64.Some(denary64.Uint64(1)),
			Expected: "denary64.Some(denary64.Uint64(1))",
		},
		{
			Value:     denary64.Some(denary64.Uint64(2)),
			Expected: "denary64.Some(denary64.Uint64(2))",
		},
		{
			Value:     denary64.Some(denary64.Uint64(3)),
			Expected: "denary64.Some(denary64.Uint64(3))",
		},
		{
			Value:     denary64.Some(denary64.Uint64(4)),
			Expected: "denary64.Some(denary64.Uint64(4))",
		},
		{
			Value:     denary64.Some(denary64.Uint64(5)),
			Expected: "denary64.Some(denary64.Uint64(5))",
		},
		{
			Value:     denary64.Some(denary64.Uint64(6)),
			Expected: "denary64.Some(denary64.Uint64(6))",
		},
		{
			Value:     denary64.Some(denary64.Uint64(7)),
			Expected: "denary64.Some(denary64.Uint64(7))",
		},
		{
			Value:     denary64.Some(denary64.Uint64(8)),
			Expected: "denary64.Some(denary64.Uint64(8))",
		},
		{
			Value:     denary64.Some(denary64.Uint64(9)),
			Expected: "denary64.Some(denary64.Uint64(9))",
		},
		{
			Value:     denary64.Some(denary64.Uint64(10)),
			Expected: "denary64.Some(denary64.Uint64(10))",
		},
		{
			Value:     denary64.Some(denary64.Uint64(11)),
			Expected: "denary64.Some(denary64.Uint64(11))",
		},
		{
			Value:     denary64.Some(denary64.Uint64(12)),
			Expected: "denary64.Some(denary64.Uint64(12))",
		},
		{
			Value:     denary64.Some(denary64.Uint64(13)),
			Expected: "denary64.Some(denary64.Uint64(13))",
		},
		{
			Value:     denary64.Some(denary64.Uint64(14)),
			Expected: "denary64.Some(denary64.Uint64(14))",
		},
		{
			Value:     denary64.Some(denary64.Uint64(15)),
			Expected: "denary64.Some(denary64.Uint64(15))",
		},
		{
			Value:     denary64.Some(denary64.Uint64(16)),
			Expected: "denary64.Some(denary64.Uint64(16))",
		},
		{
			Value:     denary64.Some(denary64.Uint64(17)),
			Expected: "denary64.Some(denary64.Uint64(17))",
		},
		{
			Value:     denary64.Some(denary64.Uint64(18)),
			Expected: "denary64.Some(denary64.Uint64(18))",
		},
		{
			Value:     denary64.Some(denary64.Uint64(19)),
			Expected: "denary64.Some(denary64.Uint64(19))",
		},
		{
			Value:     denary64.Some(denary64.Uint64(20)),
			Expected: "denary64.Some(denary64.Uint64(20))",
		},
		{
			Value:     denary64.Some(denary64.Uint64(21)),
			Expected: "denary64.Some(denary64.Uint64(21))",
		},
		{
			Value:     denary64.Some(denary64.Uint64(22)),
			Expected: "denary64.Some(denary64.Uint64(22))",
		},
		{
			Value:     denary64.Some(denary64.Uint64(23)),
			Expected: "denary64.Some(denary64.Uint64(23))",
		},
		{
			Value:     denary64.Some(denary64.Uint64(24)),
			Expected: "denary64.Some(denary64.Uint64(24))",
		},
		{
			Value:     denary64.Some(denary64.Uint64(25)),
			Expected: "denary64.Some(denary64.Uint64(25))",
		},



		{
			Value:     denary64.Some(denary64.Float(1, 1)),
			Expected: "denary64.Some(denary64.Float(1, 1))",
		},
		{
			Value:     denary64.Some(denary64.Float(2, 2)),
			Expected: "denary64.Some(denary64.Float(2, 2))",
		},
		{
			Value:     denary64.Some(denary64.Float(3, 3)),
			Expected: "denary64.Some(denary64.Float(3, 3))",
		},
		{
			Value:     denary64.Some(denary64.Float(4, 4)),
			Expected: "denary64.Some(denary64.Float(4, 4))",
		},
		{
			Value:     denary64.Some(denary64.Float(5, 5)),
			Expected: "denary64.Some(denary64.Float(5, 5))",
		},
		{
			Value:     denary64.Some(denary64.Float(6, 6)),
			Expected: "denary64.Some(denary64.Float(6, 6))",
		},
		{
			Value:     denary64.Some(denary64.Float(7, 7)),
			Expected: "denary64.Some(denary64.Float(7, 7))",
		},
		{
			Value:     denary64.Some(denary64.Float(8, 8)),
			Expected: "denary64.Some(denary64.Float(8, 8))",
		},
		{
			Value:     denary64.Some(denary64.Float(9, 9)),
			Expected: "denary64.Some(denary64.Float(9, 9))",
		},
		{
			Value:     denary64.Some(denary64.Float(10, 10)),
			Expected: "denary64.Some(denary64.Float(1, 9))",
		},
	}

	for testNumber, test := range tests {

		if expected, actual := test.Expected, test.Value.GoString(); expected != actual {
			t.Errorf("For test #%d...", testNumber)
			t.Errorf("\tEXPECTED: %q", expected)
			t.Errorf("\tACTUAL:   %q", actual)
			continue
		}
	}
}
