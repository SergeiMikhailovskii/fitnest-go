package Base

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

func MinLengthValidator(length int) Validator {
	return Validator{
		Type:       "minLength",
		Error:      "error.minLength",
		Validation: length,
	}
}

func OnlyLettersValidator() Validator {
	return RegExpValidator("^[a-zA-Z]+$")
}

func EmailValidator() Validator {
	return RegExpValidator(".{5,}@.{2,}\\..{2,}")
}
