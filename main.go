package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/vandosant/commandeer/models"
)

var commands models.Commands

func CommandHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/vnd.application+json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	c := models.Commands{
		Collection: "name",
		CommandList: []models.Command{
			{"say"},
		},
	}

	if err := json.NewEncoder(w).Encode(c); err != nil {
		panic(err)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", CommandHandler)

	fmt.Printf("Now running on port %s\n", port)
	http.ListenAndServe(":"+port, nil)

}
