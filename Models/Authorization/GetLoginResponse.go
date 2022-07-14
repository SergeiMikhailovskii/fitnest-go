package Authorization

type GetLoginResponse struct {
	Fields           GetLoginFields            `json:"fields"`
	ValidationSchema LoginValidationSchemaType `json:"validation_schema"`
}
