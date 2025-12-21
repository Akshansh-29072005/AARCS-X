package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/Akshansh-29072005/AARCS-X/backend/internals/types"
	"github.com/Akshansh-29072005/AARCS-X/backend/internals/utlis/response"

)

func New() http.HandlerFunc{
	return  func(w http.ResponseWriter, r *http.Request) {

		slog.Info("Creating a student")

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF){
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil{
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		//request validation
		
		response.WriteJson(w, http.StatusCreated, map[string] string{"success" : "OK"})
	}
}