This package can be used to send messages and photos to telegram bot.
# Installation
```
go get -u github.com/elvin-tacirzade/telegobot
```
# Usage
First we need to set the TELEGRAM_TOKEN and TELEGRAM_CHAT_ID environment variable.
```
os.Setenv("TELEGRAM_TOKEN", "XXX")
os.Setenv("TELEGRAM_CHAT_ID", "@XXX")
```
There are SendMessage() and SendPhoto() functions in the package.
### func SendMessage()
```
package main

import (
	"fmt"
	"os"
	"github.com/elvin-tacirzade/telegobot"
)

func main() {
	os.Setenv("TELEGRAM_TOKEN", "XXX")
	os.Setenv("TELEGRAM_CHAT_ID", "@XXX")
	message := "Hi Bot"
	result := telegobot.SendMessage(message)
	fmt.Println(result)
}
```
### func SendPhoto()
```
package main

import (
	"fmt"
	"os"
	"github.com/elvin-tacirzade/telegobot"
)

func main() {
	os.Setenv("TELEGRAM_TOKEN", "XXX")
	os.Setenv("TELEGRAM_CHAT_ID", "@XXX")
	photoURL := "https://go.dev/images/go_chrome_case_study.png"
	photoCaption := "Go programming language"
	btnText := "Visit Page"
	btnURL := "https://go.dev/"
	result := telegobot.SendPhoto(photoURL, photoCaption, btnText, btnURL)
	fmt.Println(result)
}
```
The SendMessage() and SendPhoto() functions return a message from a telegram bot api.

Note: Telegram only supports the following mime types of photos.
* **`image/jpeg`**
* **`image/jpg`**
* **`image/png`**
* **`image/gif`**