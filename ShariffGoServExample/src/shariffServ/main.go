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
	
	tpl = template.Must(template.ParseGlob("C:/Users/Paul/GoLang_WS/WebApp1/src/shariffServ/templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", idx)
	http.HandleFunc("/about", abt)
	http.HandleFunc("/contact", ctc)
	http.HandleFunc("/apply", apl)
	http.Handle("/favico.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func idx(w http.ResponseWriter, req *http.Request) {
	
	data := pageData{
		Title: "Index page",
	}
	
	err := tpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
func abt(w http.ResponseWriter, req *http.Request) {
	
	data := pageData{
		Title: "About page",
	}
	
	err := tpl.ExecuteTemplate(w, "about.gohtml", data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
func ctc(w http.ResponseWriter, req *http.Request) {
	
	data := pageData{
		Title: "Contact page",
	}
	
	err := tpl.ExecuteTemplate(w, "contact.gohtml", data)
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
