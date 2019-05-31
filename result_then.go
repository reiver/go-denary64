package denary64

// Then applies the function fn to the value inside of the denary64.Result if it is a denary64.Some().
func (receiver Result) Then(fn func(Type)Result) Result {
	if None() == receiver {
		return receiver
	}
	if err := receiver.err; nil != err {
		return receiver
	}

	return fn(receiver.value)
}
