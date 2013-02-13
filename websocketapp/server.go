package main

import (
	"code.google.com/p/go.net/websocket"
	"html/template"
	"io"
	"net/http"
	"log"
	"time"
)

var count = 0

var index = template.Must(template.ParseFiles("tmpl/base.html", "tmpl/index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%v] -- %v\n", r.Method, r.URL)
	index.Execute(w, nil)
}

func echoHandler(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

func main() {
	port := ":8080"

	http.Handle("/echo", websocket.Handler(echoHandler))
	http.Handle("/script/", http.FileServer(http.Dir(".")))
	http.Handle("/css/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/", indexHandler)

	log.Printf("Listening on http://localhost%v\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Printf("Error running docs webserver: %v", err)
	}
}
