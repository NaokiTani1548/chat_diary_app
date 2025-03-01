package models

type Chat struct {
	Chat_id     int    `json:"chat_id"`
	Chat_detail string `json:"chat_detail"`
	Author_id   int    `json:"author_id"`
}

type PostChat struct {
	Chat_detail string `json:"chat_detail"`
	// Author_id   int    `json:"author_id"`
}
