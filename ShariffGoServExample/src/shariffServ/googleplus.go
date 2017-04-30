package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"bytes"			//needed to formulate the JSON POST request
)

type GPlusResponse []struct {
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

func getGPluses(url string) int {
	jsonReq := []byte(fmt.Sprintf("[{\"method\":\"pos.plusones.get\",\"id\":\"p\",\"params\": {\"nolog\":true,\"id\":\"%s\",\"source\":\"widget\",\"userId\":\"@viewer\",\"groupId\":\"@self\"},\"jsonrpc\":\"2.0\",\"key\":\"p\",\"apiVersion\":\"v1\"}]", url))
	req, err := http.NewRequest("POST", "https://clients6.google.com/rpc", bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json")

	if LogOn {
		log.Println("Requesting GPlus 1+ count:", string(jsonReq))
	}

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		log.Panic(err)
		return -1
	} else {
		decoder := json.NewDecoder(response.Body)
		var t GPlusResponse
		err := decoder.Decode(&t)
		if err != nil {
			log.Panic(err)
		}
		defer response.Body.Close()
		if LogOn {
			log.Println("GPlus 1+ count for", url, "was", t[0].Result.Metadata.GlobalCounts.Count)
		}
		return int(t[0].Result.Metadata.GlobalCounts.Count)
	}
}
