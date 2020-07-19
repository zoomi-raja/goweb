package api

import (
	"fmt"
	"gotut/api/router"
	"gotut/config"
	"log"
	"net/http"
)

func Run() {
	fmt.Printf("listening [::]:%d\n", config.PORT)
	listen(config.PORT)
}
func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
