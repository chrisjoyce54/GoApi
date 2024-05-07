package models

import (
	"errors"

	"github.com/chrisjoyce54/GoApi/db"
	"github.com/chrisjoyce54/GoApi/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() (*User, error) {
	//will eventually be to db
	//events = append(events, *e)
	query := `
	INSERT INTO Users(email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	u.ID = id
	u.Password = ""
	return u, err
}

func (u *User) ValidateCredentials() error {
	query := "Select ID, password from users where email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("Invalid credentials")
	}

	passwordValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordValid {
		return errors.New("Invalid credentials")
	}

	return nil
}

func GetUsers() ([]User, error) {
	query := "SELECT * FROM users"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Email, &user.Password)

		if err != nil {
			return nil, err
		}
		user.Password = ""
		users = append(users, user)
	}

	return users, nil
}
