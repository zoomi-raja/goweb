package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zoomi-raja/goweb/api/router"
	"github.com/zoomi-raja/goweb/config"
)

func Run() {
	fmt.Printf("listening [::]:%d\n", config.PORT)
	listen(config.PORT)
}
func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
