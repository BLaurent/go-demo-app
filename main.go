package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", hello)
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = "8080"
	}
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "hello, world!")
}
