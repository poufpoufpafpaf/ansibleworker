package main

import (
	"fmt"
	"log"
	"net/http"
)

type Result struct {
	Username  string
	Platform  string
	Valid     bool
	Available bool
}

func main() {
	// Hello world, the web server

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/ansible-playbook", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello, world!\n")
}

func handler(w http.ResponseWriter, req *http.Request) {
	playbook := req.URL.Query().Get("playbook")
	fmt.Fprint(w, playbook)
}
