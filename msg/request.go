package msg

import (
    "time"
    "encoding/xml"
)

const (
    ReqText = "text"
    ReqLocation = "location"
    ReqImage = "image"
)

type Request struct {
    XMLName xml.Name `xml:"xml"`
    ToUserName string
    FromUserName string
    CreateTime time.Duration
    MsgType string
    Content string
    Location_X, Location_Y float32
    Scale int
    Label string
    PicUrl string
}

func Unmarshal(data []byte, req *Request) (err error) {
    if err = xml.Unmarshal(data, req); err != nil {
        return
    }
    req.CreateTime *= time.Second
    return
}
