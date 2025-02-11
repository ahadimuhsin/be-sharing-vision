package validator

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
    snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
    snake  = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
    return strings.ToLower(snake)
}

// ValidateStruct returns validation errors as a map of field names to error messages
func ValidateStruct(i interface{}) (map[string][]string, error) {
	errorsMap := make(map[string][]string)
	err := validate.Struct(i)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			var message string

			switch err.Tag() {
			case "email":
				message = "Invalid Email Format"
			case "required":
				message = fmt.Sprintf("%s is required", field)
			case "min":
				if field == "Password" {
					message = "Password must be at least 8 characters"
				} else {
					message = fmt.Sprintf("%s must be at least %s characters", field, err.Param())
				}
			case "oneof": // Enum validation like Laravel's "in:"
				message = fmt.Sprintf("%s must be one of %s", field, err.Param())
			default:
				message = fmt.Sprintf("Invalid value for %s", field)
			}

			// Append error message to the field's slice
			errorsMap[ToSnakeCase(field)] = append(errorsMap[field], message)
		}
		return errorsMap, errors.New("Validation failed")
	}

	return nil, nil
}
