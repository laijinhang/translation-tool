package api

import (
	"encoding/json"
	"fmt"
	"net/url"
)

const (
	GoogleApiUrl = "https://translate.google.cn/translate_a/single?client=at&sl=en&tl=zh-CN&dt=t&q="
)

func GetGoogleApi() *GoogleApi {
	return &googleApi
}

var googleApi = GoogleApi{
	base: base{
		Method: "GET",
		Name:   "谷歌翻译",
		Url:    GoogleApiUrl,
	}}

type GoogleApi struct {
	base
	AppKey string
}

func (this *GoogleApi) GetName() string {
	return this.Name
}

func (this *GoogleApi) GetAppKey() string {
	return " "
}

func (this *GoogleApi) Translation(text string) string {
	body, err := this.Get(this.Url + url.QueryEscape(text))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	var v [][][]interface{}
	json.Unmarshal(body, &v)
	return v[0][0][0].(string)
}
