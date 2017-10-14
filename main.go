package main

import (
	"fmt"
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
	fmt.Fprint(w, roottemp)
}

const roottemp = `
<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<title>Test Title</title>
	<link rel="stylesheet" href="https://unpkg.com/purecss@1.0.0/build/pure-min.css">
</head>
<body style="margin: 20px;">
	<h2>Webwatch</h2>
	<p>This will be some kind of web stop watch, that is usable by multiple people.</p>
	<p>Boring stuff explaining what to do.</p>
</body>
</html>
`
