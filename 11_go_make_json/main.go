package main

import "log"

func main() {
	log.Printf("test")
	j := makeMessage()
	log.Printf("makeMessage %#v", j)
	jl := getMessageList()
	log.Printf("getMessageList %#v", jl)
}

// Message jsonの中身の定義
type Message struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

// MessageList Messageの集合体
type MessageList []Message

func makeMessage() Message {
	data := Message{
		Name:    "とんかつ",
		Message: "あんぱん",
	}
	return data
}

func getMessageList() MessageList {
	data := MessageList{
		{
			Name:    "とんかつ",
			Message: "あんぱん",
		},
		{
			Name:    "えびふらい",
			Message: "カレーパン",
		},
		{
			Name:    "かきふらい",
			Message: "しょくぱん",
		},
	}
	return data
}
