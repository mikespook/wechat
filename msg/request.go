package msg

import (
    "time"
    "encoding/xml"
)


type Request struct {
    XMLName xml.Name `xml:"xml"`
    msgBase // base struct
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
