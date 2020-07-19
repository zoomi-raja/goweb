package database

import (
	"database/sql"
	"gotut/config"
)

func Connect() (*sql.DB, error) {
	return sql.Open(config.DBDRIVER, config.DBURL)
}
