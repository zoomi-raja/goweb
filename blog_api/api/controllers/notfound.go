package controllers

import (
	"errors"
	"net/http"

	"github.com/zoomi-raja/goweb/api/responses"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	responses.ERROR(w, http.StatusNotFound, errors.New("No such path found"))
}
