package main

import (
	"fmt"
	"net/http"
	"os"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!")
}

func CommandServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Send a command!")
	w.WriteHeader(200)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/cmds", CommandServer)

	http.ListenAndServe(":"+port, nil)
}
