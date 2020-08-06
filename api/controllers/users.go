package controllers

import (
	"net/http"

	"github.com/zoomi-raja/goweb/api/models"
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
