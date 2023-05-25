package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type Error struct {
	Error string `json:"error"`
}

// JSON converts data object to JSON and encodes onto http.responseWriter
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	log.Infoln("response.JSON returned " + strconv.Itoa(statusCode))
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Error(err.Error())
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// ERROR converts error into JSON format
func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		log.Errorln(err.Error())
		JSON(w, statusCode, Error{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
