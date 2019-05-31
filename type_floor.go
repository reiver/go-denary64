package denary64

// Floor returns the floor (i.e., whole number value) of a denary64.Type.
//
// For example:
//
//	var value denary64.Type = denary64.Parse("12.345")
//	
//	var f denary.Type = value.Floor()
//	
//	fmt.Printf(f)
//	// Output:
//	// 12
func (receiver Type) Floor() Type {
	var x Type = receiver

	for 0 < x.div10 {
		x.mantissa /= 10
		x.div10--
	}

	return x
}
