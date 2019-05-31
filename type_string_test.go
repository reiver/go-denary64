package denary64

import (
	"testing"
)

func TestTypeString(t *testing.T) {

	tests := []struct{
		Value    Type
		Expected string
	}{
		{
			Value: Uint64(0),
			Expected:    "0",
		},

		{
			Value:     Float(1, 5),
			Expected: "0.00001",
		},

		{
			Value:     Float(7, 5),
			Expected: "0.00007",
		},

		{
			Value:    Float(1, 4),
			Expected: "0.0001",
		},

		{
			Value:    Float(7, 4),
			Expected: "0.0007",
		},

		{
			Value:   Float(1, 3),
			Expected: "0.001",
		},

		{
			Value:   Float(7, 3),
			Expected: "0.007",
		},

		{
			Value:  Float(1, 2),
			Expected: "0.01",
		},

		{
			Value:  Float(7, 2),
			Expected: "0.07",
		},

		{
			Value: Float(1, 1),
			Expected: "0.1",
		},

		{
			Value: Float(7, 1),
			Expected: "0.7",
		},

		{
			Value: Uint64(1),
			Expected:    "1",
		},

		{
			Value: Uint64(2),
			Expected:    "2",
		},

		{
			Value: Uint64(3),
			Expected:    "3",
		},

		{
			Value: Float(302001, 5),
			Expected:  "3.02001",
		},

		{
			Value: Uint64(4),
			Expected:    "4",
		},

		{
			Value: Uint64(5),
			Expected:    "5",
		},

		{
			Value: Uint64(6),
			Expected:    "6",
		},

		{
			Value: Uint64(7),
			Expected:    "7",
		},

		{
			Value: Uint64(8),
			Expected:    "8",
		},

		{
			Value: Uint64(9),
			Expected:    "9",
		},

		{
			Value: Uint64(10),
			Expected:    "10",
		},

		{
			Value: Uint64(11),
			Expected:    "11",
		},

		{
			Value: Uint64(12),
			Expected:    "12",
		},

		{
			Value: Uint64(13),
			Expected:    "13",
		},

		{
			Value: Uint64(14),
			Expected:    "14",
		},

		{
			Value: Uint64(15),
			Expected:    "15",
		},

		{
			Value: Uint64(16),
			Expected:    "16",
		},

		{
			Value: Uint64(17),
			Expected:    "17",
		},

		{
			Value: Uint64(18),
			Expected:    "18",
		},

		{
			Value: Uint64(19),
			Expected:    "19",
		},

		{
			Value: Uint64(20),
			Expected:    "20",
		},

		{
			Value: Uint64(21),
			Expected:    "21",
		},

		{
			Value: Uint64(22),
			Expected:    "22",
		},

		{
			Value: Uint64(23),
			Expected:    "23",
		},

		{
			Value: Uint64(24),
			Expected:    "24",
		},

		{
			Value: Uint64(25),
			Expected:    "25",
		},

		{
			Value: Uint64(30),
			Expected:    "30",
		},

		{
			Value: Uint64(40),
			Expected:    "40",
		},

		{
			Value: Uint64(50),
			Expected:    "50",
		},

		{
			Value: Uint64(100),
			Expected:    "100",
		},

		{
			Value: Uint64(102),
			Expected:    "102",
		},

		{
			Value: Uint64(200),
			Expected:    "200",
		},

		{
			Value: Uint64(203),
			Expected:    "203",
		},

		{
			Value: Uint64(1000),
			Expected:    "1000",
		},

		{
			Value: Uint64(2000),
			Expected:    "2000",
		},

		{
			Value: Uint64(2005),
			Expected:    "2005",
		},

		{
			Value: Uint64(2040),
			Expected:    "2040",
		},

		{
			Value: Float(20400301, 4),
			Expected:   "2040.0301",
		},

		{
			Value: Uint64(2300),
			Expected:    "2300",
		},

		{
			Value: Float(23000004, 4),
			Expected:   "2300.0004",
		},

		{
			Value: Float(2300003, 3),
			Expected:   "2300.003",
		},

		{
			Value: Float(230002, 2),
			Expected:   "2300.02",
		},

		{
			Value: Float(23001, 1),
			Expected:   "2300.1",
		},

		{
			Value: Uint64(4030020001),
			Expected:    "4030020001",
		},

	}


	for testNumber, test := range tests {

		if expected, actual := test.Expected, test.Value.String(); expected != actual {
			t.Errorf("For test #%d...", testNumber)
			t.Errorf("\tEXPECTED: %q", expected)
			t.Errorf("\tACTUAL:   %q", actual)
			continue
		}
	}
}
