package Registration

type Response struct {
	Step             string      `json:"step"`
	Fields           interface{} `json:"fields"`
	ValidationSchema interface{} `json:"validation_schema"`
}
