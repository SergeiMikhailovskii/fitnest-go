package Registration

type CreateStepValidationSchemaType struct {
	FirstName []Validator `json:"first_name"`
	LastName  []Validator `json:"last_name"`
	Email     []Validator `json:"email"`
	Password  []Validator `json:"password"`
}

type Validator struct {
	Type       string      `json:"type"`
	Error      string      `json:"error"`
	Validation interface{} `json:"validation"`
}

func RequiredValidator() Validator {
	return Validator{
		Type:  "required",
		Error: "error.required",
	}
}

func RegExpValidator(validation string) Validator {
	return Validator{
		Type:       "regExp",
		Error:      "error.invalid",
		Validation: validation,
	}
}

func OnlyLettersValidator() Validator {
	return RegExpValidator("^[a-zA-Z]+$")
}

func EmailValidator() Validator {
	return RegExpValidator(".{5,}@.{2,}\\..{2,}")
}

var CreateStepValidationSchema = CreateStepValidationSchemaType{
	FirstName: []Validator{
		RequiredValidator(),
		OnlyLettersValidator(),
	},
	LastName: []Validator{
		RequiredValidator(),
		OnlyLettersValidator(),
	},
	Email: []Validator{
		RequiredValidator(),
		EmailValidator(),
	},
	Password: []Validator{
		RequiredValidator(),
	},
}
