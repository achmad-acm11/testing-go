package main

import (
	"fmt"
	"net/http"
)

func HelloHandler(wr http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(wr, "Hello World 1")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HelloHandler)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}