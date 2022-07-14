package Registration

type BaseRegistrationResponse struct {
	Step             string      `json:"step"`
	Fields           interface{} `json:"fields"`
	ValidationSchema interface{} `json:"validation_schema"`
}
