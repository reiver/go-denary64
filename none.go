package denary64

// None returns a denary64.Result that contains (the concept of) ‘nothing’.
//
// Recall that a denary64.Result is a ‘polymorphic type’ that can contain either nothing, an error, or a base-10 floating point number.
//
// denary64.None() gives you a denary64.Result that contains (the concept of) ‘nothing’.
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
