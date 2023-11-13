package models

import (
	"backend/utils"
	"log"
	"time"
)

type User struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) CreateAnUser() (*User, error) {
	statement := "INSERT INTO posts (id, name, email, password, image_url) VALUES (?,?,?,?,?)"
	u.Id = utils.GetRandomString(32)

	_, err := db.Exec(statement, u.Id, u.Name, u.Email, u.Password, u.ImageUrl)

	if err != nil {
		return nil, err
	}

	log.Printf("crated a new user with id: %s", u.Id)
	return u, nil
}

func GetAllUsers() ([]User, error) {
	var users []User
	statement := "SELECT id, name, email, image_url, created_at FROM users"

	rows, err := db.Query(statement)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user User

		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.ImageUrl, &user.CreatedAt)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func GetUserById(id string) (User, error) {
	var user User
	statement := "SELECT * FROM users WHERE id = ?"

	row, err := db.Query(statement, id)
	defer row.Close()
	if err != nil {
		return User{}, err
	}

	for row.Next() {
		err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.ImageUrl, &user.CreatedAt)

		if err != nil {
			return User{}, err
		}

	}

	return user, nil
}

func GetUserByEmail(email string) (User, error) {
	var user User
	statement := "SELECT * FROM users WHERE id = ?"

	row, err := db.Query(statement, email)
	defer row.Close()
	if err != nil {
		return User{}, err
	}

	for row.Next() {
		err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.ImageUrl, &user.CreatedAt)

		if err != nil {
			return User{}, err
		}

	}

	return user, nil
}
