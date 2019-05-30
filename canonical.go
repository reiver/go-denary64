package denary64

// canonical takes a denary64.Type, that may or may not be in canonical form, and returns a denary64.Type in canonical form.
//
// Basically, it removed extra zeros at the end.
//
// For example, it turns:
//
//	1020.3040
//
// Into:
//
//	1020.304
//
// And turns:
//
//	0.0701000
//
// Into:
//
//	0.0701
func canonical(value Type) Type {
	if 0 == value.div10 {
		return value
	}

	{
		var result Type = value

		for {
			if 0 == result.div10 {
				break
			}

			mod := result.mantissa % 10
			if 0 != mod {
				break
			}

			result.mantissa /= 10
			result.div10--
		}

		return result
	}
}
