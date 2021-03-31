package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Api interface {
	GetName() string
	GetAppKey() string
	Translation(text string) string
}

type base struct {
	Method string
	Name   string
	Url    string
}

func (this *base) translation(data map[string]string, res interface{}) error {
	if this.Method == "GET" {
		//apiUrl := this.Url + "?"
		//query := ""
		//for k, v := range data {
		//	query += fmt.Sprintf("%s=%s&", k, url.QueryEscape(fmt.Sprint(v)))
		//}
		//apiUrl += query
		//fmt.Println(apiUrl)
		//return nil
		//resp, err := http.Get(apiUrl)
		//body, err := ioutil.ReadAll(resp.Body)
		//if err != nil {
		//	// handle error
		//}
		//res := map[string]interface{}{}
		//err = json.Unmarshal(body, &res)
		//fmt.Println(res)
		//fmt.Println(string(body))
		//if err != nil {
		//	// handle error
		//}
		//return res
	} else {
		v := url.Values{}
		for key, val := range data {
			v.Set(key, val)
		}
		body := ioutil.NopCloser(strings.NewReader(v.Encode()))
		client := &http.Client{}

		request, err := http.NewRequest("POST", this.Url, body)
		if err != nil {
			return err
		}
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded;param=value")
		resp, err := client.Do(request)
		defer resp.Body.Close()
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return json.Unmarshal(content, &res)
	}
	return nil
}
