package denary64

// Nothing returns a denary64.Result that contains ‘nothing’.
//
// Recall that a denary64.Result is a ‘polymorphic type’ that can contain either nothing, an error, or a base-10 floating point number.
//
// denary64.Nothing() gives you a denary64.Result that contains ‘nothing’.
//
// Here is an example usage of it:
//
//	var result denary64.Result = denary64.Nothing()
//
// Here is another example usage of it:
//
//	var result denary64.Result
//	
//	//...
//	
//	if denary64.Nothing() == result {
//		// ...
//	}
//
// Note that an denary64.Result that has not had a value put into is has ‘nothing’ in it.
func Nothing() Result {
	return Result{}
}
