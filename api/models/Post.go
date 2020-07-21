package models

import "gotut/database"

type Post struct {
	ID    uint32 `json:"id"`
	Title string `json:"title"`
}

func (p Post) GetAll() ([]Post, error) {
	var posts []Post
	db, _ := database.Connect()
	defer db.Close()
	result, err := db.Query("SELECT id, title from posts")
	if err != nil {
		return nil, err
	}
	defer result.Close()
	for result.Next() {
		var post Post
		result.Scan(&post.ID, &post.Title)
		posts = append(posts, post)
	}
	return posts, nil
}

//todo query should be here
