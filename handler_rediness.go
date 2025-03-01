package main

import (
	"net/http"
)

func handlerRediness(w http.ResponseWriter, r *http.Request) {

	responseWithJSON(w, 200, struct{}{})

}
