package main

import (
	"FirstGolangProject/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{num:[0-9]+}", Logger(handlers.GetNum)).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("server[net/http] method [%s] connection from [%v]", r.Method, r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}
