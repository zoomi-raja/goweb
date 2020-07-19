package controllers

import (
	"encoding/json"
	"gotut/api/models"
	"gotut/database"
	"net/http"
)

// type PostHandler struct {
// 	db *sql.DB
// }

// func NewHandler(db *sql.DB) *PostHandler {
// 	return &PostHandler{
// 		db: db,
// 	}
// }

// func (p PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var posts []post
// 	result, err := p.db.Query("SELECT id, title from posts")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer result.Close()
// 	for result.Next() {
// 		var post post
// 		result.Scan(&post.ID, &post.Title)
// 		posts = append(posts, post)
// 	}
// 	json.NewEncoder(w).Encode(posts)
// }

func GetPosts(w http.ResponseWriter, r *http.Request) {
	db, _ := database.Connect()
	defer db.Close()
	w.Header().Set("Content-Type", "application/json")
	var posts []models.Post
	result, err := db.Query("SELECT id, title from posts")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var post models.Post
		result.Scan(&post.ID, &post.Title)
		posts = append(posts, post)
	}
	json.NewEncoder(w).Encode(posts)
}
