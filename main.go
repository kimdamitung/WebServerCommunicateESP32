package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
	PORT = ":8081"
)

func index(w http.ResponseWriter, r *http.Request) {
	// connect http
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("src/index.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
	}
}

func main() {
	// connect  http web server
	log.Println("Server is running on port", PORT)
	log.Println("=====================SUCCESS=====================")
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("src/img"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("src/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("src/js"))))
	http.HandleFunc("/", index)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatalf("Error starting server %s", err)
		log.Println("=====================FAILED=====================")
	}
}