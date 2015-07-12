package main

import (
	"fmt"
	"net/http"
	"os"
)

func CommandHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Fprint(w, "Send a command!")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", CommandHandler)
	http.HandleFunc("/cmds", CommandHandler)

	fmt.Printf("Now running on port %s\n", port)
	http.ListenAndServe(":"+port, nil)

}
