// Package telegobot allows to connect to telegram and send message and photo via bot
package telegobot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type (
	// Telegobot interface includes two functions. SendMessage and SendPhoto
	Telegobot interface {
		SendMessage(message string) error
		SendPhoto(photo, caption, btnText, btnUrl string) error
	}

	// telegobot struct stores TelegramToken and TelegramChatID
	telegobot struct {
		// TelegramToken is API Token for bot
		TelegramToken string
		// TelegramChatID is channel ID. For example @YourChannel
		TelegramChatID string
	}
)

// apiURL describes telegram api url for bot
const apiURL = "https://api.telegram.org/bot"

// Start function checks telegram token and returns Telegobot interface and error
func Start(telegramToken, telegramChatID string) (Telegobot, error) {
	addr := getAddr(telegramToken, "getMe")
	response, responseErr := http.Get(addr)
	if responseErr != nil {
		return nil, fmt.Errorf("failed to check token: %v", responseErr)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("telegram token is invalid")
	}

	return &telegobot{
		TelegramToken:  telegramToken,
		TelegramChatID: telegramChatID,
	}, nil
}

// SendMessage function sends message to telegram channel via bot
func (t *telegobot) SendMessage(message string) error {
	addr := getAddr(t.TelegramToken, "sendMessage")

	data := url.Values{
		"chat_id": {t.TelegramChatID},
		"text":    {message},
	}

	response, responseErr := http.PostForm(addr, data)
	if responseErr != nil {
		return fmt.Errorf("failed to get response: %v", responseErr)
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("response status code is incorrect. expected: %v, actual: %v", http.StatusOK, response.StatusCode)
	}

	return nil
}

// SendPhoto function sends photo to telegram channel via bot
func (t *telegobot) SendPhoto(photoURL, caption, btnText, btnURL string) error {
	addr := getAddr(t.TelegramToken, "sendPhoto")
	replyMarkup := ""

	if btnText != "" && btnURL != "" {
		inlineBtn, inlineBtnErr := createInlineBtn(btnText, btnURL)
		if inlineBtnErr != nil {
			return inlineBtnErr
		}
		replyMarkup = inlineBtn
	}

	data := url.Values{
		"chat_id":      {t.TelegramChatID},
		"photo":        {photoURL},
		"caption":      {caption},
		"reply_markup": {replyMarkup},
	}

	response, responseErr := http.PostForm(addr, data)
	if responseErr != nil {
		return fmt.Errorf("failed to get response: %v", responseErr)
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("response status code is incorrect. expected: %v, actual: %v", http.StatusOK, response.StatusCode)
	}

	return nil
}

// getAddr function joins apiUrL, telegramToken and path
func getAddr(telegramToken, path string) string {
	return fmt.Sprintf("%s%s/%s", apiURL, telegramToken, path)
}

// createInlineBtn function creates inlineBtn for SendPhoto function
func createInlineBtn(btnText, btnURL string) (string, error) {
	inlineBtn := map[string]interface{}{
		"inline_keyboard": []interface{}{
			[]interface{}{
				map[string]string{
					"text": btnText,
					"url":  btnURL,
				},
			},
		},
	}

	inlineBtnByte, inlineBtnByteErr := json.Marshal(inlineBtn)
	if inlineBtnByteErr != nil {
		return "", fmt.Errorf("failed to marshal inlineBtn: %v", inlineBtnByteErr)
	}

	return string(inlineBtnByte), nil
}
