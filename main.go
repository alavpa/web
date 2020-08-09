package main

import (
	"fmt"
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

	err1 := r.ParseForm()

	if err1 != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err1)
		return
	}

	name := r.FormValue("fname")
	email := r.FormValue("femail")
	phone := r.FormValue("fphone")
	message := r.FormValue("fmessage")

	err2 := sendMessage(name, email, phone, message)

	if err2 != nil {
		fmt.Fprintf(w, "Send Email err: %v", err2)
		return
	}

	http.Redirect(w, r, "/contact.html", 301)

}

func sendMessage(name string, email string, phone string, message string) error {

	senderUser := os.Getenv("SENDER_USER")
	senderPass := os.Getenv("SENDER_PASS")

	body := "To: alavpa@gmail.com\r\n" +
		"Subject: Contact alavpa form\r\n" +
		"\r\nNAME: " + name + "\nEMAIL: " + email + "\nPHONE: " + phone + "\nMESSAGE: " + message + "\r\n"
	auth := smtp.PlainAuth("", senderUser, senderPass, "smtp.gmail.com")

	// Here we do it all: connect to our server, set up a message and send it
	to := []string{"alavpa@gmail.com"}
	msg := []byte(body)
	return smtp.SendMail("smtp.gmail.com:587", auth, senderUser, to, msg)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://www.google.com", 301)
}

func main() {
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/send", writeForm)
	http.HandleFunc("/redirect", redirect)

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", http.StripPrefix("/", fs))

	log.Printf("Listening on %s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
