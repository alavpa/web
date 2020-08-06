package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
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
	auth := smtp.PlainAuth("", "alavpa@outlook.es", "@i1n5m2a1!", "smtp-mail.outlook.com")

	// Here we do it all: connect to our server, set up a message and send it
	to := []string{"alavpa@gmail.com"}
	msg := []byte(name + "\r\n" +
		email + "\r\n" +
		phone + "\r\n" +
		"\r\n" +
		message + "\r\n")
	err := smtp.SendMail("smtp-mail.outlook.com:587", auth, "alavpa@outlook.es", to, msg)
	if err != nil {
		log.Fatal(err)
	}

}

func writeMessage(name string, email string, phone string, message string) {
	t := time.Now()
	filename := t.Format("20060102150405") + ".txt"

	f, err := os.Create("public/messages/" + filename)
	check(err)
	defer f.Close()

	writeWord(f, name+"\n")
	writeWord(f, email+"\n")
	writeWord(f, phone+"\n")
	writeWord(f, message+"\n")
	f.Sync()

}

func writeWord(f *os.File, word string) {
	_, err := f.WriteString(word + "\n")
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
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
