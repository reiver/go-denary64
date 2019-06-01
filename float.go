package denary64

// Float, is a ‘right shirt’, and returns a denary64.Type with the value: [mantissa] × 10⁻ˢʰⁱᶠᵗ
func Float(mantissa uint64, shift uint64) Type {
	return canonical(Type{
		loaded:   true,
		mantissa: mantissa,
		div10:    shift,
	})
}
