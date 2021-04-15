package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func echo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["msg"]

	fmt.Fprint(w, key)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/echo/{msg}", echo)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Listening on port 10000")
	handleRequests()
}
