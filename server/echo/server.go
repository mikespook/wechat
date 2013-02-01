package main

import (
    "os"
    "flag"
    "time"
    "net/http"
    "github.com/mikespook/golib/log"
    "github.com/mikespook/golib/signal"
    "github.com/mikespook/wechat"
)

var (
    token = flag.String("token", wechat.Token, "token for wechat")
    addr = flag.String("addr", ":8080", "the address to bind to")
)

func init() {
    if !flag.Parsed() {
        flag.Parse()
    }
    if *token != "" {
        wechat.Token = *token
    }
}

func main() {
    defer func() {
        log.Message("Exit.")
        time.Sleep(time.Microsecond * 100)
    }()
    log.Message("Starting...")

    http.HandleFunc("/wechat", echo)
    go func() {
        log.Messagef("Bind to [%s]", *addr)
        if err := http.ListenAndServe(*addr, nil); err != nil {
            log.Error(err)
            signal.Send(os.Getpid(), os.Interrupt)
        }
    }()

    // signal handler
    sh := signal.NewHandler()
    sh.Bind(os.Interrupt, func() bool {return true})
    sh.Loop()
}

func echo(w http.ResponseWriter, r *http.Request) {
    wechat.Handle(w, r, func(r *wechat.Request) (resp *wechat.Response, err error){
        log.Messagef("ACCESS: %V", r)
        resp = wechat.NewResponse()
        resp.ToUserName = r.FromUserName
        resp.FromUserName = r.ToUserName
        resp.MsgType = wechat.Text
        resp.Content = r.Content
        return
    })
}
