package models

import (
	"fmt"

	"github.com/zoomi-raja/goweb/api/utils"
	"github.com/zoomi-raja/goweb/database"
)

type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"createdAt"`
}

var userTable string = "users"

// GetAllUsers will return all users with out showing password property
func (u User) GetAllUsers() ([]User, error) {
	users := make([]User, 0)
	db, _ := database.Connect()
	defer db.Close()
	result, err := db.Query(fmt.Sprintf("SELECT id, email, user_name, ifnull(avatar, ''), created_at FROM %s", userTable))
	if err != nil {
		return nil, err
	}
	defer result.Close()
	for result.Next() {
		user := User{}
		result.Scan(&user.ID, &user.Email, &user.Username, &user.Avatar, &user.CreatedAt)
		dateTime, _ := utils.FormateDate(user.CreatedAt)

		user.CreatedAt = dateTime
		users = append(users, user)
	}
	return users, nil
}
