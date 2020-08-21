package responses

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zoomi-raja/goweb/api/requests"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(struct {
		Status bool        `json:"status"`
		Data   interface{} `json:"data"`
	}{true, data})
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		errInfo := make(map[string]interface{})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		switch errType := err.(type) {
		case *requests.ErrorArr:
			errs := err.(*requests.ErrorArr)
			errInfo["ERR"] = errs.GetErros()
			errInfo["TYPE"] = errType
		default:
			errInfo["ERR"] = err.Error()
		}
		json.NewEncoder(w).Encode(struct {
			Status bool        `json:"status"`
			Error  interface{} `json:"error"`
		}{
			false, errInfo["ERR"],
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
