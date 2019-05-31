package denary64

import (
	"testing"
)

func TestTypeFloor(t *testing.T) {

	tests := []struct{
		Value    Type
		Expected Type
	}{
		{
			Value:    Uint64(0),
			Expected: Uint64(0),
		},
		{
			Value:    Uint64(1),
			Expected: Uint64(1),
		},
		{
			Value:    Uint64(2),
			Expected: Uint64(2),
		},
		{
			Value:    Uint64(3),
			Expected: Uint64(3),
		},
		{
			Value:    Uint64(4),
			Expected: Uint64(4),
		},
		{
			Value:    Uint64(5),
			Expected: Uint64(5),
		},
		{
			Value:    Uint64(6),
			Expected: Uint64(6),
		},
		{
			Value:    Uint64(7),
			Expected: Uint64(7),
		},
		{
			Value:    Uint64(8),
			Expected: Uint64(8),
		},
		{
			Value:    Uint64(9),
			Expected: Uint64(9),
		},
		{
			Value:    Uint64(10),
			Expected: Uint64(10),
		},
		{
			Value:    Uint64(11),
			Expected: Uint64(11),
		},
		{
			Value:    Uint64(12),
			Expected: Uint64(12),
		},
		{
			Value:    Uint64(13),
			Expected: Uint64(13),
		},









		{
			Value:    Float(1,1),
			Expected: Uint64(0),
		},
		{
			Value:    Float(1,2),
			Expected: Uint64(0),
		},
		{
			Value:    Float(1,3),
			Expected: Uint64(0),
		},
		{
			Value:    Float(1,4),
			Expected: Uint64(0),
		},
		{
			Value:    Float(1,5),
			Expected: Uint64(0),
		},
		{
			Value:    Float(1,6),
			Expected: Uint64(0),
		},
		{
			Value:    Float(1,7),
			Expected: Uint64(0),
		},
		{
			Value:    Float(1,8),
			Expected: Uint64(0),
		},
		{
			Value:    Float(1,9),
			Expected: Uint64(0),
		},
		{
			Value:    Float(1,10),
			Expected: Uint64(0),
		},
		{
			Value:    Float(1,11),
			Expected: Uint64(0),
		},
		{
			Value:    Float(1,12),
			Expected: Uint64(0),
		},
		{
			Value:    Float(1,13),
			Expected: Uint64(0),
		},









		{
			Value:    Float(11,1),
			Expected: Uint64(1),
		},
		{
			Value:    Float(101,2),
			Expected: Uint64(1),
		},
		{
			Value:    Float(1001,3),
			Expected: Uint64(1),
		},
		{
			Value:    Float(10001,4),
			Expected: Uint64(1),
		},
		{
			Value:    Float(100001,5),
			Expected: Uint64(1),
		},
		{
			Value:    Float(1000001,6),
			Expected: Uint64(1),
		},
		{
			Value:    Float(10000001,7),
			Expected: Uint64(1),
		},
		{
			Value:    Float(100000001,8),
			Expected: Uint64(1),
		},
		{
			Value:    Float(1000000001,9),
			Expected: Uint64(1),
		},
		{
			Value:    Float(10000000001,10),
			Expected: Uint64(1),
		},
		{
			Value:    Float(100000000001,11),
			Expected: Uint64(1),
		},
		{
			Value:    Float(1000000000001,12),
			Expected: Uint64(1),
		},
		{
			Value:    Float(10000000000001,13),
			Expected: Uint64(1),
		},









		{
			Value:    Float(129,0),
			Expected: Uint64(129),
		},
		{
			Value:    Float(1291,1),
			Expected: Uint64(129),
		},
		{
			Value:    Float(12901,2),
			Expected: Uint64(129),
		},
		{
			Value:    Float(129001,3),
			Expected: Uint64(129),
		},
		{
			Value:    Float(1290001,4),
			Expected: Uint64(129),
		},
		{
			Value:    Float(12900001,5),
			Expected: Uint64(129),
		},
		{
			Value:    Float(129000001,6),
			Expected: Uint64(129),
		},
		{
			Value:    Float(1290000001,7),
			Expected: Uint64(129),
		},
		{
			Value:    Float(12900000001,8),
			Expected: Uint64(129),
		},
		{
			Value:    Float(129000000001,9),
			Expected: Uint64(129),
		},
		{
			Value:    Float(1290000000001,10),
			Expected: Uint64(129),
		},
		{
			Value:    Float(12900000000001,11),
			Expected: Uint64(129),
		},
		{
			Value:    Float(129000000000001,12),
			Expected: Uint64(129),
		},
		{
			Value:    Float(1290000000000001,13),
			Expected: Uint64(129),
		},









		{
			Value:    Float(987,0),
			Expected: Uint64(987),
		},
		{
			Value:    Float(9876,1),
			Expected: Uint64(987),
		},
		{
			Value:    Float(98765,2),
			Expected: Uint64(987),
		},
		{
			Value:    Float(987654,3),
			Expected: Uint64(987),
		},
		{
			Value:    Float(9876543,4),
			Expected: Uint64(987),
		},
		{
			Value:    Float(98765432,5),
			Expected: Uint64(987),
		},
		{
			Value:    Float(987654321,6),
			Expected: Uint64(987),
		},
		{
			Value:    Float(9876543210,7),
			Expected: Uint64(987),
		},
		{
			Value:    Float(98765432109,8),
			Expected: Uint64(987),
		},
		{
			Value:    Float(987654321098,9),
			Expected: Uint64(987),
		},
		{
			Value:    Float(9876543210987,10),
			Expected: Uint64(987),
		},
		{
			Value:    Float(98765432109876,11),
			Expected: Uint64(987),
		},
		{
			Value:    Float(987654321098765,12),
			Expected: Uint64(987),
		},
		{
			Value:    Float(9876543210987654,13),
			Expected: Uint64(987),
		},
	}

	for testNumber, test := range tests {

		if expected, actual := test.Expected, test.Value.Floor(); expected != actual {
			t.Errorf("For test #%d...", testNumber)
			t.Errorf("\tVALUE:    %s", test.Value)
			t.Errorf("\tEXPECTED: %s", expected)
			t.Errorf("\tACTUAL:   %s", actual)
			continue
		}
	}
}
