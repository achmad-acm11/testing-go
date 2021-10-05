package main

import (
	"fmt"
	"net/http"
	"os"
)

func HelloHandler(wr http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(wr, "Hello World 1")
}

func main() {
	port := os.Getenv("PORT")
	mux := http.NewServeMux()

	mux.HandleFunc("/", HelloHandler)

	server := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
