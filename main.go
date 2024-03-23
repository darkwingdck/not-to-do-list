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

func rootHandler(writer http.ResponseWriter, request *http.Request) {
    template := template.Must(template.ParseFiles("index.html"))
    context := map[string][]Item {
	"Items": {
	    {Title: "Don't do this", Done: false},
	    {Title: "Don't do that", Done: false},
	    {Title: "Don't do any of that", Done: false},
	},
    }
    template.Execute(writer, context)
}

func addItemHandler(writer http.ResponseWriter, request *http.Request) {
    title := request.PostFormValue("title")
    if len(title) == 0 {
	return
    }
    template := template.Must(template.ParseFiles("item.html"))
    context := map[string]Item {
	"Item": {
	    Title: title,
	    Done: false,
	},
    }
    template.Execute(writer, context)
}

func main() {
    http.HandleFunc("/", rootHandler)
    http.HandleFunc("/add-item/", addItemHandler)

    fileServer := http.FileServer(http.Dir("css"))
    http.Handle("/css/", http.StripPrefix("/css/", fileServer))

    log.Fatal(http.ListenAndServe(":8000", nil))
}
