package denary64

// Something returns a denary64.Result that contains a base-10 floating point number.
//
// Recall that a denary64.Result is a ‘polymorphic type’ that can contain either nothing, an error, or a base-10 floating point number.
//
// denary64.Something() gives you a denary64.Result that contains a base-10 floating point number.
//
// Here is an example usage of it:
//
//	var value denary64.Type = denary64.Uint64(123)
//	
//	var result denary64.Result = denary64.Something(value)
func Something(value Type) Result {
	return Result {
		loaded: true,
		value:  value,
	}
}
