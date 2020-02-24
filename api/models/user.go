package models

import (
	"database/sql"
	"log"
)

type User struct {
	Id          int    `json:"id"`
	Nickname    string `json:"nickname"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
	// Avatar      string `json:"avatar"`
}

func (u *User) FindAllUsers(db *sql.DB) ([]User, error) {

	rows, err := db.Query("SELECT id, nickname, phone_number FROM user")

	if err != nil {
		log.Println("Error while retrieving all users")
		return nil, err
	}

	// var users []User
	users := make([]User, 0)

	for rows.Next() {
		user := User{}

		err = rows.Scan(&user.Id, &user.Nickname, &user.PhoneNumber)

		if err != nil {
			log.Println("Error by parsing - findAllUsers")
			return nil, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *User) CreateUser(db *sql.DB) (uint32, error) {

	var lastUserId uint32

	err := db.QueryRow("CALL create_user(?, ?, ?)", u.Nickname, u.Password, u.PhoneNumber).Scan(&lastUserId)

	if err != nil {
		return 0, err
	}

	return lastUserId, nil
}
