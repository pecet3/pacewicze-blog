package models

import (
	"backend/utils"
	"log"
	"time"
)

type Post struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	ImageUrl  string    `json:"image_url"`
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
		err := rows.Scan(&post.Id, &post.UserId, &post.Title, &post.Content, &post.ImageUrl, &post.CreatedAt)

		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (p *Post) CreateAPost() (*Post, error) {
	statement := "INSERT INTO posts (id, user_id, title, content, image_url) VALUES (?,?,?,?,?)"

	p.Id = utils.GetRandomString(32)

	_, err := db.Exec(statement, p.Id, p.UserId, p.Title, p.Content, p.ImageUrl)
	if err != nil {
		return nil, err
	}

	log.Printf("user with id: %s has created a new post with id: %s", p.UserId, p.Id)
	return p, nil
}

func GetPostById(Id string) (Post, error) {
	var post Post
	statement := "SELECT * FROM posts WHERE ID = ?"

	row, err := db.Query(statement, Id)
	if err != nil {
		return Post{}, err
	}
	for row.Next() {
		err := row.Scan(&post.Id, &post.UserId, &post.Title, &post.Content, &post.ImageUrl, &post.CreatedAt)

		if err != nil {
			return Post{}, err
		}
	}

	return post, nil
}
