package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDb() {
	var err error
	DB, err = sql.Open("sqlite", "api.db")

	if err != nil {
		panic("Could not connect to database:" + err.Error() + ".")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		Email Text NOT NULL Unique,
		Password Text Not Null
	)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create users table: " + err.Error() + ".")
	}

	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			Name TEXT NOT NULL,
			Description TEXT NOT NULL,
			Location TEXT NOT NULL,
			DateTime DATETIME NOT NULL,
			User_id INTEGER,
			Foreign key(User_id) References Users(ID)
		)
	`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table:" + err.Error() + ".")
	}

	createRegistrationsTable := `
		Create Table if Not Exists registrations (
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			Event_ID Integer,
			User_Id Integet,
			Foreign Key(Event_Id) References Events(ID),
			Foreign Key(User_Id) References Users(ID)
		)
		`

	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		panic("Could not create registration table:" + err.Error() + ".")
	}
}
