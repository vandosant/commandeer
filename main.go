package main

import (
	"fmt"
	"net/http"
	"os"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!")
}

func CommandHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Send a command!")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/cmds", CommandHandler)

	fmt.Printf("Now running on port %s\n", port)
	http.ListenAndServe(":"+port, nil)

}
