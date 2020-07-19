package controllers

import (
	"encoding/json"
	"gotut/database"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	defer db.Close()
	w.Header().Set("Content-Type", "application/json")
	var posts []post
	result, err := db.Query("SELECT id, title from posts")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var post post
		result.Scan(&post.ID, &post.Title)
		posts = append(posts, post)
	}
	json.NewEncoder(w).Encode(posts)
}
