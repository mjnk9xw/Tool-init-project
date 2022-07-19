package jsonhelper

import "encoding/json"

func ConvertJson[T any](req T) string {
	databytes, _ := json.Marshal(req)
	return string(databytes)
}
