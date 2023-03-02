# Telegobot
This package can be used to send messages and photos to telegram channel via bot.

[![Go Reference](https://pkg.go.dev/badge/github.com/elvin-tacirzade/telegobot.svg)](https://pkg.go.dev/github.com/elvin-tacirzade/telegobot)
## Installation
```
go get -u github.com/elvin-tacirzade/telegobot
```
## Usage
First we call the Start() function. The Start() function takes the following parameters:
1. `telegramToken` - Declare API Token for bot.
2. `telegramChatID` - Declare channel ID. For example @YourChannel.

This function returns interface and error. See the [example](https://github.com/elvin-tacirzade/telegobot/blob/main/exmaple/main.go) subdirectory for more information.

Note: Telegram only supports the following mime types of photos.
* **`image/jpeg`**
* **`image/jpg`**
* **`image/png`**
* **`image/gif`**