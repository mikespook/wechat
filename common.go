package wechat

import (
    "time"
)

const (
    Text = "text"
    Location = "location"
    Image = "image"
    News = "news"
)

type msgBase struct {
    ToUserName string
    FromUserName string
    CreateTime time.Duration
    MsgType string
    Content string
}
