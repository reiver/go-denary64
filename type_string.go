package denary64

import (
	"bytes"
	"fmt"
)

// String makes denary64.Type fit the fmt.Stinger interface.
//
// This gets used in funcs such as: fmt.Printf(), fmt.Sprintf(), fmt.Fprintf(), fmt.Errorf(), etc when the %s verb is used.
//
// For example:
//
//	var value denary64.Type
//	
//	// ...
//	
//	fmt.Printf("value = %s", value)
//
// But is also used in funcs such as: fmt.Print(), fmt.Println(), fmt.Sprint(), fmt.Sprintln(), etc.
//
// For example:
//
//	var value denary64.Type
//	
//	// ...
//	
//	fmt.Println(value)
func (receiver Type) String() string {

	var buffer bytes.Buffer

	{
		var n uint64 = receiver.mantissa
		var stack []uint64

		for i:=uint64(0); i<receiver.div10; i++ {
			stack = append([]uint64{ n % 10 }, stack...)
			n /= 10
		}

		fmt.Fprintf(&buffer, "%d", n)

		if n != receiver.mantissa {
			buffer.WriteRune('.')
		}

		for _, value := range stack {
			fmt.Fprintf(&buffer, "%d", value)
		}
	}

	return buffer.String()
}
