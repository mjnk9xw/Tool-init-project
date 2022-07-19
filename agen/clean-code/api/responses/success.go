package responses

type SuccessRes struct {
	Data interface{} `json:"data"`
	ErrorRes
}
