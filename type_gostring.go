package denary64

import (
	"bytes"
	"fmt"
)

// GoString makes denary64.Type fit the fmt.GoStringer interface.
//
// This gets used in funcs such as: fmt.Printf(), fmt.Sprintf(), fmt.Fprintf(), fmt.Errorf(), etc when the %#v verb is used.
//
// This outputs Go syntax (in the returned string) for a denary64.Type.
//
// For example:
//
//	var x denary64.Type = denary64.Uint64(123)
//	
//	fmt.Printf("x = %#v")
//
// Will print:
//
//	x = denary64.Uint64(123)
//
// And, for example:
//	var x denary64.Type = denary64.Uint64(123)
//	
//	var s string = x.GoString()
//	
//	fmt.Println(s)
//
// Will return:
//
//	"denary64.Uint64(123)"
//
// And, for example:
//	var x denary64.Type = denary64.Float(123,2) // = 1.23
//	
//	var s string = x.GoString()
//	
//	fmt.Println(s)
//
// Will return:
//
//	"denary64.Float(123,2)"
func (receiver Type) GoString() string {

	var buffer bytes.Buffer

	switch receiver.div10 {
	default:
		buffer.WriteString("denary64.Float(")
		fmt.Fprintf(&buffer, "%d", receiver.mantissa)
		buffer.WriteString(", ")
		fmt.Fprintf(&buffer, "%d", receiver.div10)
		buffer.WriteRune(')')
	case 0:
		buffer.WriteString("denary64.Uint64(")
		fmt.Fprintf(&buffer, "%d", receiver.mantissa)
		buffer.WriteRune(')')
	}

	return buffer.String()
}
