package wechat

import (
    "sort"
    "net/http"
    "io/ioutil"
    "crypto/sha1"
    "github.com/mikespook/golib/log"
)

var (
    Token = "thisiswechattoken"
)

type HandlerFunc func(*Request)(*Response, error)

func Handle(w http.ResponseWriter, r *http.Request, h HandlerFunc) {
    defer r.Body.Close()
    if r.Method == "post" {
        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
            log.Error(err)
            w.WriteHeader(500)
            return
        }
        var wreq *Request
        if wreq, err = DecodeRequest(body); err != nil {
            log.Error(err)
            w.WriteHeader(500)
            return
        }
        wresp, err := h(wreq)
        if err != nil {
            log.Error(err)
            w.WriteHeader(500)
            return
        }
        data, err := wresp.Encode()
        if _, err := w.Write(data); err != nil {
            log.Error(err)
            w.WriteHeader(500)
        }
        return
    } else {
        log.Debugf("%V", r)
        if Signature(Token, r.FormValue("timestamp"),
            r.FormValue("nonce")) == r.FormValue("signature") {
            w.Write([]byte(r.FormValue("echostr")))
        } else {
            w.WriteHeader(403)
        }
    }
}

/*
func wechattest(r *msg.Request) (resp *msg.Response, err error) {
    resp = NewResponse()
    resp.ToUserName = r.ToUserName
    resp.FromUserName = r.FromUserName
    resp.MsgType = msg.Text
    resp.Content = "Hello world"
    return
}
*/

func Signature(token, timestamp, nonce string) string {
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
