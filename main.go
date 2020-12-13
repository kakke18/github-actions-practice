package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type payload struct {
	Text string `json:"text"`
}

func main() {
	// webhook URLを取得
	webhookURL := os.Getenv("WEBHOOK_URL")

	// 送信データを生成
	data, err := json.Marshal(payload{Text: "test"})
	if err != nil {
		fmt.Printf("[Error] failed marshal. err=%s\n", err.Error())
		return
	}

	// POST
	resp, err := http.PostForm(webhookURL, url.Values{"payload": {string(data)}})
	if err != nil {
		fmt.Printf("[Error] failed post form. err=%s\n", err.Error())
		return
	}
	defer resp.Body.Close()
}
