package denary64_test

import (
	"github.com/reiver/go-denary64"

	"fmt"
)

func ExampleType_String() {

	var value denary64.Type = denary64.Uint64(123)

	fmt.Println(value) // This internally calls value.String()

	fmt.Printf("value = %s", value) // This also internally calls value.String()

	// Output:
	// 123
	// value = 123
}
