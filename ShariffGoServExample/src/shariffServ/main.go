package main

import (
	"os"
	"path/filepath"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type pageData struct {
	Title string
	Name string
}

func init() {
	
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
            log.Fatal(err)
    }
    log.Println("Server started from:",dir)
	
	tpl = template.Must(template.ParseGlob("C:/Users/Paul/git/ShariffGoServ/ShariffGoServExample/src/shariffServ/templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", idx)
	http.Handle("/pub/", http.StripPrefix("/pub/", http.FileServer(http.Dir("C:/Users/Paul/git/ShariffGoServ/ShariffGoServExample/src/shariffServ/pub/"))))
	http.Handle("/favico.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
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

func apl(w http.ResponseWriter, req *http.Request) {
	
	data := pageData{
		Title: "Apply page",
	}
	
	var name string
	if req.Method == http.MethodPost{
		name = req.FormValue("fname")
		data.Name=name
	}
	
	err := tpl.ExecuteTemplate(w, "apply.gohtml", data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
