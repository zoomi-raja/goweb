package responses

import (
	"encoding/json"
	"fmt"
	"gotut/api/utils"
	"net/http"
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
		// fmt.Println(reflect.TypeOf(err))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(struct {
			Status bool   `json:"status"`
			Error  string `json:"error"`
		}{
			false, err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
func ERRORARR(w http.ResponseWriter, statusCode int, errs utils.ErrorsArr) {
	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(struct {
			Status bool                   `json:"status"`
			Error  map[string]interface{} `json:"error"`
		}{
			false, errs.GetErros(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
