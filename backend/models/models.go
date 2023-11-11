package models

import (
	"fmt"
	"time"
)

type Post struct {
	Id        uint64    `json:"id"`
	UserId    uint64    `json:"user_id"`
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

		posts = append(posts, post)
	}

	return posts, nil
}

func (p *Post) CreateAPost() (*Post, error) {
	statement := "INSERT INTO posts (user_id, title, content) VALUES (?,?,?)"
	fmt.Println(p)
	_, err := db.Query(statement, 0, "title", "lorem ipsum")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return p, nil
}
