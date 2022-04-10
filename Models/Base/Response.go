package Base

type Response struct {
	Data   interface{} `json:"data"`
	Errors []Error     `json:"errors"`
	Flow   string      `json:"flow"`
}
