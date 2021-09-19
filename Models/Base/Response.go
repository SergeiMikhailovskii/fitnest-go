package Base

type Response struct {
	Errors []Error     `json:"errors"`
	Data   interface{} `json:"data"`
}
