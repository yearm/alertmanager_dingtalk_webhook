package types

type MarkdownMessage struct {
	MsgType  string    `json:"msgtype"`
	Markdown *Markdown `json:"markdown"`
	At       *At       `json:"at"`
}

type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}
