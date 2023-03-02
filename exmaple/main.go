package main

import (
	"github.com/elvin-tacirzade/telegobot"
	"log"
)

const (
	token  = "YourToken"
	chatID = "@YourChannel"
)

func main() {
	telegram, telegramErr := telegobot.Start(token, chatID)
	if telegramErr != nil {
		log.Fatal(telegramErr)
	}

	sendMessageErr := telegram.SendMessage("Hi...")
	if sendMessageErr != nil {
		log.Fatal(sendMessageErr)
	}

	imageURL := "https://www.freecodecamp.org/news/content/images/size/w2000/2021/10/golang.png"
	btnURL := "https://www.freecodecamp.org/news/what-is-go-programming-language"

	sendPhotoErr := telegram.SendPhoto(imageURL, "What is Golang?", "Visit Page", btnURL)
	if sendPhotoErr != nil {
		log.Fatal(sendPhotoErr)
	}
}
