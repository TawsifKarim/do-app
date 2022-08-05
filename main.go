package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	const port = "8080"
	r.HandleFunc("/", rootHandler)

	fmt.Println("starting server at port:", 8080)
	err := http.ListenAndServe("localhost:"+port, r) //notice the r in place of nil
	if err != nil {
		log.Fatal("Something went wrong starting the server", err)
	}

}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("params:", r.URL.Query()) //extract from query param

	// queryParam := r.URL.Query()

	type Lorem struct {
		Name    string `json:"name"`
		Purpose string `json:"purpose"`
	}

	x := Lorem{
		Name:    "tawsif",
		Purpose: "testing do app platform",
	}

	bt, err := json.Marshal(x)
	if err != nil {
		fmt.Println("json marshal error", err.Error())
	}

	val := r.FormValue("name")
	fmt.Println("value from body:", val)

	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bt)

}
