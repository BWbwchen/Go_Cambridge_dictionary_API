package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

var TOKEN string = "YOUR-TELEGRAM-TOKEN"

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}
type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}
type Chat struct {
	Id int `json:"id"`
}

func SendTextToTelegramChat(chatId int, text string) (string, error) {
	var telegramApi string = "https://api.telegram.org/bot" + TOKEN + "/sendMessage"
	parse_mode := "Markdown"
	response, err := http.PostForm(
		telegramApi,
		url.Values{
			"chat_id":    {strconv.Itoa(chatId)},
			"text":       {text},
			"parse_mode": {parse_mode},
		})

	if err != nil {
		log.Printf("error when posting text to the chat: %s", err.Error())
		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		log.Printf("error in parsing telegram answer %s", errRead.Error())
		return "", err
	}
	bodyString := string(bodyBytes)

	return bodyString, nil
}
