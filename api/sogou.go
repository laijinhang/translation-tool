package api

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// api接入文档：https://deepi.sogou.com/doccenter/texttranslatedoc
const (
	SoGouApiUrl = "http://fanyi.sogou.com/reventondc/api/sogouTranslate"
	SoGouAppKey = ""
	SoGouSecKey = ""
)

func GetSoGouApi() *SoGouApi {
	return &soGouApi
}

var soGouApi = SoGouApi{
	AppKey: SoGouAppKey,
	base: base{
		Method: "POST",
		Name:   "搜狗翻译",
		Url:    SoGouApiUrl,
	}}

type SoGouApi struct {
	base
	AppKey string
}

func (this *SoGouApi) GetName() string {
	return this.Name
}

func (this *SoGouApi) GetAppKey() string {
	return this.AppKey
}

func (this *SoGouApi) Translation(text string) string {
	salt := uuid.NewV1().String()
	sign := md5String(SoGouAppKey + text + salt + SoGouSecKey)
	data := map[string]string{
		"from": "en",
		"to":   "zh-CHS",
		"pid":  SoGouAppKey,
		"q":    text,
		"salt": salt,
		"sign": sign,
	}
	var resp SoGouResp
	err := this.base.Post(data, &resp)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if resp.ErrorCode != "" && resp.ErrorCode != "0" {
		fmt.Println("err:", resp.ErrorCode)
		return ""
	}
	return resp.Translation
}

type SoGouResp struct {
	ErrorCode   string `json:"error_code"`
	Query       string `json:"query"`
	Translation string `json:"translation"`
}
