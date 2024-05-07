package models

import (
	"database/sql"
)

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FullName  string `json:"fullName"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func CreateTables(db *sql.DB) error {
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT UNIQUE NOT NULL,
        email TEXT UNIQUE NOT NULL,
        firstname TEXT NOT NULL,
        lastname TEXT NOT NULL,
        password TEXT NOT NULL
    );`
	_, err := db.Exec(createTableSQL)
	return err
}

func GetUserByUsername(db *sql.DB, username string) (*User, error) {
	var user User
	query := "SELECT id, username, email, firstname, lastname FROM users WHERE username = ?"
	err := db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Email, &user.FirstName, &user.LastName)
	if err != nil {
		return nil, err
	}
	user.FullName = user.FirstName + " " + user.LastName
	return &user, nil
}
