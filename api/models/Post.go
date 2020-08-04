package models

import (
	"fmt"
	"gotut/api/utils"
	"gotut/database"
)

type Post struct {
	ID        uint32 `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"createdAt"`
}

var table string = "posts"

func (p Post) GetAll() ([]Post, error) {
	posts := make([]Post, 0)
	db, _ := database.Connect()
	defer db.Close()
	result, err := db.Query(fmt.Sprintf("SELECT id, title, body, created_at FROM %s", table))
	if err != nil {
		return nil, err
	}
	defer result.Close()
	for result.Next() {
		var post Post
		result.Scan(&post.ID, &post.Title, &post.Body, &post.CreatedAt)
		dateTime, _ := utils.FormateDate(post.CreatedAt)
		post.CreatedAt = dateTime
		posts = append(posts, post)
	}
	return posts, nil
}

//todo query should be here
