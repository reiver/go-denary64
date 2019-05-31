package denary64_test

import (
	"github.com/reiver/go-denary64"

	"fmt"
	"math/rand"
	"time"

	"testing"
)

func TestTypeGoString(t *testing.T) {

	tests := []struct{
		Value    denary64.Type
		Expected string
	}{
		{
			Value:     denary64.Uint64(0),
			Expected: "denary64.Uint64(0)",
		},
		{
			Value:     denary64.Uint64(1),
			Expected: "denary64.Uint64(1)",
		},
		{
			Value:     denary64.Uint64(2),
			Expected: "denary64.Uint64(2)",
		},
		{
			Value:     denary64.Uint64(3),
			Expected: "denary64.Uint64(3)",
		},
		{
			Value:     denary64.Uint64(4),
			Expected: "denary64.Uint64(4)",
		},
		{
			Value:     denary64.Uint64(5),
			Expected: "denary64.Uint64(5)",
		},
		{
			Value:     denary64.Uint64(6),
			Expected: "denary64.Uint64(6)",
		},
		{
			Value:     denary64.Uint64(7),
			Expected: "denary64.Uint64(7)",
		},
		{
			Value:     denary64.Uint64(8),
			Expected: "denary64.Uint64(8)",
		},
		{
			Value:     denary64.Uint64(9),
			Expected: "denary64.Uint64(9)",
		},
		{
			Value:     denary64.Uint64(10),
			Expected: "denary64.Uint64(10)",
		},
		{
			Value:     denary64.Uint64(11),
			Expected: "denary64.Uint64(11)",
		},
		{
			Value:     denary64.Uint64(12),
			Expected: "denary64.Uint64(12)",
		},
		{
			Value:     denary64.Uint64(13),
			Expected: "denary64.Uint64(13)",
		},



		{
			Value:      denary64.Float(0, 0),
			Expected: "denary64.Uint64(0)",
		},
		{
			Value:      denary64.Float(1, 0),
			Expected: "denary64.Uint64(1)",
		},
		{
			Value:      denary64.Float(2, 0),
			Expected: "denary64.Uint64(2)",
		},
		{
			Value: denary64.Float(3, 0),
			Expected: "denary64.Uint64(3)",
		},
		{
			Value:      denary64.Float(4, 0),
			Expected: "denary64.Uint64(4)",
		},
		{
			Value:      denary64.Float(5, 0),
			Expected: "denary64.Uint64(5)",
		},
		{
			Value:      denary64.Float(6, 0),
			Expected: "denary64.Uint64(6)",
		},
		{
			Value:      denary64.Float(7, 0),
			Expected: "denary64.Uint64(7)",
		},
		{
			Value:      denary64.Float(8, 0),
			Expected: "denary64.Uint64(8)",
		},
		{
			Value:      denary64.Float(9, 0),
			Expected: "denary64.Uint64(9)",
		},
		{
			Value:      denary64.Float(10, 0),
			Expected: "denary64.Uint64(10)",
		},
		{
			Value:      denary64.Float(11, 0),
			Expected: "denary64.Uint64(11)",
		},
		{
			Value:      denary64.Float(12, 0),
			Expected: "denary64.Uint64(12)",
		},
		{
			Value:      denary64.Float(13, 0),
			Expected: "denary64.Uint64(13)",
		},



		{
			Value:     denary64.Float(1, 1),
			Expected: "denary64.Float(1, 1)",
		},
		{
			Value:     denary64.Float(1, 2),
			Expected: "denary64.Float(1, 2)",
		},
		{
			Value:     denary64.Float(1, 3),
			Expected: "denary64.Float(1, 3)",
		},
		{
			Value:     denary64.Float(1, 4),
			Expected: "denary64.Float(1, 4)",
		},
		{
			Value:     denary64.Float(1, 5),
			Expected: "denary64.Float(1, 5)",
		},
		{
			Value:     denary64.Float(1, 6),
			Expected: "denary64.Float(1, 6)",
		},
		{
			Value:     denary64.Float(1, 7),
			Expected: "denary64.Float(1, 7)",
		},
		{
			Value:     denary64.Float(1, 8),
			Expected: "denary64.Float(1, 8)",
		},
		{
			Value:     denary64.Float(1, 9),
			Expected: "denary64.Float(1, 9)",
		},
		{
			Value:     denary64.Float(1, 10),
			Expected: "denary64.Float(1, 10)",
		},
		{
			Value:     denary64.Float(1, 11),
			Expected: "denary64.Float(1, 11)",
		},
		{
			Value:     denary64.Float(1, 12),
			Expected: "denary64.Float(1, 12)",
		},
		{
			Value:     denary64.Float(1, 13),
			Expected: "denary64.Float(1, 13)",
		},



		{
			Value:      denary64.Float(700, 1),
			Expected: "denary64.Uint64(70)",
		},
		{
			Value:      denary64.Float(700, 2),
			Expected: "denary64.Uint64(7)",
		},
		{
			Value:     denary64.Float(700, 3),
			Expected: "denary64.Float(7, 1)",
		},
		{
			Value:     denary64.Float(700, 4),
			Expected: "denary64.Float(7, 2)",
		},
		{
			Value:     denary64.Float(700, 5),
			Expected: "denary64.Float(7, 3)",
		},
		{
			Value:     denary64.Float(700, 6),
			Expected: "denary64.Float(7, 4)",
		},
		{
			Value:     denary64.Float(700, 7),
			Expected: "denary64.Float(7, 5)",
		},
		{
			Value:     denary64.Float(700, 8),
			Expected: "denary64.Float(7, 6)",
		},
		{
			Value:     denary64.Float(700, 9),
			Expected: "denary64.Float(7, 7)",
		},
		{
			Value:     denary64.Float(700, 10),
			Expected: "denary64.Float(7, 8)",
		},
		{
			Value:     denary64.Float(700, 11),
			Expected: "denary64.Float(7, 9)",
		},
		{
			Value:     denary64.Float(700, 12),
			Expected: "denary64.Float(7, 10)",
		},
		{
			Value:     denary64.Float(700, 13),
			Expected: "denary64.Float(7, 11)",
		},



		{
			Value:     denary64.Float(40300200010000, 401),
			Expected: "denary64.Float(4030020001, 397)",
		},



		{
			Value:     denary64.Float(200000000003, 1),
			Expected: "denary64.Float(200000000003, 1)",
		},
	}

	{
		randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

		for i:=0; i<100; i++ {

			n     := randomness.Uint64()
			shift := randomness.Uint64()

			for shift > 0 && 0 == n % 10 {
				n /= 10
				shift--
			}

			switch shift {
			default:
				test := struct{
					Value    denary64.Type
					Expected string
				}{
					Value:                 denary64.Float(n, shift),
					Expected: fmt.Sprintf("denary64.Float(%d, %d)", n, shift),
				}

				tests = append(tests, test)
			case 0:

				test := struct{
					Value    denary64.Type
					Expected string
				}{
					Value:                 denary64.Float(n, shift),
					Expected: fmt.Sprintf("denary64.Uint64(%d)", n),
				}

				tests = append(tests, test)
			}
		}

		for i:=0; i<100; i++ {

			n := randomness.Uint64()

			test := struct{
				Value    denary64.Type
				Expected string
			}{
				Value:                 denary64.Float(n, 0),
				Expected: fmt.Sprintf("denary64.Uint64(%d)", n),
			}

			tests = append(tests, test)
		}
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
