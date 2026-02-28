package response

import (
	"encoding/json"
	"net/http"
)

// for encoding from struct to json

//   any / interface{}   both are same dude


func WriteJson( writer http.ResponseWriter, statusCode int , data any ) error {

	writer.Header().Set("Content-Type","application/json")
	writer.WriteHeader(statusCode)

	return json.NewEncoder(writer).Encode(data)
}