package arrhelper

type Constraints interface {
	string | float64 | int64 | int | int32 | float32 | bool
}

func CheckItemInArray[T Constraints](arr []T, item T) int {
	for i, v := range arr {
		if v == item {
			return i
		}
	}
	return -1
}
