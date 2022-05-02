package telegobot

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

const botApi = "https://api.telegram.org/bot"

func connectBot() {
	if os.Getenv("TELEGRAM_TOKEN") == "" {
		log.Fatal("TELEGRAM_TOKEN not found. Please set environment variable TELEGRAM_TOKEN.")
	}
	if os.Getenv("TELEGRAM_CHAT_ID") == "" {
		log.Fatal("TELEGRAM_CHAT_ID not found. Please set environment variable TELEGRAM_CHAT_ID.")
	}
	response, err := http.Get(botApi + os.Getenv("TELEGRAM_TOKEN") + "/getMe")
	checkError(err)
	body, err := ioutil.ReadAll(response.Body)
	checkError(err)
	checkError(response.Body.Close())
	var g getMe
	checkError(json.Unmarshal(body, &g))
	if !g.Ok {
		log.Fatal(g)
	}
}

func SendMessage(message string) string {
	connectBot()
	response, err := http.PostForm(botApi+os.Getenv("TELEGRAM_TOKEN")+"/sendMessage",
		url.Values{
			"chat_id": {os.Getenv("TELEGRAM_CHAT_ID")},
			"text":    {message},
		})
	checkError(err)
	body, err := ioutil.ReadAll(response.Body)
	checkError(err)
	checkError(response.Body.Close())
	return string(body)
}

func SendPhoto(photo, caption, btnText, btnUrl string) string {
	connectBot()
	inlineBtn := map[string]interface{}{
		"inline_keyboard": []interface{}{
			[]interface{}{
				map[string]string{
					"text": btnText,
					"url":  btnUrl,
				},
			},
		},
	}
	inlineBtnByte, err := json.Marshal(inlineBtn)
	checkError(err)
	response, err := http.PostForm(botApi+os.Getenv("TELEGRAM_TOKEN")+"/sendPhoto",
		url.Values{
			"chat_id":      {os.Getenv("TELEGRAM_CHAT_ID")},
			"photo":        {photo},
			"caption":      {caption},
			"reply_markup": {string(inlineBtnByte)},
		})
	checkError(err)
	body, err := ioutil.ReadAll(response.Body)
	checkError(err)
	checkError(response.Body.Close())
	return string(body)
}
