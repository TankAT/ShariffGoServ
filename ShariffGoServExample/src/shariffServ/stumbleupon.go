package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type SUResponse struct {
		Result struct {
			URL        string `json:"url"`
			InIndex    bool   `json:"in_index"`
			Publicid   string `json:"publicid"`
			Views      int    `json:"views"`
			Title      string `json:"title"`
			Thumbnail  string `json:"thumbnail"`
			ThumbnailB string `json:"thumbnail_b"`
			SubmitLink string `json:"submit_link"`
			BadgeLink  string `json:"badge_link"`
			InfoLink   string `json:"info_link"`
		} `json:"result"`
		Timestamp int  `json:"timestamp"`
		Success   bool `json:"success"`
	}

func getStumbleUponViews(url string) int {
	
	requesturl := fmt.Sprintf("http://www.stumbleupon.com/services/1.01/badge.getinfo?url=%s", url)
	if(LogOn){
	log.Println("Requesting StumbeUpon view count:", requesturl)
	}
	response, err := http.Get(requesturl)
	if err != nil {
		log.Panic(err)
		return -1
	} else {
		decoder := json.NewDecoder(response.Body)
		var t SUResponse
		err := decoder.Decode(&t)
		if err != nil {
			log.Panic(err)
		}
		defer response.Body.Close()
		if(LogOn){
		log.Println("StumbeUpon view count for", url, "was", t.Result.Views)
		}
		return t.Result.Views
	}
}
