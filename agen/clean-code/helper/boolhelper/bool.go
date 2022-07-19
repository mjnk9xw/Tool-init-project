package boolhelper

func ConvertBool2I(b bool) int {
	if b {
		return 1
	}
	return 0
}

func ConvertI2Bool(i int) bool {
	if i == 1 {
		return true
	}
	return false
}
