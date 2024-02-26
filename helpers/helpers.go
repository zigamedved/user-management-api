package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(data any) error {
	if err := validate.Struct(data); err != nil {
		return err
	}

	return nil
}

func RequestToStruct(request *http.Request, result any) error {
	defer request.Body.Close()
	if err := json.NewDecoder(request.Body).Decode(result); err != nil {
		return err
	}

	return nil
}

func GetValidator() *validator.Validate {
	return validate
}
