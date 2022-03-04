package helper

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

func (err ErrorResponse) Error() string {
	var errorBody ErrorBody
	return fmt.Sprintf("%v", errorBody)
}

func ErrorArrayToError(errorBody []validator.FieldError) error {
	var errorResponse ErrorResponse
	errorResponse.TimeStamp = time.Now().Format(time.RFC3339)
	errorResponse.ErrorReference = uuid.New()

	for _, value := range errorBody {
		body := ErrorBody{
			Message: value.Error(),
			Code:    "400{validation} error",
			Source:  Config.AppName,
		}
		errorResponse.Errors = append(errorResponse.Errors, body)
	}
	return errorResponse
}

func PrintErrorMessage(code, message string) error {
	var errorResponse ErrorResponse
	errorResponse.TimeStamp = time.Now().Format(time.RFC3339)
	errorResponse.ErrorReference = uuid.New()
	errorResponse.Errors = append(errorResponse.Errors, ErrorBody{
		Code:    code,
		Message: message,
		Source:  Config.AppName,
	})
	return errorResponse
}
