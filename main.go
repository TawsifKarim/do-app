package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const port = "8080"

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", rootHandler)

	fmt.Println("starting server at port:", port)

	err := http.ListenAndServe("localhost:"+port, r)
	if err != nil {
		log.Fatal("Something went wrong starting the server", err)
	}
}

type Person struct {
	Name    string `json:"name"`
	Purpose string `json:"purpose"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	x := Person{
		Name:    "tawsif",
		Purpose: "testing do app platform",
	}

	bt, err := json.Marshal(x)
	if err != nil {
		fmt.Println("json marshal error", err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bt)

}
