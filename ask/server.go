package main

import (
    "os"
    "flag"
    "time"
    "net/http"
    "github.com/mikespook/golib/log"
    "github.com/mikespook/golib/signal"
    "github.com/mikespook/wechat/handle"
)

var (
    token = flag.String("token", handle.DefaultToken, "token for wechat")
    addr = flag.String("addr", ":8080", "the address to bind to")
)

func init() {
    if !flag.Parsed() {
        flag.Parse()
    }
    if *token != "" {
        handle.DefaultToken = *token
    }
}

func main() {
    defer func() {
        log.Message("Exit.")
        time.Sleep(time.Microsecond * 100)
    }()
    log.Message("Starting...")

    http.HandleFunc("/wechat", handle.Wechat)
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
