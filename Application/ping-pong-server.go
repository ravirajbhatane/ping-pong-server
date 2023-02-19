package main

import (
	"fmt"
	"net/http"
)

func pong(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

func main() {
	http.HandleFunc("/ping", pong)

	http.ListenAndServe(":8090", nil)
}
