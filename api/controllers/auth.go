package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/zoomi-raja/goweb/api/cookies"
	"github.com/zoomi-raja/goweb/api/security"

	"github.com/zoomi-raja/goweb/api/models"
	"github.com/zoomi-raja/goweb/api/responses"

	"github.com/zoomi-raja/goweb/api/requests"
)

func Login(w http.ResponseWriter, r *http.Request) {
	req := requests.Credentials{}
	json.NewDecoder(r.Body).Decode(&req)
	if err := req.ValidateLoginRequest(); err.HasError() {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	authModel := models.Auth{Email: req.Email, Password: req.Password}
	if user, err := authModel.GetUser(); err != nil {
		switch err {
		case sql.ErrNoRows:
			responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("Username or password is invalid"))
		default:
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
		}
	} else {
		//validate password
		if err := security.VerifyPassword(user.Password, req.Password); err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("Username or password is invalid"))
		} else {
			token, err := user.CreatenToken()
			if err != nil {
				responses.ERROR(w, http.StatusInternalServerError, err)
				return
			}
			userData := map[string]string{
				"id":       strconv.FormatInt(user.ID, 10),
				"username": user.Username,
				"email":    user.Email,
				"avatar":   user.Avatar,
			}
			//set auth info in cookie
			cookies.SetAuthCookie(w, token)
			responses.JSON(w, http.StatusOK, userData)
		}
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := requests.User{}
	json.NewDecoder(r.Body).Decode(&user)
	err := user.ValidateUserCreate()
	if err.HasError() {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	authModel := models.Auth{Email: user.Email, Username: user.Username, Password: user.Password}
	if userId, err := authModel.CreateUser(); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err) //*mysql.MySQLError error type
	} else {
		token, err := authModel.CreatenToken()
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		//set auth info in cookie
		cookies.SetAuthCookie(w, token)
		responses.JSON(w, http.StatusOK, struct {
			UserId int64  `json:"userID"`
			Token  string `json:"token"`
		}{userId, token.Token})
	}
}
