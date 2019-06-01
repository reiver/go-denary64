package denary64

// Undefined returns a denary64.Type with an undefined value.
//
// Note that an uninitialized denary64.Type starts with a value of undefined.
//
// I.e.,...
//
//	var value denary64.Type // Nothing is assigned to this.
//
//	if denary64.Undefined() == value {
//		fmt.Println("‘value’ is undefined.")
//	} else {
//		fmt.Printf("value = %s\n", value)
//	}
//
//	// Output:
//	// ‘value’ is undefined.
func Undefined() Type {
	return Type{}
}
