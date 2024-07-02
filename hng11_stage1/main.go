package main

import (
	"github.com/Manuelshub/hng11_stage1/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/hello", helloHandler)

	addr := ":8080"
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("server could not start: %v", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	handlers.HelloHandler(w, r)
}
