package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/zoomi-raja/goweb/api/responses"

	"github.com/zoomi-raja/goweb/api/utils"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Addmiddlewares(h http.HandlerFunc, m []Middleware) http.HandlerFunc {
	if len(m) < 1 {
		return h
	}

	wrapped := h

	// loop in reverse to preserve middleware order
	for i := len(m) - 1; i >= 0; i-- {
		fmt.Println(utils.GetFunctionName(m[i]))
		wrapped = m[i](wrapped)
	}

	return wrapped
}
func SetMiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s%s %s", r.Method, r.Host, r.RequestURI, r.Proto)
		next(w, r)
	}
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("validate cookie")
		if c, err := r.Cookie("token"); err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("protected resource"))
		} else {
			atClaims := jwt.MapClaims{}
			tkn, err := jwt.ParseWithClaims(c.Value, atClaims, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte("mywebblog"), nil
			})
			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}

				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if !tkn.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			fmt.Println(atClaims)
			next(w, r)
		}
	}
}
