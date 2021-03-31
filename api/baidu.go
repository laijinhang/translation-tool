package api

import (
	"crypto/md5"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// api接入文档：http://api.fanyi.baidu.com/doc/21
const (
	BaiDuApiUrl = "https://fanyi-api.baidu.com/api/trans/vip/translate"
	BaiDuAppKey = ""
	BaiDuSecKey = ""
)

func GetBaiDuApi() *BaiDuApi {
	return &baiDuApi
}

var baiDuApi = BaiDuApi{
	AppKey: BaiDuAppKey,
	base: base{
		Method: "POST",
		Name:   "百度翻译",
		Url:    BaiDuApiUrl,
	}}

type BaiDuApi struct {
	base
	AppKey string
}

func (this *BaiDuApi) GetName() string {
	return this.Name
}

func (this *BaiDuApi) GetAppKey() string {
	return this.AppKey
}

func (this *BaiDuApi) Translation(text string) string {
	salt := uuid.NewV1().String()
	sign := md5String(BaiDuAppKey + text + salt + BaiDuSecKey)
	data := map[string]string{
		"from":  "en",
		"to":    "zh",
		"appid": BaiDuAppKey,
		"q":     text,
		"salt":  salt,
		"sign":  sign,
	}
	var resp BaiDuResp
	err := this.base.translation(data, &resp)
	if err != nil {
		fmt.Println(err)
	}
	if len(resp.TransResult) == 0 {
		return ""
	}
	return resp.TransResult[0].Dst
}

func md5String(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return fmt.Sprintf("%x", m.Sum(nil))
}

type BaiDuResp struct {
	From        string `json:"from"`
	To          string `json:"to"`
	TransResult []struct {
		Src string `json:"src"`
		Dst string `json:"dst"`
	} `json:"trans_result"`
}
