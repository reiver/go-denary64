package denary64

// Uint64 returns a denary64.Type with the same logical value as the uint64 passed to it.
//
// For example:
//
//	var value denary64.Type = denary64.Uint64(123)
//	
//	fmt.Printf("value = %s", value)
//	// Output: value = 123
func Uint64(value uint64) Type {
	return Type{
		loaded:   true,
		mantissa: value,
		div10:    0,
	}
}
