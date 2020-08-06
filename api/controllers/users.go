package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/zoomi-raja/goweb/api/cookies"
	"github.com/zoomi-raja/goweb/api/models"
	"github.com/zoomi-raja/goweb/api/requests"
	"github.com/zoomi-raja/goweb/api/responses"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	user := models.User{}
	if users, err := user.GetAllUsers(); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	} else {
		responses.JSON(w, http.StatusOK, users)
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
	userModel := models.User{Email: user.Email, Username: user.Username, Password: user.Password}
	if userId, err := userModel.CreateUser(); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err) //*mysql.MySQLError error type
	} else {
		// os.Setenv("ACCESS_SECRET", "mywebblog")
		token, err := userModel.CreatenToken()
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
