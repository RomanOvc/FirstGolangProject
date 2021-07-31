package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{num:[0-9]+}", Logger(getNum)).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getNum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	par, err := strconv.Atoi(params["num"])
	if err == nil {
		// fmt.Println(par)
	}
	type Resp struct {
		Result int
	}

	res := Resp{
		Result: par * par,
	}
	// fmt.Println(res)

	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("server[net/http] method [%s] connection from [%v]", r.Method, r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}
