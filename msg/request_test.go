package msg

import (
    "time"
    "testing"
)

const (
    TextCase = `<xml>
<ToUserName><![CDATA[toUser]]></ToUserName>
<FromUserName><![CDATA[fromUser]]></FromUserName>
<CreateTime>1348831860</CreateTime>
<MsgType><![CDATA[text]]></MsgType>
<Content><![CDATA[this is a test]]></Content>
</xml>`

    LocationCase = `<xml>
<ToUserName><![CDATA[toUser]]></ToUserName>
<FromUserName><![CDATA[fromUser]]></FromUserName>
<CreateTime>1351776360</CreateTime>
<MsgType><![CDATA[location]]></MsgType>
<Location_X>23.134521</Location_X>
<Location_Y>113.358803</Location_Y>
<Scale>20</Scale>
<Label><![CDATA[位置信息]]></Label>
</xml>`

    ImageCase = `<xml>
<ToUserName><![CDATA[toUser]]></ToUserName>
<FromUserName><![CDATA[fromUser]]></FromUserName>
<CreateTime>1348831860</CreateTime>
<MsgType><![CDATA[image]]></MsgType>
<PicUrl><![CDATA[this is a url]]></PicUrl>
</xml>`

)

func TestRequestText(t *testing.T) {
    req := &Request{}
    if err := Unmarshal([]byte(TextCase), req); err != nil {
        t.Error(err)
        return
    }
    if req.ToUserName != "toUser" {
        t.Errorf("ToUserName: %s", req.ToUserName)
        return
    }
    if req.CreateTime != 1348831860 * time.Second {
        t.Errorf("%d", req.CreateTime)
        return
    }
}

func TestRequestLocation(t *testing.T) {
    req := Request{}
    if err := Unmarshal([]byte(LocationCase), &req); err != nil {
        t.Error(err)
        return
    }
    if req.Content != "" {
        t.Errorf("%s", req.Content)
        return
    }
    if req.Scale != 20 {
        t.Errorf("%d", req.Scale)
        return
    }
}

func TestRequestImage(t *testing.T) {
    req := Request{}
    if err := Unmarshal([]byte(ImageCase), &req); err != nil {
        t.Error(err)
        return
    }
    if req.PicUrl != "this is a url" {
        t.Errorf("%s", req.PicUrl)
        return
    }
}
