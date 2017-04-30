package main

import (
	"encoding/json" //supplies json encoding and decoding
	"fmt"           //supplies basic formating functionality
	"log"           //supplies basic logging capabilites
	"net/http"      //supplies http functionality
)

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

func getFBShares(url string) int {
	requesturl := fmt.Sprintf("http://graph.facebook.com/?fields=og_object%%7Blikes.summary(true).limit(0)%%7D,share&id=%s", url)
	if LogOn {
		log.Println("Requesting FB Share count:", requesturl)
	}
	response, err := http.Get(requesturl)
	if err != nil {
		log.Panic(err)
		return -1
	} else {
		decoder := json.NewDecoder(response.Body)
		var t FBResponse
		err := decoder.Decode(&t)
		if err != nil {
			log.Panic(err)
		}
		defer response.Body.Close()
		if LogOn {
			log.Println("Facebook share count for", url, "was", t.Share.ShareCount)
		}
		return t.Share.ShareCount
	}
}
