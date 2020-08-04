package models

import (
	"fmt"

	"github.com/zoomi-raja/goweb/api/security"
	"github.com/zoomi-raja/goweb/api/utils"
	"github.com/zoomi-raja/goweb/database"
)

type User struct {
	ID        uint32 `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"createdAt"`
}

var userTable string = "users"

func (u User) GetAllUsers() ([]User, error) {
	users := make([]User, 0)
	db, _ := database.Connect()
	defer db.Close()
	result, err := db.Query(fmt.Sprintf("SELECT id, email, user_name, avatar, created_at FROM %s", userTable))
	if err != nil {
		return nil, err
	}
	defer result.Close()
	for result.Next() {
		var user User
		result.Scan(&user.ID, &user.Email, &user.Username, &user.Avatar, &user.CreatedAt)
		dateTime, _ := utils.FormateDate(user.CreatedAt)
		user.CreatedAt = dateTime
		users = append(users, user)
	}
	return users, nil
}

func (u User) CreateUser() (int64, error) {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return 0, err
	}
	db, _ := database.Connect()
	defer db.Close()
	result, err := db.Exec(fmt.Sprintf("insert into %s (email, user_name, password) values (?,?,?)", userTable), u.Email, u.Username, hashedPassword)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	return id, nil
}
