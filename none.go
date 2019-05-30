package denary64

// None returns a denary64.Result that represents the concept of ‘nothing’.
//
// Recall that a denary64.Result is a ‘polymorphic type’ that represents either nothing, and error, or a base-10 floating point number.
//
// denary64.None() gives you that ‘nothing’ value for denary64.Result.
//
// Here is an example usage of it:
//
//	var result denary64.Result
//	
//	//...
//	
//	if denary64.None() == result {
//		// ...
//	}
func None() Result {
	return Result{}
}
