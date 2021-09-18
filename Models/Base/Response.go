package Base

type Response struct {
	Errors []string    `json:"errors"`
	Data   interface{} `json:"data"`
}
