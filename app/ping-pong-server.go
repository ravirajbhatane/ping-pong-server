package main

import (
	"fmt"
	"net/http"
)

func ping(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

func main() {
	http.HandleFunc("/ping", ping)

	fmt.Println("Listening on port 8090")
	http.ListenAndServe(":8090", nil)
}
