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
	fmt.Fprintln(w, "<html><head></head><body><b>Hola me llamo Alex</b></body></html>")
}
func main() {
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", hello)
	log.Printf("Listening on %s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
