package denary64

// A Type is a base-10 floating point number, and is safe to use to store money values, and do money calculations with them.
//
// Note that this is as opposed to the built-in Golang types float32, and float64 which are base-2 floating point number types (rather than base-10) and which are NOT safe to use for money values.
//
// An uninitialized denary64.Type has a value of zero (0).
//
// I.e.,:...
//
//	var value denary64.Type // Start off with a value of zero (0).
type Type struct {
	mantissa uint64
	div10    uint64
}
