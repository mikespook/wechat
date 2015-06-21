package wechat

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	UrlMenuCreate = "https://api.weixin.qq.com/cgi-bin/menu/create"
	UrlMenuGet    = "https://api.weixin.qq.com/cgi-bin/menu/get"
	UrlMenuDelete = "https://api.weixin.qq.com/cgi-bin/menu/delete"
)

type Menu struct {
	Button []Button `json:"button"`
}

type Button struct {
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Key       string   `json:"key"`
	Url       string   `json:"url"`
	SubButton []Button `json:"sub_button"`
}

func NewMenu(accessToken string, menu *Menu) error {
	b, err := json.Marshal(menu)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(b)
	req, err := http.NewRequest("POST", UrlMenuCreate, buf)
	if err != nil {
		return err
	}
	values := url.Values{
		"access_token": []string{accessToken},
	}
	req.URL.RawQuery = values.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var data Error
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}
	if data.Code != 0 {
		return &data
	}
	return nil
}

func GetMenu(accessToken string) (*Menu, error) {
	req, err := http.NewRequest("GET", UrlMenuGet, nil)
	if err != nil {
		return nil, err
	}
	values := url.Values{
		"access_token": []string{accessToken},
	}
	req.URL.RawQuery = values.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var data struct {
		Menu
		Error
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	if data.Code != 0 {
		return nil, &data.Error
	}
	return &data.Menu, nil
}

func DeleteMenu(accessToken string) error {
	req, err := http.NewRequest("GET", UrlMenuDelete, nil)
	if err != nil {
		return err
	}
	values := url.Values{
		"access_token": []string{accessToken},
	}
	req.URL.RawQuery = values.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var data Error
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}
	if data.Code != 0 {
		return &data
	}
	return nil
}
