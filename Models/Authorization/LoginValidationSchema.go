package Authorization

import "TestProject/Models/Base"

type LoginValidationSchemaType struct {
	Login    []Base.Validator `json:"login"`
	Password []Base.Validator `json:"password"`
}

var LoginValidationSchema = LoginValidationSchemaType{
	Login: []Base.Validator{
		Base.RequiredValidator(),
	},
	Password: []Base.Validator{
		Base.RequiredValidator(),
	},
}
