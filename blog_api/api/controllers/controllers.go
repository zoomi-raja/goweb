package controllers

import (
	"database/sql"
	"net/http"
)

type Controllers interface {
	action(*sql.DB, string) func(http.ResponseWriter, *http.Request)
}
