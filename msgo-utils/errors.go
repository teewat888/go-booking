package msgoutils

import (
	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
)

type ValidationError struct {
	Field string `json:"field"`
	Rule  string `json:"rule"`
	Value any    `json:"value"`
}

func ParseValidationErrors(err error) []ValidationError {
	vErrArr := make([]ValidationError, 0)

	for _, err := range err.(validator.ValidationErrors) {
		vErr := ValidationError{
			Field: strcase.ToLowerCamel(err.StructField()),
			Rule:  err.ActualTag(),
			Value: err.Value(),
		}
		vErrArr = append(vErrArr, vErr)
	}

	return vErrArr
}
