package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type clock struct {
	cid    string
	create time.Time
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
	t, _ := template.ParseFiles("templates/base.html", "templates/root.html")
	t.Execute(w, nil)
}

func create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		_ = r.ParseForm()
		clockid := r.FormValue("clockid")
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
		if clock.cid == cid {
			c = &clock
		}
	}
	if c != nil {
		fmt.Fprint(w, "yeay, we have clock "+c.cid+c.create.String())
	} else {
		fmt.Fprint(w, "NO clock found, stop enumerating.")
	}
}
