package middleware

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/zoomi-raja/goweb/api/models"

	"github.com/zoomi-raja/goweb/api/responses"
)

//Middleware type with handerfunc to maintain the array of middlewares
type Middleware func(http.HandlerFunc) http.HandlerFunc

//Addmiddlewares helper function to arrange middle ware functions
func Addmiddlewares(h http.HandlerFunc, m []Middleware) http.HandlerFunc {
	if len(m) < 1 {
		return h
	}

	wrapped := h

	// loop in reverse to preserve middleware order
	for i := len(m) - 1; i >= 0; i-- {
		// fmt.Println(utils.GetFunctionName(m[i]))
		wrapped = m[i](wrapped)
	}

	return wrapped
}

//SetMiddlewareLogger our dummy logger as of now
func SetMiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s%s %s", r.Method, r.Host, r.RequestURI, r.Proto)
		next(w, r)
	}
}

//AuthMiddleware to protext route only for logged in user
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("validate cookie")
		if c, err := r.Cookie("token"); err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("protected resource"))
		} else {
			auth := models.Auth{}
			if claim, err := auth.VerifyToken(c.Value); err != nil {
				responses.ERROR(w, http.StatusUnauthorized, err)
			} else {
				ctx := context.WithValue(r.Context(), models.AuthCTX("user"), claim) //context science
				next(w, r.WithContext(ctx))
			}
		}
	}
}
