package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Text string
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	p := Page{Text: "Imersão Full Cycle"}
	t, err := template.ParseFiles("./assets/text.html")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Fprintf(w, "<title>Go</title>")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", index_handler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8000", nil)
}
