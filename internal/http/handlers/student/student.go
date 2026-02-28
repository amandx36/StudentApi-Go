package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/amandx36/studentCrudApiGo/internal/types"
	"github.com/amandx36/studentCrudApiGo/internal/utils/response"
)

// 1    New() returns an HTTP handler function for creating a student
func New() http.HandlerFunc {

	// 2️  Handler receives:
	//    - responseSender :) for sending request 
	//    - requesting     get from frontEnd 
	return func(responseSender http.ResponseWriter, requesting *http.Request) {

		// 3️  Create empty Student struct
		var student types.Student

		// 4️  Decode JSON request body into Student struct
		err := json.NewDecoder(requesting.Body).Decode(&student)

		// 5️  Handle decoding errors
		if err != nil {

			// 5   If body is empty → return 400 Bad Request
			if errors.Is(err, io.EOF) {
				response.WriteJson(responseSender, http.StatusBadRequest, map[string]string{
					"error": "empty request body",
				})
				return
			}

			// 5   If JSON is invalid → return 400 Bad Request
			response.WriteJson(responseSender, http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			return
		}

		// 6  Log 
		slog.Info("Creating a student")

		// 7️ for storing the data into database 
		// service.CreateStudent(student)

		// 8️  Send success response
		response.WriteJson(responseSender, http.StatusCreated, map[string]string{
			"message": "Successfully created",
		})
	}
}