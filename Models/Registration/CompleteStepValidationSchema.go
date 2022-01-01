package Registration

import "TestProject/Models/Base"

type CompleteStepValidationSchemaType struct {
	Sex         []Base.Validator `json:"sex"`
	DateOfBirth []Base.Validator `json:"date_of_birth"`
	Weight      []Base.Validator `json:"weight"`
	Height      []Base.Validator `json:"height"`
}

var CompleteStepValidationSchema = CompleteStepValidationSchemaType{
	Sex: []Base.Validator{
		Base.RequiredValidator(),
	},
	DateOfBirth: []Base.Validator{
		Base.RequiredValidator(),
		Base.MinAgeValidator(18),
		Base.MaxAgeValidator(80),
	},
	Weight: []Base.Validator{
		Base.RequiredValidator(),
		Base.MinValueValidator(20),
		Base.MaxValueValidator(150),
	},
	Height: []Base.Validator{
		Base.RequiredValidator(),
		Base.MinValueValidator(120),
		Base.MaxValueValidator(250),
	},
}
