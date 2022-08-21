package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var req int
var messages = make(chan interface{}, 10)

func main() {

	for i := 0; i < 4; i++ {
		go worker(messages, i+1)
	}

	r := mux.NewRouter()

	r.HandleFunc("/", rootHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("starting golang server at port:", port)

	err := http.ListenAndServe("0.0.0.0:"+port, r)
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
	// req++
	// go func(reqx int) {
	// 	for i := 0; i < 500; i++ {
	// 		messages <- fmt.Sprintf("Task: %d of req %d", i, reqx)
	// 	}

	// }(req)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bt)

}

func worker(c chan interface{}, workerNo int) {

	fmt.Println("worker:", workerNo, " Ready")
	for {
		time.Sleep(time.Millisecond * 20)
		fmt.Println("worker:", workerNo, " received ", <-c)
	}
}
