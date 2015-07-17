package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/vandosant/commandeer/models"
)

var commands models.Commands
var command models.Command

func CommandHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/vnd.application+json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	command.Name = "say"
	commands.Commands = append(commands.Commands, command)

	if err := json.NewEncoder(w).Encode(commands); err != nil {
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
