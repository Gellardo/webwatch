package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type clock struct {
	cid string
}

var clocklist []clock

func main() {
	fmt.Println("Running:D")
	http.HandleFunc("/", root)
	http.HandleFunc("/create", create)
	http.HandleFunc("/clock", clockHandler)
	err := http.ListenAndServe(":1337", nil)
	if err != nil {
		panic(err)
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/root.html")
	t.Execute(w, nil)
}

func create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		_ = r.ParseForm()
		clockid := r.FormValue("clockid")
		clocklist = append(clocklist, clock{clockid})
		fmt.Println("added clock; list: ", clocklist)

		http.Redirect(w, r, "/clock", http.StatusSeeOther)

		return
	}

	fmt.Fprint(w, "404 you should not see this")
}

func clockHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "yeay, created a clock")
}
