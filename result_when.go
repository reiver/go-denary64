package denary64

// WhenNone with call the func fn when the denary64.Result is none.
func (receiver Result) WhenNone(fn func()) {
	if None() == receiver {
		fn()
	}
}

// WhenNone with call the func fn when the denary64.Result is an error.
func (receiver Result) WhenError(fn func(error)) {
	if None() != receiver && nil != receiver.err {
		fn(receiver.err)
	}
}

// WhenNone with call the func fn when the denary64.Result is some.
func (receiver Result) WhenSome(fn func(Type)) {
	if None() != receiver && nil== receiver.err {
		fn(receiver.value)
	}
}
