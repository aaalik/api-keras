package helper

import (
	"encoding/json"
	"net/http"
)

const HTTPStatusSuccess string = "success"
const HTTPStatusError string = "error"

type ErrorResponse struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Code   int    `json:"code"`
}

//ToJSON - convert struct to JSON
func JSONResponse(w http.ResponseWriter, status string, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)

	if status == HTTPStatusSuccess {
		Log.Info(body)
	} else {
		Log.Error(body)
	}
}

//ToJSON -> convert interface to JSON
func ToJSON(d interface{}) string {
	j, _ := json.Marshal(d)
	return string(j)
}
