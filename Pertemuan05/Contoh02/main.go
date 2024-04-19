package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main()  {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"Name": "John Wick",
			"Message": "Have a nice day",
		}

		var t, err = template.ParseFiles("views/home/index.html")
		if err != nil{
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)
	})

	fmt.Println("Starting Web Server at Port http://localhost:8080/")
	http.ListenAndServe(":7000", nil)
}
