package wechat

import "fmt"

type Error struct {
	Code int    `json:"errcode"`
	Msg  string `json:"errmsg"`
}

func (err *Error) Error() string {
	return fmt.Sprintf("%s (%d)", err.Msg, err.Code)
}
