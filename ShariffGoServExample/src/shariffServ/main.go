package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var LogOn bool = false

var tpl *template.Template

type pageData struct {
	Title string
	Name  string
}

func init() {
	progArgs := os.Args[1:]

	for _, progArg := range progArgs {
		if (strings.EqualFold(progArg, "--log")||strings.EqualFold(progArg, "-l")) {
			LogOn = true
		}
	}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server started from:", dir)

	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", idx)
	http.Handle("/pub/", http.StripPrefix("/pub/", http.FileServer(http.Dir("pub/"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/shariff/", shariff_backend)
	http.ListenAndServe(":80", nil)
}

func idx(w http.ResponseWriter, req *http.Request) {

	data := pageData{
		Title: "Go Title",
	}

	err := tpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
