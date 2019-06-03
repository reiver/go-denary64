package denary64

// Return more or less turns a denary.Result into a denary.Type.
func (receiver Result) Return() (Type, error) {
	switch {
	case ! receiver.loaded:
		return Undefined(), nil
	case nil != receiver.err:
		return receiver.value, receiver.err
	default:
		return receiver.value, nil
	}
}
