package denary64

// A Result is a ‘polymorphic type’ that can contain either nothing, an error, or a base-10 floating point number.
//
// A way to create a denary64.Result that represents nothing would be to either to use the denary64.Nothing() func:
//
//      var result denary64.Result = denary64.Nothing()
//
// Or just create a denary64.Result, and don't assign anything to it, as in:
//
//      var result denary.Result // <-- It starts off with the value denary64.Nothing()
//
// Note that the denary64.Nothing() func can be useful if you want to check if a denary64.Result has a value of denary64.Nothing():
//
//	var result denary64.Result
//	
//	//...
//	
//	if denary64.Nothing() == result {
//		//...
//	}
type Result struct {
	loaded bool
	err    error
	value  Type
}
