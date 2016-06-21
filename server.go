package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/leonardokl/go-rest-api/controllers"
	"gopkg.in/mgo.v2"
)

func index(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	fmt.Fprint(res, "Rest API in Golang!\n")
}

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return session
}

func main() {
	router := httprouter.New()
	userCtrl := controllers.NewUserController(getSession())

	router.GET("/", index)
	router.GET("/users", userCtrl.GetAll)
	router.POST("/users", userCtrl.Create)
	router.GET("/users/:id", userCtrl.GetById)
	router.DELETE("/users/:id", userCtrl.Delete)

	http.ListenAndServe("localhost:3000", router)
}
