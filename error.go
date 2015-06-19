package wechat

import "fmt"

type Error struct {
	Code int
	Msg  string
}

func (err *Error) Error() string {
	return fmt.Sprintf("%s (%d)", err.Msg, err.Code)
}
