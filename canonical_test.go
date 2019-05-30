package denary64

import (
	"testing"
)

func TestCanonical(t *testing.T) {

	tests := []struct{
		Value    Type
		Expected Type
	}{
		{
			Value:    Type{mantissa: 0, div10: 0}, // 0
			Expected: Type{mantissa: 0, div10: 0}, // 0
		},
		{
			Value:    Type{mantissa: 0, div10: 1}, // 0.0
			Expected: Type{mantissa: 0, div10: 0}, // 0
		},
		{
			Value:    Type{mantissa: 0, div10: 2}, // 0.00
			Expected: Type{mantissa: 0, div10: 0}, // 0
		},
		{
			Value:    Type{mantissa: 0, div10: 3}, // 0.000
			Expected: Type{mantissa: 0, div10: 0}, // 0
		},
		{
			Value:    Type{mantissa: 0, div10: 4}, // 0.0000
			Expected: Type{mantissa: 0, div10: 0}, // 0
		},
		{
			Value:    Type{mantissa: 0, div10: 5}, // 0.00000
			Expected: Type{mantissa: 0, div10: 0}, // 0
		},
		{
			Value:    Type{mantissa: 0, div10: 6}, // 0.000000
			Expected: Type{mantissa: 0, div10: 0}, // 0
		},
		{
			Value:    Type{mantissa: 0, div10: 7}, // 0.0000000
			Expected: Type{mantissa: 0, div10: 0}, // 0
		},
		{
			Value:    Type{mantissa: 0, div10: 8}, // 0.00000000
			Expected: Type{mantissa: 0, div10: 0}, // 0
		},
		{
			Value:    Type{mantissa: 0, div10: 9}, // 0.000000000
			Expected: Type{mantissa: 0, div10: 0}, // 0
		},
		{
			Value:    Type{mantissa: 0, div10: 10}, // 0.0000000000
			Expected: Type{mantissa: 0, div10:  0}, // 0
		},
		{
			Value:    Type{mantissa: 0, div10: 11}, // 0.00000000000
			Expected: Type{mantissa: 0, div10:  0}, // 0
		},
		{
			Value:    Type{mantissa: 0, div10: 12}, // 0.000000000000
			Expected: Type{mantissa: 0, div10:  0}, // 0
		},
		{
			Value:    Type{mantissa: 0, div10: 13}, // 0.0000000000000
			Expected: Type{mantissa: 0, div10:  0}, // 0
		},









		{
			Value:    Type{mantissa: 1, div10: 0}, // 1
			Expected: Type{mantissa: 1, div10: 0}, // 1
		},
		{
			Value:    Type{mantissa: 2, div10: 0}, // 2
			Expected: Type{mantissa: 2, div10: 0}, // 2
		},
		{
			Value:    Type{mantissa: 3, div10: 0}, // 3
			Expected: Type{mantissa: 3, div10: 0}, // 3
		},
		{
			Value:    Type{mantissa: 4, div10: 0}, // 4
			Expected: Type{mantissa: 4, div10: 0}, // 4
		},
		{
			Value:    Type{mantissa: 5, div10: 0}, // 5
			Expected: Type{mantissa: 5, div10: 0}, // 5
		},
		{
			Value:    Type{mantissa: 6, div10: 0}, // 6
			Expected: Type{mantissa: 6, div10: 0}, // 6
		},
		{
			Value:    Type{mantissa: 7, div10: 0}, // 7
			Expected: Type{mantissa: 7, div10: 0}, // 7
		},
		{
			Value:    Type{mantissa: 8, div10: 0}, // 8
			Expected: Type{mantissa: 8, div10: 0}, // 8
		},
		{
			Value:    Type{mantissa: 9, div10: 0}, // 9
			Expected: Type{mantissa: 9, div10: 0}, // 9
		},
		{
			Value:    Type{mantissa: 10, div10: 0}, // 10
			Expected: Type{mantissa: 10, div10: 0}, // 10
		},
		{
			Value:    Type{mantissa: 11, div10: 0}, // 11
			Expected: Type{mantissa: 11, div10: 0}, // 11
		},
		{
			Value:    Type{mantissa: 12, div10: 0}, // 12
			Expected: Type{mantissa: 12, div10: 0}, // 12
		},
		{
			Value:    Type{mantissa: 13, div10: 0}, // 13
			Expected: Type{mantissa: 13, div10: 0}, // 13
		},









		{
			Value:    Type{mantissa: 10, div10: 1}, // 1.0
			Expected: Type{mantissa: 1,  div10: 0}, // 1
		},
		{
			Value:    Type{mantissa: 20, div10: 1}, // 2.0
			Expected: Type{mantissa: 2,  div10: 0}, // 2
		},
		{
			Value:    Type{mantissa: 30, div10: 1}, // 3.0
			Expected: Type{mantissa: 3,  div10: 0}, // 3
		},
		{
			Value:    Type{mantissa: 40, div10: 1}, // 4.0
			Expected: Type{mantissa: 4,  div10: 0}, // 4
		},
		{
			Value:    Type{mantissa: 50, div10: 1}, // 5.0
			Expected: Type{mantissa: 5,  div10: 0}, // 5
		},
		{
			Value:    Type{mantissa: 60, div10: 1}, // 6.0
			Expected: Type{mantissa: 6,  div10: 0}, // 6
		},
		{
			Value:    Type{mantissa: 70, div10: 1}, // 7.0
			Expected: Type{mantissa: 7,  div10: 0}, // 7
		},
		{
			Value:    Type{mantissa: 80, div10: 1}, // 8.0
			Expected: Type{mantissa: 8,  div10: 0}, // 8
		},
		{
			Value:    Type{mantissa: 90, div10: 1}, // 9.0
			Expected: Type{mantissa: 9,  div10: 0}, // 9
		},
		{
			Value:    Type{mantissa: 100, div10: 1}, // 10.0
			Expected: Type{mantissa: 10,  div10: 0}, // 10
		},
		{
			Value:    Type{mantissa: 110, div10: 1}, // 11.0
			Expected: Type{mantissa: 11,  div10: 0}, // 11
		},
		{
			Value:    Type{mantissa: 120, div10: 1}, // 12.0
			Expected: Type{mantissa: 12,  div10: 0}, // 12
		},
		{
			Value:    Type{mantissa: 130, div10: 1}, // 13.0
			Expected: Type{mantissa: 13,  div10: 0}, // 13
		},









		{
			Value:    Type{mantissa: 100, div10: 2}, // 1.00
			Expected: Type{mantissa: 1,   div10: 0}, // 1
		},
		{
			Value:    Type{mantissa: 200, div10: 2}, // 2.00
			Expected: Type{mantissa: 2,   div10: 0}, // 2
		},
		{
			Value:    Type{mantissa: 300, div10: 2}, // 3.00
			Expected: Type{mantissa: 3,   div10: 0}, // 3
		},
		{
			Value:    Type{mantissa: 400, div10: 2}, // 4.00
			Expected: Type{mantissa: 4,   div10: 0}, // 4
		},
		{
			Value:    Type{mantissa: 500, div10: 2}, // 5.00
			Expected: Type{mantissa: 5,   div10: 0}, // 5
		},
		{
			Value:    Type{mantissa: 600, div10: 2}, // 6.00
			Expected: Type{mantissa: 6,   div10: 0}, // 6
		},
		{
			Value:    Type{mantissa: 700, div10: 2}, // 7.00
			Expected: Type{mantissa: 7,   div10: 0}, // 7
		},
		{
			Value:    Type{mantissa: 800, div10: 2}, // 8.00
			Expected: Type{mantissa: 8,   div10: 0}, // 8
		},
		{
			Value:    Type{mantissa: 900, div10: 2}, // 9.00
			Expected: Type{mantissa: 9,   div10: 0}, // 9
		},
		{
			Value:    Type{mantissa: 1000, div10: 2}, // 10.00
			Expected: Type{mantissa: 10,   div10: 0}, // 10
		},
		{
			Value:    Type{mantissa: 1100, div10: 2}, // 11.00
			Expected: Type{mantissa: 11,   div10: 0}, // 11
		},
		{
			Value:    Type{mantissa: 1200, div10: 2}, // 12.00
			Expected: Type{mantissa: 12,   div10: 0}, // 12
		},
		{
			Value:    Type{mantissa: 1300, div10: 2}, // 13.00
			Expected: Type{mantissa: 13,   div10: 0}, // 13
		},









		{
			Value:    Type{mantissa: 10, div10: 2}, // 0.10
			Expected: Type{mantissa: 1,  div10: 1}, // 0.1
		},
		{
			Value:    Type{mantissa: 20, div10: 2}, // 0.20
			Expected: Type{mantissa: 2,  div10: 1}, // 0.2
		},
		{
			Value:    Type{mantissa: 30, div10: 2}, // 0.30
			Expected: Type{mantissa: 3,  div10: 1}, // 0.3
		},
		{
			Value:    Type{mantissa: 40, div10: 2}, // 0.40
			Expected: Type{mantissa: 4,  div10: 1}, // 0.4
		},
		{
			Value:    Type{mantissa: 50, div10: 2}, // 0.50
			Expected: Type{mantissa: 5,  div10: 1}, // 0.5
		},
		{
			Value:    Type{mantissa: 60, div10: 2}, // 0.60
			Expected: Type{mantissa: 6,  div10: 1}, // 0.6
		},
		{
			Value:    Type{mantissa: 70, div10: 2}, // 0.70
			Expected: Type{mantissa: 7,  div10: 1}, // 0.7
		},
		{
			Value:    Type{mantissa: 80, div10: 2}, // 0.80
			Expected: Type{mantissa: 8,  div10: 1}, // 0.8
		},
		{
			Value:    Type{mantissa: 90, div10: 2}, // 0.90
			Expected: Type{mantissa: 9,  div10: 1}, // 0.9
		},
		{
			Value:    Type{mantissa: 100, div10: 2}, // 1.00
			Expected: Type{mantissa: 1,   div10: 0}, // 1
		},
		{
			Value:    Type{mantissa: 110, div10: 2}, // 1.10
			Expected: Type{mantissa: 11,  div10: 1}, // 1.1
		},
		{
			Value:    Type{mantissa: 120, div10: 2}, // 1.20
			Expected: Type{mantissa: 12,  div10: 1}, // 1.2
		},
		{
			Value:    Type{mantissa: 130, div10: 2}, // 1.30
			Expected: Type{mantissa: 13,  div10: 1}, // 1.3
		},









		{
			Value:    Type{mantissa: 40300200010000, div10: 14}, // 0.40300200010000
			Expected: Type{mantissa: 4030020001,     div10: 10}, // 0.4030020001
		},









		{
			Value:    Type{mantissa: 700800000, div10: 0}, // 700800000
			Expected: Type{mantissa: 700800000, div10: 0}, // 700800000
		},
	}

	for testNumber, test := range tests {

		if expected, actual := test.Expected, canonical(test.Value); expected != actual {
			t.Errorf("For test #%d...", testNumber)
			t.Errorf("\tVALUE:    %#v", test.Value)
			t.Errorf("\tEXPECTED: %#v", expected)
			t.Errorf("\tACTUAL:   %#v", actual)
			continue
		}
	}
}
