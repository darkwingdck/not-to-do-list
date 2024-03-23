package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

type Task struct {
    Title   string
    Done    bool
}

func rootHandler(writer http.ResponseWriter, request *http.Request) {
    tmpl := template.Must(template.ParseFiles("index.html"))
    context := map[string][]Task {
	"Tasks": {
	    {Title: "Don't do this", Done: false},
	    {Title: "Don't do that", Done: false},
	    {Title: "Don't do any of that", Done: false},
	},
    }
    tmpl.Execute(writer, context)
}

func addTaskHandler(writer http.ResponseWriter, request *http.Request) {
    title := request.PostFormValue("title")
    if len(title) == 0 {
	return
    }
    tmpl := template.Must(template.ParseFiles("task.html"))
    context := map[string]Task {
	"Task": {
	    Title: title,
	    Done: false,
	},
    }
    tmpl.Execute(writer, context)
}

func removeTaskHandler(writer http.ResponseWriter, request *http.Request) {
    tmpl, _ := template.New("t").Parse("")
    tmpl.Execute(writer, nil)
}

func markTaskHandler(writer http.ResponseWriter, request *http.Request) {
    title := request.PostFormValue("title")
    done := request.PostFormValue("done") == "true"
    tmpl := template.Must(template.ParseFiles("task.html"))
    context := map[string]Task {
	"Task": {
	    Title: title,
	    Done: done,
	},
    }
    tmpl.Execute(writer, context)
}

func faviconHandler(writer http.ResponseWriter, request *http.Request) {
    http.ServeFile(writer, request, "./favicon.ico")
}

func main() {
    http.HandleFunc("/", rootHandler)
    http.HandleFunc("/add-task/", addTaskHandler)
    http.HandleFunc("/remove-task/", removeTaskHandler)
    http.HandleFunc("/mark-task/", markTaskHandler)
    http.HandleFunc("/favicon.ico", faviconHandler)

    fileServer := http.FileServer(http.Dir("css"))
    http.Handle("/css/", http.StripPrefix("/css/", fileServer))

    port := os.Getenv("PORT")

    if port == "" {
	port = "3000"
    }

    log.Fatal(http.ListenAndServe("0.0.0.0:" + port, nil))
}
