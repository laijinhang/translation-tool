package api

import (
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tmt "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tmt/v20180321"
)

// api接入文档：https://cloud.tencent.com/document/product/551/15619
const (
	TMTApiUrl = "tmt.tencentcloudapi.com"
	TMTAppKey = ""
	TMTSecKey = ""
)

func GetTMTApi() *TMTApi {
	return &tmtApi
}

var tmtApi = TMTApi{
	AppKey: TMTAppKey,
	base: base{
		Method: "POST",
		Name:   "腾讯翻译君",
		Url:    TMTApiUrl,
	}}

type TMTApi struct {
	base
	AppKey string
}

func (this *TMTApi) GetName() string {
	return this.Name
}

func (this *TMTApi) GetAppKey() string {
	return this.AppKey
}

func (this *TMTApi) Translation(text string) string {
	credential := common.NewCredential(
		TMTAppKey,
		TMTSecKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = TMTApiUrl
	client, _ := tmt.NewClient(credential, "ap-guangzhou", cpf)

	request := tmt.NewTextTranslateRequest()
	request.SourceText = common.StringPtr(text)
	request.ProjectId = common.Int64Ptr(0)
	request.Source = common.StringPtr("en")
	request.Target = common.StringPtr("zh")
	response, err := client.TextTranslate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s\n", err)
		return ""
	}
	if err != nil {
		panic(err)
	}

	return *response.Response.TargetText
}
