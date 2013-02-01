package wechat

import (
    "time"
    "testing"
)

const (
    RespTextCase = `<xml><ToUserName>toUser</ToUserName><FromUserName>fromUser</FromUserName><CreateTime>12345678</CreateTime><MsgType>text</MsgType><Content>content</Content><Articles></Articles><FuncFlag>0</FuncFlag></xml>`

    RespNewsCase = `<xml><ToUserName>toUser</ToUserName><FromUserName>fromUser</FromUserName><CreateTime>12345678</CreateTime><MsgType>news</MsgType><Content></Content><ArticleCount>2</ArticleCount><Articles><item><Title>title1</Title><Description>description1</Description><PicUrl>picurl</PicUrl><Url>url</Url></item><item><Title>title</Title><Description>description</Description><PicUrl>picurl</PicUrl><Url>url</Url></item></Articles><FuncFlag>1</FuncFlag></xml>`
)

func TestResponseNews(t *testing.T) {
    resp := &Response{
        msgBase: msgBase {
            ToUserName: "toUser",
            FromUserName: "fromUser",
            CreateTime: 12345678000000000,
            MsgType: News,
        },
        ArticleCount: 2,
        Articles: []*item{
            &item{
                Title: "title1",
                Description: "description1",
                PicUrl: "picurl",
                Url: "url",
            },
            &item{
                Title: "title",
                Description: "description",
                PicUrl: "picurl",
                Url: "url",
            },
        },
        FuncFlag: 1,
    }
    data, err := resp.Encode()
    if err != nil {
        t.Error(err)
        return
    }
    if resp.CreateTime != 12345678 * time.Second {
        t.Errorf("CreateTime shouldn't change: %d", resp.CreateTime)
        return
    }
    if string(data) != RespNewsCase {
        t.Errorf("%s", data)
        return
    }
}

func TestResponseText(t *testing.T) {
    resp := &Response{
        msgBase: msgBase {
            ToUserName: "toUser",
            FromUserName: "fromUser",
            CreateTime: 12345678000000000,
            MsgType: Text,
            Content: "content",
        },
    }
    data, err := resp.Encode()
    if err != nil {
        t.Error(err)
        return
    }
    if resp.CreateTime != 12345678 * time.Second {
        t.Errorf("CreateTime shouldn't change: %d", resp.CreateTime)
        return
    }
    if string(data) != RespTextCase {
        t.Errorf("%s", data)
        return
    }
}
