package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const view string = `<html>
	<head>
		<title>Themplate</title>
	<head>
	<body>
		<h1>Hallo</h1>
	</body>
</html>`

func main() {
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.New("main-template").Parse(view))
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusTemporaryRedirect)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/index", http.StatusTemporaryRedirect)
	})

	fmt.Println("Server Started At localhost:9090")
	http.ListenAndServe(":9090", nil)
}
