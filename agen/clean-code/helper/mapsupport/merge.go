package mapsupport

type KMap interface {
	string | float64 | int64 | int32 | int | float32
}

func Merge[K KMap, V any](one, second map[K]V) map[K]V {

	for k, v := range second {
		one[k] = v
	}
	return one
}
