package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/zoomi-raja/goweb/api/router"
	"github.com/zoomi-raja/goweb/config"
)

func Run() {
	fmt.Printf("listening [::]:%d\n", config.PORT)
	listen(config.PORT)
}
func listen(port int) {
	r := router.New()
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:3001/"})
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		handlers.CORS(header, methods, origins)(http.TimeoutHandler(r, time.Second*1, "timeout"))))
}
