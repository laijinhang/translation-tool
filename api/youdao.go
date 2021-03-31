package api

import (
	"crypto/sha256"
	"fmt"
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

var youDaoApi = YouDaoApi{
	AppKey: YouDaoAppKey,
	base: base{
		Method: "POST",
		Name:   "有道翻译",
		Url:    YouDaoApiUrl,
	}}

type YouDaoApi struct {
	base
	AppKey string
}

func (this *YouDaoApi) GetName() string {
	return this.Name
}

func (this *YouDaoApi) GetAppKey() string {
	return this.AppKey
}

func (this *YouDaoApi) Translation(text string) string {
	curtime := strconv.FormatInt(time.Now().Unix(), 10)
	salt := uuid.NewV1().String()
	sign := sha256String(YouDaoAppKey + truncate(text) + salt + curtime + YouDaoSecKey)
	data := map[string]string{
		"from":     "en",
		"to":       "zh-CHS",
		"signType": "v3",
		"curtime":  curtime,
		"appKey":   YouDaoAppKey,
		"q":        text,
		"salt":     salt,
		"sign":     sign,
		//"voice":    "",
		//"strict":   "",
		//"vocabId":  "",
	}
	//data['vocabId'] = "您的用户词表ID"
	var resp YouDaoResp
	err := this.base.translation(data, &resp)
	if err != nil {
		fmt.Println(err)
	}
	if resp.ErrorCode != "0" {
		fmt.Println(resp)
		return ""
	}
	return resp.Translation[0]
}

func truncate(q string) string {
	if len(q) <= 20 {
		return q
	}
	return fmt.Sprintf("%s%d%s", q[:10], len(q), q[len(q)-10:])
}

func sha256String(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}

type YouDaoReq struct {
}

type YouDaoResp struct {
	ErrorCode    string
	Query        interface{}
	Translation  []string
	Basic        interface{}
	Web          interface{}
	L            interface{}
	Dict         interface{}
	Webdict      interface{}
	TSpeakUrl    interface{}
	SpeakUrl     interface{}
	ReturnPhrase interface{}
}
