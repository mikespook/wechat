package wechat

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
)

func Signature(token, timestamp, nonce string) string {
	strs := sort.StringSlice{token, timestamp, nonce}
	sort.Strings(strs)
	h := sha1.New()
	for _, s := range strs {
		io.WriteString(h, s)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
