package denary64

// A Result is a ‘polymorphic type’ that represents either nothing, and error, or a base-10 floating point number.
type Result struct {
	loaded bool
	err    error
	value  Type
}
