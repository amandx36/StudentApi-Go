package response

import (
	"encoding/json"
	"fmt"
	"strings"

	"net/http"

	"github.com/go-playground/validator/v10"
)

// for encoding from struct to json

//   any / interface{}   both are same dude

const (
	StatusOk    = "OK"
	StatusError = "ERROR"
)

// making a custom struct for Bad request
type Response struct {
	Status string
	Error  string
}

func WriteJson(writer http.ResponseWriter, statusCode int, data any) error {

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	return json.NewEncoder(writer).Encode(data)
}

func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

// making a validation error dude

func ValidationError(errs validator.ValidationErrors) Response {
	//  check each field which one  has the errror dude
	var errMssg []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMssg = append(errMssg, fmt.Sprintf("field %s is required field", err.Field()))
		default:
			errMssg = append(errMssg, fmt.Sprintf("field  %s is invalid field", err.Field()))
		}

	}
	return Response{
		Status: StatusError,
		Error:  strings.Join(errMssg, ","),
	}
}
