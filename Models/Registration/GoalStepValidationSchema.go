package Registration

import "TestProject/Models/Base"

type GoalStepValidationSchemaType struct {
	Goal []Base.Validator `json:"goal"`
}

var GoalStepValidationSchema = GoalStepValidationSchemaType{
	Goal: []Base.Validator{
		Base.RequiredValidator(),
		Base.EnumValidator([]string{
			"IMPROVE_SHAPE",
			"LEAN_TONE",
			"LOSE_FAT",
		}),
	},
}
