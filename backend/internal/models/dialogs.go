package models

type Message struct {
	From int64  `josn:"from"`
	To   int64  `json:"to"`
	Text string `json:"text"`
}
