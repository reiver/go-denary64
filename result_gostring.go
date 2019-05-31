package denary64

import (
	"bytes"
	"fmt"
)

// GoString makes denary64.Result fit the fmt.GoStringer interface.
//
// This gets used in funcs such as: fmt.Printf(), fmt.Sprintf(), fmt.Fprintf(), fmt.Errorf(), etc when the %#v verb is used.
//
// This outputs Go syntax (in the returned string) for a denary64.Result.
//
// For example:
//
//	var x denary64.Result = denary64.Some(denary64.Uint64(123))
//	
//	fmt.Printf("x = %#v")
//
// Will print:
//
//	x = denary64.Some(denary64.Uint64(123))
//
// And, for example:
//	var x denary64.Result = denary64.Some(denary64.Uint64(123))
//	
//	var s string = x.GoString()
//	
//	fmt.Println(s)
//
// Will return:
//
//	"denary64.Some((denary64.Uint64(123)))"
//
// And, for example:
//	var x denary64.Type = denary64.Some(denary64.ShiftRight(123,2)) // = 1.23
//	
//	var s string = x.GoString()
//	
//	fmt.Println(s)
//
// Will return:
//
//	"denary64.Some(denary64.ShiftRight(123,2))"
func (receiver Result) GoString() string {

	var buffer bytes.Buffer

	switch receiver {
	case Nothing():
		buffer.WriteString("denary64.Nothing()")
	default:
		switch receiver.err {
		default:
			buffer.WriteString("denary64.Error(")
			switch casted := receiver.err.(type) {
			case interface{GoString()string}:
				buffer.WriteString(casted.GoString())
			default:
				fmt.Fprintf(&buffer, "%q", receiver.err.Error())
			}
			buffer.WriteRune(')')
		case nil:
			buffer.WriteString("denary64.Some(")
			buffer.WriteString(receiver.value.GoString())
			buffer.WriteRune(')')
		}
	}

	return buffer.String()
}
