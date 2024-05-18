package pkg

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func ExtractValidationErrors(error interface{}) []string {
	var errMessage []string
	errors := error.(validator.ValidationErrors)

	for _, err := range errors {
		if err.ActualTag() == "required" {
			errMessage = append(errMessage, fmt.Sprintf("field %v is required", err.Field()))
		} else {
			errMessage = append(errMessage, fmt.Sprintf("field %v is not valid", err.Field()))
		}
	}

	return errMessage
}
