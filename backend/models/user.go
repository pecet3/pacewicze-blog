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
	Salt      string    `json:"salt"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) CreateAnUser() (*User, error) {
	statement := "INSERT INTO users (id, name, email, password, salt, image_url) VALUES (?,?,?,?,?,?)"
	u.Id = utils.GetRandomString(32)
	salt, err := utils.GenerateSalt(16)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	hashedPassword, err := utils.HashPassword(u.Password, salt)
	if err != nil {
		return nil, err
	}

	u.Password = hashedPassword
	u.Salt = salt

	_, err = db.Exec(statement, u.Id, u.Name, u.Email, u.Password, u.Salt, u.ImageUrl)
	if err != nil {
		log.Println(err)
	}
	if err != nil {
		log.Println(err, 2)
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
		err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.ImageUrl, &user.Salt, &user.CreatedAt)

		if err != nil {
			return User{}, err
		}
	}

	return user, nil
}

func GetUserByEmail(email string) (User, error) {
	var user User
	statement := "SELECT * FROM users WHERE email = ?"

	row, err := db.Query(statement, email)
	defer row.Close()
	if err != nil {
		log.Println(err)
		return User{}, err
	}

	for row.Next() {
		err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.ImageUrl, &user.Salt, &user.CreatedAt)

		if err != nil {
			log.Println(err)
			return User{}, err

		}
	}

	return user, nil
}
