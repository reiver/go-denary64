package denary64

// A Type is a base-10 floating point number, and is safe to use to store money values, and do money calculations with them.
//
// Note that this is as opposed to the built-in Golang types float32, and float64 which are base-2 floating point number types (rather than base-10) and which are NOT safe to use for money values.
//
// Undefined
//
// An uninitialized denary64.Type has a value of denary64.Undefined().
//
// Note that this does NOT mean that an uninitialized denary64.Type has some numerical value, but you won't know what it is.
// Having a value of denary64.Undefined() is a non-numerical value that a denary64.Type can have.
//
// You can understand denary64.Type as being something like what other programming languages might call an ‘option type’.
//
// Assignment
//
// If you want to give a denary64.Type a ‘whole number’ value, then you can do that using the denary64.Uint() func:
//
//      var value denary64.Type = denary64.Uint(123)
//
// If you want to give a denary64.Type a numeric value that inclues a fractional part, then you can do that using the denary64.Float() func:
//
//	var value denay64.Type = denary64.Float(123,2) // = 123 × 10⁻² = 1.23
type Type struct {
	mantissa uint64
	div10    uint64
	loaded   bool
}
