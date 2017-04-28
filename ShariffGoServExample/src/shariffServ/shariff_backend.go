package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type ShariffResponse struct {
	Facebook    int `json:"facebook"`
	Googleplus  int `json:"googleplus"`
	Linkedin    int `json:"linkedin"`
	Stumbleupon int `json:"stumbleupon"`
	Pinterest   int `json:"pinterest"`
}

func shariff_backend(w http.ResponseWriter, req *http.Request) {
	var t ShariffResponse
	request := fmt.Sprintf("http://www.%s", strings.TrimLeft(strings.TrimLeft(req.Host, "http://"), "www."))

	t.Facebook = getFBShares(request)
	t.Googleplus = getGPluses(request)
	t.Pinterest = getPinterestPins(request)
	t.Linkedin = getLinkedInShare(request)
	t.Stumbleupon = getStumbleUponViews(request)

	json.NewEncoder(w).Encode(t)
}
