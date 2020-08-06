package cookies

import (
	"net/http"
	"time"

	"github.com/zoomi-raja/goweb/api/models"
)

func SetAuthCookie(w http.ResponseWriter, t models.UserToken) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    t.Token,
		HttpOnly: true,
		Expires:  time.Unix(t.ExpTime, 0),
	})
}
