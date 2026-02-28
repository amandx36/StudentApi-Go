package sqlite

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/amandx36/studentCrudApiGo/internal/types"


	"github.com/amandx36/studentCrudApiGo/internal/config"
	_ "github.com/mattn/go-sqlite3" // register sqlite driver
)

// Sqlite wraps DB connection
type Sqlite struct {
	// filed name and data type
	Db *sql.DB // connection manager
}

// New initializes DB and returns repository
func New(cfg *config.Config) (*Sqlite, error) {

	// 1. Open DB (creates connection pool)
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	// 2. Creating  table if not exists
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS students (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			email TEXT,
			age INTEGER
		)
	`)
	// error handling
	if err != nil {
		return nil, err
	}

	// create empty Sqlite struct
	repo := Sqlite{}
	// Access its filed Db and store the connection into this dude

	repo.Db = db

	// 3. Return repository with DB
	return &repo, nil
}

// implementing the interface dude  and attaching the interface to sqlite dude


func (s *Sqlite) CreateStudent(name string, email string, age int64) (int64, error) {
	statement, err := s.Db.Prepare("INSERT INTO students(name,email,age)values(?,?,?)")

	if err != nil {
		return 0, nil
	}
	defer statement.Close() // free the resource

	result, err := statement.Exec(name, email, age)
	if err != nil {
		return 0, nil
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}


// now again attach to the struct dude 

func (s *Sqlite)GetStudentById(id int64) (types.Student , error){ 

	statement , err := s.Db.Prepare(

		"SELECT id, name, email, age FROM students WHERE id = ? LIMIT 1",

	)
	if err != nil{
		return types.Student{},err
	}
	defer statement.Close()

	// all clear than return the data dude 
	var student types.Student

	// now put into to the data base dude
	err = statement.QueryRow(id).Scan(&student.Id,&student.Name,&student.Email,&student.Age)

	// if we got  the error than we do this 
	if err !=nil{
		// if no user found than do this dude 
		if err ==sql.ErrNoRows{
			return types.Student{} , fmt.Errorf("No Student with this id  %d found ",id)
		}

		// if other error happends than do this  dude 
		return types.Student{} , fmt.Errorf("Querry error %w",err)
	}
	return student , nil

}

// implementing to all the list dude

func (s *Sqlite)GetStudents()([] types.Student, error){
	statement , err := s.Db.Prepare(

		"SELECT id, name, email, age FROM students ",

	)
	if err !=nil{
		slog.Info("Error in getting the users",err)
	}
	
	defer statement.Close()


	rows , err := statement.Query()
	if err !=nil{
		return nil , err 

	}
	defer rows.Close()

	var students []types.Student 

	// for loop in all the loops and than put inside dude 

	for rows.Next(){
		var student  types.Student

		error := rows.Scan(&student.Id , & student.Name , &student.Email , &student.Age )
		 if error !=nil{
			return nil , error 
		 }
		 students = append(students , student)

	}
	return students , nil 
}


