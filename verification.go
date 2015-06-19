package wechat

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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

const UrlToken = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential"

func AccessToken(id, secret string) (string, int, error) {
	// https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET
	req, err := http.NewRequest("GET", UrlToken, nil)
	if err != nil {
		return "", 0, err
	}
	req.URL.Query().Add("appid", id)
	req.URL.Query().Add("secret", secret)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var data struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		ErrCode     int    `json:"errcode"`
		ErrMsg      string `json:"errmsg"`
	}
	if err := json.Unmarshal(body, data); err != nil {
		return "", 0, err
	}
	if data.ErrCode != 0 {
		return "", 0, &Error{data.ErrCode, data.ErrMsg}
	}
	return data.AccessToken, data.ExpiresIn, nil
}