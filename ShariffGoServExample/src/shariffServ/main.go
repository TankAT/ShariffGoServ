package main

import (
	"os"
	"path/filepath"
	"html/template"
	"log"
	"net/http"
	"fmt"
)

var tpl *template.Template

type pageData struct {
	Title string
	Name string
	FBShareCount int
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
	http.HandleFunc("/shariff/", shariff_backend)
	http.ListenAndServe(":8080", nil)
}

func idx(w http.ResponseWriter, req *http.Request) {
	
	fmt.Println(req.URL.Path)
	
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

func shariff_backend(w http.ResponseWriter, req *http.Request){
	fb_request :=fmt.Sprintf("%s%s\n", req.Host, req.URL.Path) 
	
	counts:=getFBShares(fb_request)
	
	fmt.Println(counts)
}

func getFBShares (url string) int{
	requesturl := fmt.Sprintf("https://api.facebook.com/method/links.getStats?urls=%s&format=xml", url)
	log.Println(requesturl)
	request, _:=http.NewRequest(http.MethodGet, requesturl, nil)
	client :=&http.Client{}
	response, _:= client.Do(request)
	log.Println(response)
	
	return 4
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
