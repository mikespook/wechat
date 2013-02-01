package handle

import (
    "time"
    "sort"
    "net/http"
    "io/ioutil"
    "crypto/sha1"
    "github.com/mikespook/wechat/msg"
    "github.com/mikespook/golib/log"
)

var (
    DefaultToken = "thisiswechathahahaha"
)

func Wechat(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()
    if r.Method == "post" {
        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
            log.Error(err)
            return
        }
        wreq := &msg.Request{}
        if err = msg.Unmarshal(body, wreq); err != nil {
            log.Error(err)
            return
        }
        wresp := wechattest(wreq)
        data, err := msg.Marshal(wresp)
        if _, err := w.Write(data); err != nil {
            log.Error(err)
        }
        return
    } else {
        if wechatSignature(DefaultToken, r.FormValue("timestamp"),
            r.FormValue("nonce")) == r.FormValue("signature") {
            w.Write([]byte(r.FormValue("echostr")))
        }
    }
}

func wechattest(r *msg.Request) (resp msg.Response) {
    resp.ToUserName = r.ToUserName
    resp.FromUserName = r.FromUserName
    resp.CreateTime = time.Duration(time.Now().Unix())
    resp.MsgType = msg.Text
    resp.Content = "Hello world"
    return
}

func wechatSignature(token, timestamp, nonce string) string {
    strs := sort.StringSlice{token, timestamp, nonce}
    sort.Strings(strs)
    str := ""
    for _, s := range strs {
        str += s
    }
    h := sha1.New()
    h.Write([]byte(str))
    return string(h.Sum(nil))
}
