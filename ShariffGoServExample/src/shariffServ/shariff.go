package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var LogOn bool = false

type ShariffResponse struct {
	Facebook    int `json:"facebook"`
	Googleplus  int `json:"googleplus"`
	Linkedin    int `json:"linkedin"`
	Stumbleupon int `json:"stumbleupon"`
	Pinterest   int `json:"pinterest"`
}

func shariff(w http.ResponseWriter, req *http.Request) {
	var t ShariffResponse
	request := fmt.Sprintf("http://%s", strings.TrimLeft(strings.TrimLeft(req.Host, "http://"), "www."))

	t.Facebook = getFBShares(request)
	t.Googleplus = getGPluses(request)
	t.Pinterest = getPinterestPins(request)
	t.Linkedin = getLinkedInShare(request)
	t.Stumbleupon = getStumbleUponViews(request)

	json.NewEncoder(w).Encode(t)
}

func shariffLogging(logging bool){
	LogOn=logging
}