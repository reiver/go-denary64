package denary64

// WhenNothing with call the func fn when the denary64.Result is nothing.
func (receiver Result) WhenNothing(fn func()) {
	if Nothing() == receiver {
		fn()
	}
}

// WhenError with call the func fn when the denary64.Result is an error.
func (receiver Result) WhenError(fn func(error)) {
	if Nothing() != receiver && nil != receiver.err {
		fn(receiver.err)
	}
}

// WhenSomething with call the func fn when the denary64.Result is some.
func (receiver Result) WhenSomething(fn func(Type)) {
	if Nothing() != receiver && nil== receiver.err {
		fn(receiver.value)
	}
}
