package controllers

import (
	"errors"
	"gotut/api/responses"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	responses.ERROR(w, http.StatusNotFound, errors.New("No such path found"))
}
