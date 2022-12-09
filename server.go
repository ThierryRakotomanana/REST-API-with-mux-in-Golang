package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8080"
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and Running...")
	})
	router.HandleFunc("/posts", getPost).Methods("GET")
	router.HandleFunc("/posts", addPost).Methods("POST")
	log.Println("server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
