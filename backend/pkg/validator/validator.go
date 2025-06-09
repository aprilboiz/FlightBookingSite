package validator

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *validator.Validate
}

func New() *Validator {
	v := validator.New()

	// Register field name to JSON tag mapping
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		if jsonTag := fld.Tag.Get("json"); jsonTag != "" {
			return jsonTag
		}
		return ""
	})

	return &Validator{validator: v}
}

func (v *Validator) Validate(i any) error {
	return v.validator.Struct(i)
}

func ValidationErrors(err error) map[string]string {
	if err == nil {
		return nil
	}

	errors := make(map[string]string)

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		errors["generic"] = err.Error()
		return errors
	}

	for _, e := range validationErrors {
		errors[e.Field()] = validationMessage(e)
	}

	return errors
}

func validationMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email address"
	case "url":
		return "Invalid URL"
	case "min":
		return "Minimum value is " + e.Param()
	case "max":
		return "Maximum value is " + e.Param()
	case "len":
		return "Length must be exactly " + e.Param()
	default:
		return e.Error()
	}
}
