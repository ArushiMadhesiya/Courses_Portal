package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//Let's start registering a couple of URL paths and handlers:
	fmt.Println("server starting")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	// now creating a serverin golang
	//http.ListenAndServe(":4000",r)
	// it may throw error 
	log.Fatal(http.ListenAndServe(":4000",r))
	// we do go build . instead of go run main.go
}
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>raboo here</h1>"))
}
