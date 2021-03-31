package api

var apis []Api
var names []string

func init() {
	register(GetYouDaoApi()) // 有道翻译
}

func GetApis() []Api {
	return apis
}

func register(api Api) {
	if api.GetAppKey() == "" {
		return
	}
	apis = append(apis, api)
	names = append(names, api.GetName())
}
