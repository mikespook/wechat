package wechat

import (
	"encoding/xml"
	"io"
)

type Message struct {
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
	Content      string
	MsgId        string
}

func NewMessage(r io.Reader) (*Message, error) {
	d := xml.NewDecoder(r)
	var data Message
	if err := d.Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
