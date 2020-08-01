package responses

import (
	"encoding/json"
	"fmt"
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
		JSON(w, statusCode, struct {
			Status bool   `json:"status"`
			Error  string `json:"error"`
		}{
			false, err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
