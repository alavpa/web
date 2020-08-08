package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
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

	name := r.FormValue("fname")
	email := r.FormValue("femail")
	phone := r.FormValue("fphone")
	message := r.FormValue("fmessage")

	sendMessage(name, email, phone, message)

	// http.Redirect(w, r, "/contact.html", 301)

}

func sendMessage(name string, email string, phone string, message string) {

	senderUser := os.Getenv("SENDER_USER")
	senderPass := os.Getenv("SENDER_PASS")

	body := "To: " + email + "\r\n" +
		"Subject: Contact alavpa form\r\n" +
		"\r\n NAME: " + name + "\nPHONE: " + phone + "\nMESSAGE: " + message + "\r\n"
	auth := smtp.PlainAuth("", senderUser, senderPass, "smtp.gmail.com")

	// Here we do it all: connect to our server, set up a message and send it
	to := []string{email}
	msg := []byte(body)
	err := smtp.SendMail("smtp.gmail.com:587", auth, senderUser, to, msg)
	if err != nil {
		log.Fatal(err)
	}
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
