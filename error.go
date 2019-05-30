package denary64

import (
	"errors"
	"fmt"
)

// Err returns a denary64.Result that contains an error.
//
// Recall that a denary64.Result is a ‘polymorphic type’ that can contain either nothing, an error, or a base-10 floating point number.
//
// denary64.Err() gives you a denary64.Result that contains an error.
func Err(err error) Result {
	if nil == err {
		return None()
	}

	return Result {
		loaded: true,
		err:    err,
	}
}

// Error returns a denary64.Result that contains an error.
//
// Recall that a denary64.Result is a ‘polymorphic type’ that can contain either nothing, an error, or a base-10 floating point number.
//
// denary64.Error() gives you a denary64.Result that contains an error.
//
// denary64.Error() creates a Golang error from the string passed to it, (internally) using the errors.New() func.
//
// This is a convenience func for not having to do something such as:
//
//	result := denary64.Err( errors.New(msg) )
//
// And instead just doing:
//
//	result := denary64.Error(msg)
func Error(msg string) Result {
	err := errors.New(msg)

	return Err(err)
}

// Errorf returns a denary64.Result that contains an error.
//
// Recall that a denary64.Result is a ‘polymorphic type’ that can contain either nothing, an error, or a base-10 floating point number.
//
// denary64.Errorf() gives you a denary64.Result that contains an error.
//
// denary64.Errorf() creates a Golang error from the parameters passed to it, (internally) using the fmt.Errorf() func.
//
// This is a convenience func for not having to do something such as:
//
//	result := denary64.Err( fmt.Errorf(format, a...) )
//
// And instead just doing:
//
//	result := denary64.Errorf(format, a...)
func Errorf(format string, a ...interface{}) Result {
	err := fmt.Errorf(format, a...)

	return Err(err)
}
