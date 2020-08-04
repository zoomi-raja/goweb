package controllers

import (
	"encoding/json"
	"gotut/api/models"
	"gotut/api/requests"
	"gotut/api/responses"
	"net/http"
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
		responses.ERRORARR(w, http.StatusInternalServerError, err)
		return
	}
	userModel := models.User{Email: user.Email, Username: user.Username, Password: user.Password}
	if result, err := userModel.CreateUser(); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	} else {
		responses.JSON(w, http.StatusOK, result)
	}
}
