package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ping(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "pong\n")
}

func main() {
	r := httprouter.New()

	r.GET("/ping", ping)

	fmt.Println("Listening on port 8090")
	http.ListenAndServe(":8090", r)
}
