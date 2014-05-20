package main

import (
	"encoding/json"
	"net/http"
)

func renderJson(w http.ResponseWriter, i interface{}) (int, []byte) {
	b, e := json.Marshal(i)
	if e != nil {
		return returnError(e)
	}
	w.Header().Set("Content-Type", "application/json")
	return http.StatusOK, b
}

func returnError(e error) (int, []byte) {
	logger.Printf("ERROR: " + e.Error())
	return http.StatusInternalServerError, []byte(e.Error())
}
