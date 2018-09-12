package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}
func hello(w http.ResponseWriter, r *http.Request) {

	c := make(chan string)

	go func() {
		for {
			time.Sleep(time.Second)
			//c <- "pong"
			fmt.Fprintln(w, "pong")
		}
	}()

	//	for i := 0; i < 10; i++ {
	//		fmt.Fprintln(w, <-c)
	//	}

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
