package database

import (
	"database/sql"

	"github.com/zoomi-raja/goweb/config"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	return sql.Open(config.DBDRIVER, config.DBURL)
}
