package denary64

// Result returns a denary64.Result with the value of this denary64.Type inside of it.
func (receiver Type) Result() Result {
	switch receiver {
	case Undefined():
		return Nothing()
	default:
		return Something(receiver)
	}
}
