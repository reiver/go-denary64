package denary64

import (
	"bytes"
	"fmt"
)

// Overflow represents when the result of an arithmetic calculation's result is larger than the maximum value of a denary64.Type.
//
// For example:
//
//	var a denary64.Type = denary64.Uint64(18446744073709551615)
//	var b denary64.Type = denary64.Uint64(18446744073709551615)
//	
//	result := a.Add(b) // This is overflow.
//	
//	value, err := result.Return()
//	switch err.(type) {
//	case denary64.Overflow:
//		//@TODO: Deal with overflow here.
//	default:
//		// ...
//	}
type Overflow interface {
	error
	Overflow()
}

func overflowf(format string, a ...interface{}) error {
	var buffer bytes.Buffer

	buffer.WriteString("denary64: overflow; ")

	fmt.Fprintf(&buffer, format, a...)

	return internalOverflow{
		msg: buffer.String(),
	}
}

type internalOverflow struct {
	msg string
}

func (receiver internalOverflow) Error() string {
	return receiver.msg
}

func (internalOverflow) Overflow() {
	// Nothing here.
}
