package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"

	"github.com/ravirajbhatane/ping-pong-server/controllers"
)

func ping(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "pong\n")
}

func main() {
	r := httprouter.New()

	uc := controllers.NewUserConyroller(getSession())

	r.GET("/ping", ping)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	fmt.Println("Listening on port 8090")
	http.ListenAndServe(":8090", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://mongo:27017")
	if err != nil {
		panic(err)
	}

	return s
}
