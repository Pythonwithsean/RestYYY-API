package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Trans struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, fmt.Sprintf("Hello From Server \n"))
}
func payHandler(w http.ResponseWriter, r *http.Request) {

	var trans Trans
	if r.Method != http.MethodPost {
		fmt.Fprint(w, "Wrong Method Type \n")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error Reading Request Body \n", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &trans)
	if err != nil {
		fmt.Println("Error Unmarshaling Body ")
		http.Error(w, "Error Unmarshaling Body \n", http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()
	fmt.Println(trans)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data Recieved"))

}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/pay", payHandler)
	fmt.Println("Server Listening on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
