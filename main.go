package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/vandosant/commandeer/models"
	"gopkg.in/mgo.v2"
)

var commands models.Commands

var (
	mongoURL        = os.Getenv("mongo_url")
	mongoUser       = os.Getenv("mongo_user")
	mongoPassword   = os.Getenv("mongo_password")
	mongoCollection = os.Getenv("mongo_collection")
	mongoSession    *mgo.Session
	database        *mgo.Database
	collection      *mgo.Collection
	err             error
)

func CommandHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/vnd.application+json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	c := models.Commands{
		Collection: "name",
		CommandList: []models.Command{
			{"say"},
		},
	}

	cmd := exec.Command("say", "hello")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	if err := json.NewEncoder(w).Encode(c); err != nil {
		panic(err)
	}
}

func main() {
	addr := fmt.Sprintf("mongodb://%s:%s@%s", mongoUser, mongoPassword, mongoURL)
	if mongoSession, err = mgo.Dial(addr); err != nil {
		log.Fatal(err)
	}
	defer mongoSession.Close()

	database = mongoSession.DB("")
	collection = database.C(mongoCollection)
	fmt.Printf("Collection: %+v", collection)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", CommandHandler)

	fmt.Printf("Now running on port %s\n", port)
	http.ListenAndServe(":"+port, nil)

}
