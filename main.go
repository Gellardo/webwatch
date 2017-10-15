package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	fmt.Println("Running:D")
	http.HandleFunc("/", root)
	err := http.ListenAndServe(":1337", nil)
	if err != nil {
		panic(err)
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/root.html")
	t.Execute(w, nil)
}
