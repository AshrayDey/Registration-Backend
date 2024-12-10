package models

import (
	"errors"
	"registrationApp/db"
	"registrationApp/utils"
)

type User struct {
	ID       int64
	Name     string `binding:"required"`
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `
	INSERT INTO users (name, email, password)
	VALUES (?, ?, ?)
	`

	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	_, err = statement.Exec(u.Name, u.Email, hashPassword)
	if err != nil {
		return err
	}
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id,password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrivedPassword string
	var userId int64

	err := row.Scan(&userId, &retrivedPassword)
	if err != nil {
		return err
	}

	passwordValid := utils.CheckPasswordHash(u.Password, retrivedPassword)
	if !passwordValid {
		return errors.New("Incorrect Credentials")
	}

	return nil
}

func GetAllUser() ([]User, error) {
	query := "SELECT * FROM users"
	rows, err := db.DB.Query(query)

	var users []User
	var user User
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, err
}
