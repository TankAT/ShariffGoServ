package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"io/ioutil"
	"strings"
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
	request := fmt.Sprintf("%s%s\n", req.Host, req.URL.Path)

	request = "http://z0r.de"
	request = "http://www.cookingclassy.com/sesame-noodles-with-chicken-and-broccoli/"

	counts := getFBShares(request)

	fmt.Println(counts)

	counts = getGPluses(request)

	fmt.Println(counts)
	
	counts = getPinterestPins(request)
	
	fmt.Println(counts)
	
	counts = getLinkedInShare(request)
	
	fmt.Println(counts)
}

func getFBShares(url string) int {
	type FBResponse struct {
		OgObject struct {
			Likes struct {
				Data    []interface{} `json:"data"`
				Summary struct {
					TotalCount int `json:"total_count"`
				} `json:"summary"`
			} `json:"likes"`
			ID string `json:"id"`
		} `json:"og_object"`
		Share struct {
			CommentCount int `json:"comment_count"`
			ShareCount   int `json:"share_count"`
		} `json:"share"`
		ID string `json:"id"`
	}

	requesturl := fmt.Sprintf("http://graph.facebook.com/?fields=og_object%%7Blikes.summary(true).limit(0)%%7D,share&id=%s", url)
	log.Println("Requesting FB Share count:", requesturl)

	response, err := http.Get(requesturl)
	if err != nil {
		log.Fatal(err)
		return -1
	} else {
		decoder := json.NewDecoder(response.Body)
		var t FBResponse
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()
		log.Println("Facebook share count for",url,"was",t.Share.ShareCount)
		return t.Share.ShareCount
	}
}

func getGPluses(url string) int {
	type GPlusResponse [] struct {
		ID     string `json:"id"`
		Result struct {
			Kind          string `json:"kind"`
			ID            string `json:"id"`
			IsSetByViewer bool   `json:"isSetByViewer"`
			Metadata      struct {
				Type         string `json:"type"`
				GlobalCounts struct {
					Count float64 `json:"count"`
				} `json:"globalCounts"`
			} `json:"metadata"`
			Abtk string `json:"abtk"`
		} `json:"result"`
	}

	jsonReq := []byte(fmt.Sprintf("[{\"method\":\"pos.plusones.get\",\"id\":\"p\",\"params\": {\"nolog\":true,\"id\":\"%s\",\"source\":\"widget\",\"userId\":\"@viewer\",\"groupId\":\"@self\"},\"jsonrpc\":\"2.0\",\"key\":\"p\",\"apiVersion\":\"v1\"}]", url))
	req, err := http.NewRequest("POST", "https://clients6.google.com/rpc", bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
		return -1
	} else {
		decoder := json.NewDecoder(response.Body)
		var t GPlusResponse
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()
		log.Println("GPlus 1+ count for",url,"was",t[0].Result.Metadata.GlobalCounts.Count)
		return int(t[0].Result.Metadata.GlobalCounts.Count)
	}
}

func getPinterestPins(url string) int {
type PinterestResponse struct {
	URL string `json:"url"`
	Count int `json:"count"`
}

	requesturl := fmt.Sprintf("http://api.pinterest.com/v1/urls/count.json?url=%s", url)
	log.Println("Requesting Pinterest Pin count:", requesturl)

	response, err := http.Get(requesturl)
	if err != nil {
		log.Fatal(err)
		return -1
	} else {
		body, err := ioutil.ReadAll(response.Body)
		bodystring := string(body)	
		bodystring = strings.TrimLeft(strings.TrimRight(bodystring,")"),"receiveCount(")
		body = []byte(bodystring)
		var t PinterestResponse
		err = json.Unmarshal(body, &t)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()
		log.Println("Pinterest Pin count for",url,"was",t.Count)
		return t.Count
	}
}

func getLinkedInShare (url string) int {
	type LinkedInResponse struct {
	Count int `json:"count"`
	FCnt string `json:"fCnt"`
	FCntPlusOne string `json:"fCntPlusOne"`
	URL string `json:"url"`
}

	requesturl := fmt.Sprintf("https://www.linkedin.com/countserv/count/share?url=%s&format=json", url)
	log.Println("LinkedInShare count:", requesturl)

	response, err := http.Get(requesturl)
	if err != nil {
		log.Fatal(err)
		return -1
	} else {
		decoder := json.NewDecoder(response.Body)
		var t LinkedInResponse
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()
		log.Println("LinkedIn share count for",url,"was",t.Count)
		return t.Count
	}
}
