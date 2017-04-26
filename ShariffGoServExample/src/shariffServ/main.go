package main

import (
//	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
//	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *template.Template

type pageData struct {
	Title        string
	Name         string
	FBShareCount int
}

func init() {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server started from:", dir)

	tpl = template.Must(template.ParseGlob("C:/Users/Paul/git/ShariffGoServ/ShariffGoServExample/src/shariffServ/templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", idx)
	http.Handle("/pub/", http.StripPrefix("/pub/", http.FileServer(http.Dir("C:/Users/Paul/git/ShariffGoServ/ShariffGoServExample/src/shariffServ/pub/"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/shariff/", shariff_backend)
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

func shariff_backend(w http.ResponseWriter, req *http.Request) {
	fb_request := fmt.Sprintf("%s%s\n", req.Host, req.URL.Path)

	counts := getFBShares(fb_request)

	fmt.Println(counts)
}

func getFBShares(url string) int {
	
	type AutoGenerated struct {
	OgObject struct {
		Likes struct {
			Data []interface{} `json:"data"`
			Summary struct {
				TotalCount int `json:"total_count"`
			} `json:"summary"`
		} `json:"likes"`
		ID string `json:"id"`
	} `json:"og_object"`
	Share struct {
		CommentCount int `json:"comment_count"`
		ShareCount int `json:"share_count"`
	} `json:"share"`
	ID string `json:"id"`
}

	url = "http://mephisto-mori.eu:9001/"
	url = "https://messycow.com"
	
	requesturl := fmt.Sprintf("http://graph.facebook.com/?fields=og_object%%7Blikes.summary(true).limit(0)%%7D,share&id=%s", url)
	log.Println(requesturl)

	response, err := http.Get(requesturl)
	if err != nil {
		log.Fatal(err)
		return -1
	} else {
		decoder := json.NewDecoder(response.Body)
		var t AutoGenerated
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()
		log.Println(t.Share.ShareCount)

		//		buf := new(bytes.Buffer)
		//		defer response.Body.Close()
		//		_, err := io.Copy(buf, response.Body)
		//		if err != nil {
		//			log.Fatal(err)
		//		}
		//
		//		decoder := json.NewDecoder(buf)
		//
		//		var v map[string]string
		//		err = json.Unmarshal([]byte(buf.String()), &v)
		//		if err != nil {
		//			log.Fatal(err)
		//		}
		//		fmt.Println(v)

		return 4
	}
}

func getGPluses() {

}

func apl(w http.ResponseWriter, req *http.Request) {

	data := pageData{
		Title: "Apply page",
	}

	var name string
	if req.Method == http.MethodPost {
		name = req.FormValue("fname")
		data.Name = name
	}

	err := tpl.ExecuteTemplate(w, "apply.gohtml", data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
