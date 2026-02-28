package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/amandx36/studentCrudApiGo/internal/storage"
	"github.com/amandx36/studentCrudApiGo/internal/types"
	"github.com/amandx36/studentCrudApiGo/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

// 1    New() returns an HTTP handler function for creating a student
// pass the interface of all type so  that we can implemnt it dude
func New(storage storage.Storage) http.HandlerFunc {

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

			// 5   If body is eStudentmpty → return 400 Bad Request
			if errors.Is(err, io.EOF) {
				response.GeneralError(err)
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

		// 7 Request validation
		// using inbuilt package
		//   go get github.com/go-playground/validator/v10

		// we have to add the field  validation according to  in response struct

		if err := validator.New().Struct(student); err != nil {

			// now type casting because we need to transform into other one dude
			validateErrors := err.(validator.ValidationErrors)

			response.WriteJson(responseSender, http.StatusBadRequest, response.ValidationError(validateErrors))
			return
		}

		// 7️ for storing the data into database
		//
		lastId, err := storage.CreateStudent(
			student.Name,
			student.Email,
			int64(student.Age),
		)
		if err != nil {
			response.WriteJson(responseSender, http.StatusInternalServerError, err)
			return
		}

		// 8️  Send success response
		response.WriteJson(responseSender, http.StatusCreated, map[string]string{
			"message": "Successfully created",
			"Id":      string(lastId),
		})
		slog.Info("User created sucessfully dude ", slog.String("With User id  ", fmt.Sprint(lastId)))
	}
}

func GetStudentById(storage storage.Storage) http.HandlerFunc {
	return func(responseSender http.ResponseWriter, requesting *http.Request) {
		id := requesting.PathValue("id")

		slog.Info("Getting a student", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.WriteJson(responseSender, http.StatusBadRequest, response.GeneralError(err))
		}
		student, err := storage.GetStudentById(intId)

		if err != nil {
			slog.Error("Error  in the getting users", slog.String("id", fmt.Sprint(intId)))

			response.WriteJson(responseSender, http.StatusInternalServerError, response.GeneralError(err))
			return

		}
		response.WriteJson(responseSender, http.StatusOK, student)
	}

}
// implementing other api end points dude 

func GetList(storage storage.Storage) http.HandlerFunc {
	return func(responseSender http.ResponseWriter, requesting *http.Request) {
		
		slog.Info("Getting all the students ")
		students , err := storage.GetStudents()
		if err !=nil{
			response.WriteJson(responseSender,http.StatusInternalServerError,err )
		}

		response.WriteJson(responseSender,http.StatusOK,students)
	}

}
