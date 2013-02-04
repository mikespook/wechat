package wechat

import (
    "testing"
)

const (
    SING = "f7c3bc1d808e04732adf679965ccc34ca7ae3441"
)

func TestSignature(t *testing.T) {
    s := Signature("123", "456", "789")
    if s !=  SING {
        t.Errorf("Singature error: %s", s)
        return
    }
}
