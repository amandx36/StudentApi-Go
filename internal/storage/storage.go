package storage

// import "github.com/amandx36/studentCrudApiGo/internal/types"

// making a interface for making all things
type Storage interface {
	CreateStudent(name string, email string, age int64) (int64 , error)

	// GetStudentById(id int64) (types.Student, error)

	// GetStudents() ([]types.Student, error)
}
