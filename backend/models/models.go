package models

import (
	"backend/utils"
	"fmt"
	"time"
)

type Post struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func GetAllPosts() ([]Post, error) {
	var posts []Post

	statement := "SELECT * FROM posts"

	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		err := rows.Scan(&post.Id, &post.UserId, &post.Title, &post.Content, &post.CreatedAt)

		if err != nil {
			return nil, err
		}
		fmt.Println(post)
		posts = append(posts, post)
	}

	return posts, nil
}

func (p *Post) CreateAPost() (*Post, error) {
	statement := "INSERT INTO posts (id, user_id, title, content) VALUES (?,?,?,?)"

	p.Id = utils.GetRandomString(32)

	_, err := db.Exec(statement, p.Id, p.UserId, p.Title, p.Content)
	if err != nil {
		return nil, err
	}

	return p, nil
}
