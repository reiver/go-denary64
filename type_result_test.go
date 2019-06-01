package denary64_test

import (
	"github.com/reiver/go-denary64"

	"testing"
)

func TestTypeResult(t *testing.T) {

	tests := []struct{
		Value    denary64.Type
		Expected denary64.Result
	}{
		{
			Value:    denary64.Undefined(),
			Expected: denary64.Nothing(),
		},



		{
			Value:                       denary64.Uint64(0),
			Expected: denary64.Something(denary64.Uint64(0)),
		},
		{
			Value:                       denary64.Uint64(1),
			Expected: denary64.Something(denary64.Uint64(1)),
		},
		{
			Value:                       denary64.Uint64(2),
			Expected: denary64.Something(denary64.Uint64(2)),
		},
		{
			Value:                       denary64.Uint64(3),
			Expected: denary64.Something(denary64.Uint64(3)),
		},
		{
			Value:                       denary64.Uint64(4),
			Expected: denary64.Something(denary64.Uint64(4)),
		},
		{
			Value:                       denary64.Uint64(5),
			Expected: denary64.Something(denary64.Uint64(5)),
		},
		{
			Value:                       denary64.Uint64(6),
			Expected: denary64.Something(denary64.Uint64(6)),
		},
		{
			Value:                       denary64.Uint64(7),
			Expected: denary64.Something(denary64.Uint64(7)),
		},
		{
			Value:                       denary64.Uint64(8),
			Expected: denary64.Something(denary64.Uint64(8)),
		},
		{
			Value:                       denary64.Uint64(9),
			Expected: denary64.Something(denary64.Uint64(9)),
		},
		{
			Value:                       denary64.Uint64(10),
			Expected: denary64.Something(denary64.Uint64(10)),
		},
		{
			Value:                       denary64.Uint64(11),
			Expected: denary64.Something(denary64.Uint64(11)),
		},
		{
			Value:                       denary64.Uint64(12),
			Expected: denary64.Something(denary64.Uint64(12)),
		},
		{
			Value:                       denary64.Uint64(13),
			Expected: denary64.Something(denary64.Uint64(13)),
		},









		{
			Value:                       denary64.Float(1, 0),
			Expected: denary64.Something(denary64.Float(1, 0)),
		},
		{
			Value:                       denary64.Float(1, 1),
			Expected: denary64.Something(denary64.Float(1, 1)),
		},
		{
			Value:                       denary64.Float(1, 2),
			Expected: denary64.Something(denary64.Float(1, 2)),
		},
		{
			Value:                       denary64.Float(1, 3),
			Expected: denary64.Something(denary64.Float(1, 3)),
		},
		{
			Value:                       denary64.Float(1, 4),
			Expected: denary64.Something(denary64.Float(1, 4)),
		},
		{
			Value:                       denary64.Float(1, 5),
			Expected: denary64.Something(denary64.Float(1, 5)),
		},
		{
			Value:                       denary64.Float(1, 6),
			Expected: denary64.Something(denary64.Float(1, 6)),
		},
		{
			Value:                       denary64.Float(1, 7),
			Expected: denary64.Something(denary64.Float(1, 7)),
		},
		{
			Value:                       denary64.Float(1, 8),
			Expected: denary64.Something(denary64.Float(1, 8)),
		},
		{
			Value:                       denary64.Float(1, 9),
			Expected: denary64.Something(denary64.Float(1, 9)),
		},
		{
			Value:                       denary64.Float(1, 10),
			Expected: denary64.Something(denary64.Float(1, 10)),
		},
		{
			Value:                       denary64.Float(1, 11),
			Expected: denary64.Something(denary64.Float(1, 11)),
		},
		{
			Value:                       denary64.Float(1, 12),
			Expected: denary64.Something(denary64.Float(1, 12)),
		},
		{
			Value:                       denary64.Float(1, 13),
			Expected: denary64.Something(denary64.Float(1, 13)),
		},









		{
			Value:                       denary64.Float(700, 0),
			Expected: denary64.Something(denary64.Float(700, 0)),
		},
		{
			Value:                       denary64.Float(700, 1),
			Expected: denary64.Something(denary64.Float(700, 1)),
		},
		{
			Value:                       denary64.Float(700, 2),
			Expected: denary64.Something(denary64.Float(700, 2)),
		},
		{
			Value:                       denary64.Float(700, 3),
			Expected: denary64.Something(denary64.Float(700, 3)),
		},
		{
			Value:                       denary64.Float(700, 4),
			Expected: denary64.Something(denary64.Float(700, 4)),
		},
		{
			Value:                       denary64.Float(700, 5),
			Expected: denary64.Something(denary64.Float(700, 5)),
		},
		{
			Value:                       denary64.Float(700, 6),
			Expected: denary64.Something(denary64.Float(700, 6)),
		},
		{
			Value:                       denary64.Float(700, 7),
			Expected: denary64.Something(denary64.Float(700, 7)),
		},
		{
			Value:                       denary64.Float(700, 8),
			Expected: denary64.Something(denary64.Float(700, 8)),
		},
		{
			Value:                       denary64.Float(700, 9),
			Expected: denary64.Something(denary64.Float(700, 9)),
		},
		{
			Value:                       denary64.Float(700, 10),
			Expected: denary64.Something(denary64.Float(700, 10)),
		},
		{
			Value:                       denary64.Float(700, 11),
			Expected: denary64.Something(denary64.Float(700, 11)),
		},
		{
			Value:                       denary64.Float(700, 12),
			Expected: denary64.Something(denary64.Float(700, 12)),
		},
		{
			Value:                       denary64.Float(700, 13),
			Expected: denary64.Something(denary64.Float(700, 13)),
		},









		{
			Value:                       denary64.Float(40300200010000, 401),
			Expected: denary64.Something(denary64.Float(4030020001,     397)),
		},



		{
			Value:                       denary64.Float(200000000003, 1),
			Expected: denary64.Something(denary64.Float(200000000003, 1)),
		},
	}

	for testNumber, test := range tests {

		if expected, actual := test.Expected, test.Value.Result(); expected != actual {
			t.Errorf("For test #%d, expected %#v, but actually got %#v", testNumber, expected, actual)
			continue
		}

	}
}
