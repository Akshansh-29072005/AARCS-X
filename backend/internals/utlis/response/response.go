package response

import (
	"encoding/json"
	"net/http"
)

const(
	StatusOK = "OK"
	StatusError = "Error"
)

type Response struct{
	Status string `json:"status"`
	Error string `json:"error"`
}


func WriteJson(w http.ResponseWriter, status int, data interface{}) error{

	w.Header().Set("content-type" , "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response{
	return Response{
		Status: StatusError,
		Error: err.Error(),
	}
}