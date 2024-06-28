package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func headers(w http.ResponseWriter, r *http.Request) {
	for name, header := range r.Header {
		fmt.Println(name, header)
		w.Write([]byte(fmt.Sprintf("%s, %s", name, header)))
	}
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/headers", headers)
	fmt.Println("Server Listening on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
