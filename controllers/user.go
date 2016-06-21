package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/julienschmidt/httprouter"
	"github.com/leonardokl/go-rest-api/models"
	"gopkg.in/mgo.v2/bson"
)

type (
	UserController struct {
		session *mgo.Session
	}
)

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (u UserController) GetAll(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	fmt.Fprint(res, "usersctrl\n")
}

func (userCtrl UserController) Create(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	user := models.User{}

	json.NewDecoder(req.Body).Decode(&user)
	user.Id = bson.NewObjectId()
	userCtrl.session.DB("go_rest_api").C("users").Insert(user)
	userJson, _ := json.Marshal(user)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(201)
	fmt.Fprintf(res, "%s", userJson)
}

func (userCtrl UserController) GetById(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	userID := params.ByName("id")
	fmt.Println("id", userID)

	if !bson.IsObjectIdHex(userID) {
		res.WriteHeader(500)
		return
	}

	userObjectId := bson.ObjectIdHex(userID)
	user := models.User{}

	if err := userCtrl.session.DB("go_rest_api").C("users").FindId(userObjectId).One(&user); err != nil {
		res.WriteHeader(404)
		return
	}

	userJSON, _ := json.Marshal(user)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200)
	fmt.Fprintf(res, "%s", userJSON)
}

func (userCtrl UserController) Delete(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	userID := params.ByName("id")

	if !bson.IsObjectIdHex(userID) {
		res.WriteHeader(404)
		return
	}

	userObjectId := bson.ObjectIdHex(userID)

	if err := userCtrl.session.DB("go_rest_api").C("users").RemoveId(userObjectId); err != nil {
		res.WriteHeader(404)
		return
	}

	fmt.Fprint(res, "delete\n")
}
