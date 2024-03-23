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
    tmpl := template.Must(template.ParseFiles("index.html"))
    context := map[string][]Item {
	"Items": {
	    {Title: "Don't do this", Done: false},
	    {Title: "Don't do that", Done: false},
	    {Title: "Don't do any of that", Done: false},
	},
    }
    tmpl.Execute(writer, context)
}

func addItemHandler(writer http.ResponseWriter, request *http.Request) {
    title := request.PostFormValue("title")
    if len(title) == 0 {
	return
    }
    tmpl := template.Must(template.ParseFiles("item.html"))
    context := map[string]Item {
	"Item": {
	    Title: title,
	    Done: false,
	},
    }
    tmpl.Execute(writer, context)
}

func removeItemHandler(writer http.ResponseWriter, request *http.Request) {
    tmpl, _ := template.New("t").Parse("")
    tmpl.Execute(writer, nil)
}

func main() {
    http.HandleFunc("/", rootHandler)
    http.HandleFunc("/add-item/", addItemHandler)
    http.HandleFunc("/remove-item/", removeItemHandler)

    fileServer := http.FileServer(http.Dir("css"))
    http.Handle("/css/", http.StripPrefix("/css/", fileServer))

    log.Fatal(http.ListenAndServe(":8000", nil))
}
