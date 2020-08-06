package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/zoomi-raja/goweb/api/security"
	"github.com/zoomi-raja/goweb/api/utils"
	"github.com/zoomi-raja/goweb/database"
)

type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
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

func (u User) GetUser() (User, error) {
	user := User{}

	if u.Email == "" {
		return user, errors.New("email missing for query")
	}
	db, _ := database.Connect()
	defer db.Close()
	row := db.QueryRow(`SELECT id, email, user_name, ifnull(avatar, ''), created_at,password FROM users WHERE email = ?`, u.Email)
	if err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Avatar, &user.CreatedAt, &user.Password); err != nil {
		return user, err
	}
	dateTime, _ := utils.FormateDate(user.CreatedAt)
	user.CreatedAt = dateTime
	return user, nil
}

//CreateUser create user and will also hash the password
func (u User) CreateUser() (int64, error) {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return 0, err
	}

	db, _ := database.Connect()
	defer db.Close()
	result, err := db.Exec(fmt.Sprintf("insert into %s (email, user_name, password) values (?,?,?)", userTable), u.Email, u.Username, hashedPassword)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	id, _ := result.LastInsertId()
	u.ID = id
	return id, nil
}

type UserToken struct {
	Token   string
	ExpTime int64
}

func (u User) CreatenToken() (UserToken, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = u.ID
	atClaims["name"] = u.Username
	atClaims["email"] = u.Email
	atClaims["avatar"] = u.Avatar
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte("mywebblog"))
	if err != nil {
		return UserToken{}, err
	}
	return UserToken{token, atClaims["exp"].(int64)}, nil //os.Getenv("ACCESS_SECRET")
}
