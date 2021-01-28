package app

import (
	"encoding/json"
	"log"
	"net/http"

)

// resJson send json
func resJson(w http.ResponseWriter, iData interface{}) {

	data, err := json.Marshal(iData)

	if err != nil {
		errWriter(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)

	if err != nil {

		log.Print(err)
	}
}

// function for writing an error in responseWriter
func errWriter(w http.ResponseWriter, httpSts int, err error) {
	log.Print(err)
	http.Error(w, http.StatusText(httpSts), httpSts)
}