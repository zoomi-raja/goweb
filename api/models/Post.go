package models

import (
	"fmt"
	"gotut/database"
)

type Post struct {
	ID        uint32 `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"createdAt"`
}

func (p Post) GetAll() ([]Post, error) {
	var posts []Post
	db, _ := database.Connect()
	defer db.Close()
	result, err := db.Query("SELECT id, title, body, createdAt FROM posts")
	if err != nil {
		return nil, err
	}
	defer result.Close()
	for result.Next() {
		var post Post
		result.Scan(&post.ID, &post.Title, &post.Body, &post.CreatedAt)
		fmt.Println(post)
		posts = append(posts, post)
	}
	return posts, nil
}

//todo query should be here
