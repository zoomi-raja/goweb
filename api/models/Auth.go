package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/zoomi-raja/goweb/api/security"
	"github.com/zoomi-raja/goweb/database"
)

type Auth struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

var authTable string = "users"

func (u Auth) GetUser() (Auth, error) {
	user := Auth{}

	if u.Email == "" {
		return user, errors.New("email missing for query")
	}
	db, _ := database.Connect()
	defer db.Close()
	row := db.QueryRow(`SELECT id, email, user_name, ifnull(avatar, ''), password FROM users WHERE email = ?`, u.Email)
	if err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Avatar, &user.Password); err != nil {
		return user, err
	}
	return user, nil
}

//CreateUser create user and will also hash the password
func (a Auth) CreateUser() (int64, error) {
	hashedPassword, err := security.Hash(a.Password)
	if err != nil {
		return 0, err
	}

	db, _ := database.Connect()
	defer db.Close()
	result, err := db.Exec(fmt.Sprintf("insert into %s (email, user_name, password) values (?,?,?)", authTable), a.Email, a.Username, hashedPassword)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	id, _ := result.LastInsertId()
	a.ID = id
	return id, nil
}

type UserToken struct {
	Token   string
	ExpTime int64
}

func (a Auth) CreatenToken() (UserToken, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = a.ID
	atClaims["name"] = a.Username
	atClaims["email"] = a.Email
	atClaims["avatar"] = a.Avatar
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte("mywebblog"))
	if err != nil {
		return UserToken{}, err
	}
	return UserToken{token, atClaims["exp"].(int64)}, nil //os.Getenv("ACCESS_SECRET")
}
