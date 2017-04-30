package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"io/ioutil"			//needed parse the response into a String
	"strings"			//needed to trim down the string to format
)

type PinterestResponse struct {
	URL   string `json:"url"`
	Count int    `json:"count"`
}

func getPinterestPins(url string) int {
	requesturl := fmt.Sprintf("http://api.pinterest.com/v1/urls/count.json?url=%s", url)

	if LogOn {
		log.Println("Requesting Pinterest Pin count:", requesturl)
	}

	response, err := http.Get(requesturl)
	if err != nil {
		log.Panic(err)
		return -1
	} else {
		body, err := ioutil.ReadAll(response.Body)
		bodystring := string(body)
		bodystring = strings.TrimLeft(strings.TrimRight(bodystring, ")"), "receiveCount(")
		body = []byte(bodystring)
		var t PinterestResponse
		err = json.Unmarshal(body, &t)
		if err != nil {
			log.Panic(err)
		}
		defer response.Body.Close()
		if LogOn {
			log.Println("Requesting Pinterest Pin count for", url, "was", t.Count)
		}
		return t.Count
	}
}
