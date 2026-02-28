package sqlite

import (
	"database/sql"

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
