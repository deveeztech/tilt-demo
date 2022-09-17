package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", HelloHandler).Methods(http.MethodGet, http.MethodOptions)
	r.Use(mux.CORSMethodMiddleware(r))

	// Bind to a port and pass our router in
	log.Println("Starting Server... ")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	println("New request recived")
	message := "HELLO WORLD!!\n"
	_, _ = w.Write([]byte(message))
}
