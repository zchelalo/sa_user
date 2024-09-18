package util

func IsErrorType(err error, errArray []error) bool {
	for _, e := range errArray {
		if err == e {
			return true
		}
	}
	return false
}
