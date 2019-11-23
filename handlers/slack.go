package handlers

import (
	"encoding/json"
	"komodorIntegrations/db"
	"time"
)

// SlackClient API for message sending
type SlackHandler struct {
	HandlerFunc func(string, *db.Details) error
	TokenURL    string
}

type SlackMessage struct {
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Fallback   string  `json:"fallback"`
	Color      string  `json:"color"`
	Pretext    string  `json:"pretext"`
	AuthorName string  `json:"author_name"`
	AuthorLink string  `json:"author_link"`
	AuthorIcon string  `json:"author_icon"`
	Title      string  `json:"title"`
	TitleLink  string  `json:"title_link"`
	Text       string  `json:"text"`
	ImageURL   string  `json:"image_url"`
	ThumbURL   string  `json:"thumb_url"`
	Footer     string  `json:"footer"`
	FooterIcon string  `json:"footer_icon"`
	Ts         int64   `json:"ts"`
	Fields     []Field `json:"fields"`
}

type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

func SetSlackMessage(req *Request) ([]byte, error) {
	attachment := Attachment{
		Color:      "#36a64f",
		Pretext:    "Notification Alert from Komodor",
		AuthorName: "Komodor.io",
		AuthorLink: "https://komodor.io/",
		Title:      req.Title,
		TitleLink:  req.Link,
		Text:       req.Details,
		ImageURL:   "http://my-website.com/path/to/image.jpg",
		ThumbURL:   "http://example.com/path/to/thumb.png",
		Footer:     "Komodor API",
		FooterIcon: "https://komodor.io/wp-content/uploads/2019/11/LogoMakr_6ZPQCs.png",
		Ts:         time.Now().Unix(),
	}

	message := SlackMessage{
		Attachments: []Attachment{attachment},
	}

	return json.Marshal(message)
}
