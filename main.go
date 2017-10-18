package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gellardo/webwatch/rand"
)

type clock struct {
	Cid    string
	Create time.Time
}

var clocklist []clock

func main() {
	fmt.Println("Running:D")
	http.HandleFunc("/", root)
	http.HandleFunc("/clock", clockHandler)
	http.HandleFunc("/clock/create", clockCreate)
	err := http.ListenAndServe(":1337", nil)
	if err != nil {
		panic(err)
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/base.html", "templates/root.html")
	t.Execute(w, nil)
}

func clockCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		clockid := rand.RandomString(5)
		clocklist = append(clocklist, clock{clockid, time.Now()})
		fmt.Println("added clock; list: ", clocklist)

		http.Redirect(w, r, "/clock?cid="+clockid, http.StatusSeeOther)

		return
	}

	fmt.Fprint(w, "404 you should not see this")
}

func clockHandler(w http.ResponseWriter, r *http.Request) {
	cid := r.URL.Query().Get("cid")
	var c *clock
	for _, clock := range clocklist {
		if clock.Cid == cid {
			c = &clock
		}
	}
	t, err := template.ParseFiles("templates/base.html", "templates/clock.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, c)
}
