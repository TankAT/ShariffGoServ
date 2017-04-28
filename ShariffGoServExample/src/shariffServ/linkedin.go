package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type LinkedInResponse struct {
	Count       int    `json:"count"`
	FCnt        string `json:"fCnt"`
	FCntPlusOne string `json:"fCntPlusOne"`
	URL         string `json:"url"`
}

func getLinkedInShare(url string) int {
	requesturl := fmt.Sprintf("https://www.linkedin.com/countserv/count/share?url=%s&format=json", url)

	if LogOn {
		log.Println("Requesting LinkedInShare count:", requesturl)
	}

	response, err := http.Get(requesturl)
	if err != nil {
		log.Panic(err)
		return -1
	} else {
		decoder := json.NewDecoder(response.Body)
		var t LinkedInResponse
		err := decoder.Decode(&t)
		if err != nil {
			log.Panic(err)
		}
		defer response.Body.Close()
		if LogOn {
			log.Println("LinkedIn share count for", url, "was", t.Count)
		}
		return t.Count
	}
}
