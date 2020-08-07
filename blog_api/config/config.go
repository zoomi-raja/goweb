package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	PORT     = 0
	DBDRIVER = ""
	DBURL    = ""
)

func loadEnv() {
	var err error
	if value, _ := strconv.ParseBool(os.Getenv("DEBUGGER")); !value {
		err = godotenv.Load()
		if err != nil {
			fmt.Println(err)
			log.Fatal("Error loading .env file")
		}
	}
}
func Load() {
	var err error
	loadEnv()
	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Println(err)
	}
	DBDRIVER = os.Getenv("DB_DRIVER")
	DBURL = fmt.Sprintf("%s:%s@tcp(go_db:3306)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
}
