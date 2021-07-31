package handlers

import (
	"FirstGolangProject/model"
	"FirstGolangProject/usecase"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func GetNum(w http.ResponseWriter, r *http.Request) {
	var (
		err error
	)
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			resErr := errors.New("")
			err1 := model.ErrorResp{Err: resErr.Error()}
			bytes, err := json.Marshal(err1)
			if err != nil {
				log.Println(err)
			}
			w.Write(bytes)
		}
	}()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	par, err := strconv.Atoi(params["num"])
	if err != nil {
		log.Println(err)
		err = errors.New("long int, fuck you!")
		return
	}

	res, err := usecase.SqrtMain(par)
	if err != nil {
		return
	}
	// fmt.Println(res)

	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
