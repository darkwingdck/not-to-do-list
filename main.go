package main

import (
    "log"
    "net/http"
    "text/template"
)

type Item struct {
    Title   string
    Done    bool
}

func main() {
    handler := func (writer http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("index.html"))
	items := map[string][]Item {
	    "Items": {
		{Title: "Don't do this", Done: false},
		{Title: "Don't do that", Done: false},
		{Title: "Don't do any of that", Done: false},
	    },
	}
	template.Execute(writer, items)
    }
    http.HandleFunc("/", handler)

    fileServer := http.FileServer(http.Dir("css"))
    http.Handle("/css/", http.StripPrefix("/css/", fileServer))

    log.Fatal(http.ListenAndServe(":8000", nil))
}
