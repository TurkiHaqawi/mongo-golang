package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TurkiHaqawi/mongo-golang/models"
	"github.com/julienschmidt/httprouter"
	// "github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	uc.session.DB("monog-golang").C("users").Insert(u)

	uj, err := json.Marshal(u)

	if err != nil {
		fmt.Print(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)

}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p *httprouter.Params) {

}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p *httprouter.Params) {

}
