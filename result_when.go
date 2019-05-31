package denary64

// WhenNone with call the func fn when the denary64.Result is none.
func (receiver Result) WhenNothing(fn func()) {
	if Nothing() == receiver {
		fn()
	}
}

// WhenNone with call the func fn when the denary64.Result is an error.
func (receiver Result) WhenError(fn func(error)) {
	if Nothing() != receiver && nil != receiver.err {
		fn(receiver.err)
	}
}

// WhenNone with call the func fn when the denary64.Result is some.
func (receiver Result) WhenSome(fn func(Type)) {
	if Nothing() != receiver && nil== receiver.err {
		fn(receiver.value)
	}
}
