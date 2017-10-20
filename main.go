package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Gellardo/webwatch/clock"
)

func main() {
	//debug
	clock.Create("1")

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
		clockid := clock.Create("")

		cookie := http.Cookie{Name: "token", Value: "this is a test token"}
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, "/clock?cid="+clockid, http.StatusSeeOther)

		return
	}

	fmt.Fprint(w, "404 you should not see this")
}

func clockHandler(w http.ResponseWriter, r *http.Request) {
	cid := r.URL.Query().Get("cid")
	auth := false
	if cookie, _ := r.Cookie("token"); cookie.Value == "this is a test token" {
		auth = true
	}

	c, _ := clock.Get(cid)

	t, err := template.ParseFiles("templates/base.html", "templates/clock.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, struct {
		Clock *clock.Clock
		Auth  bool
	}{Clock: c, Auth: auth})
}
