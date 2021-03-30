package api

import (
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
)

// api接入文档：https://ai.youdao.com/DOCSIRMA/html/%E8%87%AA%E7%84%B6%E8%AF%AD%E8%A8%80%E7%BF%BB%E8%AF%91/API%E6%96%87%E6%A1%A3/%E6%96%87%E6%9C%AC%E7%BF%BB%E8%AF%91%E6%9C%8D%E5%8A%A1/%E6%96%87%E6%9C%AC%E7%BF%BB%E8%AF%91%E6%9C%8D%E5%8A%A1-API%E6%96%87%E6%A1%A3.html
const (
	YouDaoApiUrl = "https://openapi.youdao.com/api"
	YouDaoAppKey = ""
	YouDaoSecKey = ""
)

func GetYouDaoApi() *YouDaoApi {
	return &youDaoApi
}

var youDaoApi = YouDaoApi{base{
	Method: "POST",
	Name:   "",
	Url:    "",
}}

type YouDaoApi struct {
	base
}

func (this *YouDaoApi) Translation(text string) string {
	curtime := strconv.FormatInt(time.Now().Unix(), 10)
	salt := uuid.NewV1().String()
	sign := YouDaoAppKey + text + salt + curtime + YouDaoSecKey
	data := map[string]interface{}{
		"from":     "en",
		"to":       "zh-CHS",
		"signType": "v3",
		"curtime":  curtime,
		"appKey":   YouDaoAppKey,
		"q":        text,
		"salt":     salt,
		"sign":     sign,
	}
	//data['vocabId'] = "您的用户词表ID"
	this.base.translation(YouDaoApiUrl, data)

	return ""
}
