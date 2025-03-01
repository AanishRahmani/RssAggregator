package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithError(w http.ResponseWriter, code int, msg string) {

	if code > 499 {
		log.Println("responding with 5XX error:", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	responseWithJSON(w, code, errResponse{
		Error: msg,
	})

}

func responseWithJSON(w http.ResponseWriter, code int, paylod interface{}) {

	data, err := json.Marshal(paylod)
	if err != nil {
		log.Printf("failed to marshal json response%v", paylod)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)

}
