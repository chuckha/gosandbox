package main

import (
	"html/template"
	"net/http"
)

type Message struct {
	Author string
	Message string
}

var Messages []Message

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	if r.Method == "POST" {
		m := Message{r.RemoteAddr, r.FormValue("message")}
		Messages = append(Messages, m)
		t.Execute(w, Messages)
	} else {
		t.Execute(w, Messages)
	}

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
