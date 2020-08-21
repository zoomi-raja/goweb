package models

import (
	"encoding/base64"
	"fmt"

	"github.com/zoomi-raja/goweb/api/utils"
	"github.com/zoomi-raja/goweb/database"
)

type Post struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"userId"`
	Claps     int64  `json:"claps"`
	Intro     string `json:"intro"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

var postTable string = "posts"

func (p Post) GetAll() ([]Post, error) {
	posts := make([]Post, 0)
	db, _ := database.Connect()
	defer db.Close()
	result, err := db.Query(fmt.Sprintf("SELECT id, claps, intro, title, created_at, updated_at FROM %s", postTable))
	if err != nil {
		return nil, err
	}
	defer result.Close()
	for result.Next() {
		var post Post
		result.Scan(&post.ID, &post.Claps, &post.Intro, &post.Title, &post.CreatedAt, &post.UpdatedAt)
		dateTime, _ := utils.FormateDate(post.CreatedAt)
		post.CreatedAt = dateTime
		dateTime2, _ := utils.FormateDate(post.UpdatedAt)
		post.UpdatedAt = dateTime2

		posts = append(posts, post)
	}
	return posts, nil
}

func (p Post) GetById(id int) (Post, error) {
	var post Post
	db, _ := database.Connect()
	defer db.Close()
	sqlQuery := fmt.Sprintf(`SELECT id, claps, intro, title, body, created_at, updated_at FROM %s WHERE id=?;`, postTable)

	row := db.QueryRow(sqlQuery, id)
	if err := row.Scan(&post.ID, &post.Claps, &post.Intro, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt); err != nil {
		return post, err
	} else {
		dateTime, _ := utils.FormateDate(post.CreatedAt)
		post.CreatedAt = dateTime
		dateTime2, _ := utils.FormateDate(post.UpdatedAt)
		post.UpdatedAt = dateTime2
		if data, err := base64.StdEncoding.DecodeString(post.Body); err != nil {
			return post, err
		} else {
			post.Body = string(data)
		}
		return post, nil
	}
}

/*CreatePost to create post in db*/
func (p Post) CreatePost() (int64, error) {
	db, _ := database.Connect()
	defer db.Close()
	result, err := db.Exec(fmt.Sprintf("insert into %s (title, intro, body, user_id) values (?,?,?,?)", postTable), p.Title, p.Intro, p.Body, p.UserID)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	p.ID = id
	return id, nil
}

//todo query should be here
