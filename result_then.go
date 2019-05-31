package denary64

// Then applies the function fn to the value inside of the denary64.Result if it is a denary64.Some().
//
// If the denary64.Result contain nothing, or and error, then Then has no effect.
func (receiver Result) Then(fn func(Type)Result) Result {
	if Nothing() == receiver {
		return receiver
	}
	if err := receiver.err; nil != err {
		return receiver
	}

	return fn(receiver.value)
}
