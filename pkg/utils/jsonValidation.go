package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator"
	"githum.com/athunlal/bookNowAdmin-svc/pkg/domain"
)

func ValidateUser(user domain.Admin) error {
	validate := validator.New()

	err := validate.Struct(user)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make([]string, len(validationErrors))

		for i, validationErr := range validationErrors {
			fieldName := validationErr.Field()
			switch fieldName {
			case "Email":
				errorMessages[i] = "Invalid Email"
				break
			case "Username":
				errorMessages[i] = "Invalid Username, Minimum 8 letters or Maximum 24 letters required"
				break
			case "Password":
				errorMessages[i] = "Invalid password, Minimum 8 letters or Maximum 16 letters required"
				break
			case "Phone":
				errorMessages[i] = "Invalid Phone Number"
			default:
				errorMessages[i] = "Validation failed"
			}
		}

		return fmt.Errorf(strings.Join(errorMessages, ", "))
	}

	return nil
}
