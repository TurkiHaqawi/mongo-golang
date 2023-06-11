package main

import (
	"net/http"

	"log"
	"os"

	"github.com/TurkiHaqawi/mongo-golang/controllers"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSessions())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)

}

func getSessions() *mgo.Session {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s, err := mgo.Dial(os.Getenv("MONGO_URL"))
	if err != nil {
		panic(err)
	}
	return s
}
