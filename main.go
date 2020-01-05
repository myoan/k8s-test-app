package main

import (
	"log"
	"net/http"
)

type User struct {
	ID int
	Name string
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hogehoge"))
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
