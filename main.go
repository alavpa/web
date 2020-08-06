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
func writeForm(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)

	name := r.FormValue("fname")
	email := r.FormValue("femail")
	phone := r.FormValue("fphone")
	message := r.FormValue("fmessage")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "email = %s\n", email)
	fmt.Fprintf(w, "phone = %s\n", phone)
	fmt.Fprintf(w, "message = %s\n", message)

	http.Redirect(w, r, "contact.html", 301)

}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://www.google.com", 301)
}

func main() {
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/ping", writeForm)
	http.HandleFunc("/redirect", redirect)

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", http.StripPrefix("/", fs))

	log.Printf("Listening on %s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
