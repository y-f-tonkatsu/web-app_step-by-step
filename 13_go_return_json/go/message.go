package main

// Message jsonの中身の定義
type Message struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

// MessageList Messageの集合体
type MessageList []Message

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
