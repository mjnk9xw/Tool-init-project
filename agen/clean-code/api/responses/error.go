package responses

type ErrorRes struct {
	Code    int    `json:"code"`
	Messgae string `json:"message"`
}
