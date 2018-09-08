package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func serveFiles(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("public"))
	http.StripPrefix("/public/", fs)
}

func main() {
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/ping", hello)

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", http.StripPrefix("/", fs))

	log.Printf("Listening on %s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
