package main

import (
	"encoding/json"
	"fibonacci/fibonacci"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func NumberHandler(w http.ResponseWriter, r *http.Request) {
	var buf []byte
	vars := mux.Vars(r)
	n, err := strconv.Atoi(vars["n"])
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	result := fibonacci.Result(n)
	res := map[string]int{
		"n":     n,
		"value": result,
	}

	buf, err = json.Marshal(res)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(buf)
	return
}

func EvenSumHandler(w http.ResponseWriter, r *http.Request) {
	var buf []byte
	vars := mux.Vars(r)
	n, err := strconv.Atoi(vars["n"])
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	result := fibonacci.EvenSumResult(n)
	res := map[string]int{
		"n":     n,
		"value": result,
	}

	buf, err = json.Marshal(res)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(buf)
	return
}

func main() {
	log.Print("Listening on port :8080")
	r := mux.NewRouter()
	r.HandleFunc("/fib/number/{n}", NumberHandler)
	r.HandleFunc("/fib/evensum/{n}", EvenSumHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
