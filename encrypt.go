package wechat

import (
	"encoding/xml"
	"io"
)

type EncryptMessage struct {
	ToUserName     string
	Encrypt        string
	token, key, id string
}

func (data *EncryptMessage) Decode(sign, ts, nonce string) io.Reader {
	sign0 := innerSignature(data.token, ts, nonce, data.Encrypt)
	if sign0 != sign {
		return nil
	}
	return nil
}

func NewEncryptMessage(r io.Reader, token, key, id string) (*EncryptMessage, error) {
	d := xml.NewDecoder(r)
	var data EncryptMessage
	if err := d.Decode(&data); err != nil {
		return nil, err
	}
	data.token = token
	data.key = key
	data.id = id
	return &data, nil
}
