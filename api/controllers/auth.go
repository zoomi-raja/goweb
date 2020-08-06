package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

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
	userModel := models.User{Email: req.Email, Password: req.Password}
	if user, err := userModel.GetUser(); err != nil {
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
			user.Password = ""
			//set auth info in cookie
			cookies.SetAuthCookie(w, token)
			responses.JSON(w, http.StatusOK, user)
		}
	}
}
