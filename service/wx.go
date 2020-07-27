package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const APPID = "wx3a8da23d5c4ae903"
const APPSECRET = "315cee879d47dc24204f17cf80003364"

type BodyStruct struct {
	Openid      string
	Session_key string
	Unionid     string
	Errcode     string
	Errmsg      string
}

func GetOpenID(code string) (BodyStruct, error) {
	var bodyStruct BodyStruct
	Url, _ := url.Parse("https://api.weixin.qq.com/sns/jscode2session")
	params := url.Values{}
	params.Set("appid", APPID)
	params.Set("secret", APPSECRET)
	params.Set("js_code", code)
	params.Set("grant_type", "authorization_code")
	Url.RawQuery = params.Encode()
	res, err := http.Get(Url.String())
	if err != nil {
		return bodyStruct, err
	}

	body, err := ioutil.ReadAll(res.Body)

	json.Unmarshal(body, &bodyStruct)
	return bodyStruct, nil
}
