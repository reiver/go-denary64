package denary64

// add does a uint64 addition (i.e., +) but also detects overflow.
func add(left uint64, right uint64) (uint64, error) {

	result := left + right

	if (result < left && right > 0) || (result < right && left > 0) {
//@TODO: return a typed error! so that it can be detected with a Golang type switch.
		return result, overflowf("%d + %d", left, right)
	}

	return result, nil
}
