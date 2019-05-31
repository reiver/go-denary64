package denary64

// Result returns a denary64.Result with the value of this denary64.Type inside of it.
func (receiver Type) Result() Result {
	return Something(receiver)
}
