package models

import (
	"github.com/look4suman/events-api/db"
	"github.com/look4suman/events-api/routes/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u User) Save() (*User, error) {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}

	result, err := db.DB.Exec(query, u.Email, hashedPassword)
	if err != nil {
		return &u, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return &u, err
	}
	u.ID = id
	return &u, nil
}

func GetAllUsers() ([]User, error) {
	query := `SELECT id, email, password FROM users`
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
		users = append(users, user)
	}

	return users, nil
}

func GetUserByEmail(user User) (*User, error) {
	query := `SELECT id, email, password FROM users where email = ?`
	row := db.DB.QueryRow(query, user.Email)

	var dbUser User
	err := row.Scan(&dbUser.ID, &dbUser.Email, &dbUser.Password)
	if err != nil {
		return nil, err
	}

	isValidPassword := utils.CheckPasswordHash(user.Password, dbUser.Password)
	if isValidPassword {
		return &user, nil
	}
	return nil, nil
}
