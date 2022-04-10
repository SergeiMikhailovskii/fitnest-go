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

func MinAgeValidator(age int) Validator {
	return Validator{
		Type:       "minAge",
		Error:      "error.minAge",
		Validation: age,
	}
}

func MaxAgeValidator(age int) Validator {
	return Validator{
		Type:       "maxAge",
		Error:      "error.maxAge",
		Validation: age,
	}
}

func MinValueValidator(value int) Validator {
	return Validator{
		Type:       "minValue",
		Error:      "error.minValue",
		Validation: value,
	}
}

func MaxValueValidator(value int) Validator {
	return Validator{
		Type:       "maxValue",
		Error:      "error.maxValue",
		Validation: value,
	}
}

func EnumValidator(value []string) Validator {
	return Validator{
		Type:       "enum",
		Error:      "error.invalid",
		Validation: value,
	}
}

func OnlyLettersValidator() Validator {
	return RegExpValidator("^[a-zA-Z]+$")
}

func EmailValidator() Validator {
	return RegExpValidator(".{5,}@.{2,}\\..{2,}")
}
