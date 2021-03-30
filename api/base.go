package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Api interface {
	Translation(text string) string
}

type base struct {
	Method string
	Name   string
	Url    string
}

func (this *base) translation(url string, data map[string]interface{}) map[string]interface{} {
	if this.Method == "GET" {

	} else {
		b, _ := json.Marshal(data)
		resp, err := http.DefaultClient.Post(url, "application/x-www-form-urlencoded", bytes.NewReader(b))
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
		}
		res := map[string]interface{}{}
		err = json.Unmarshal(body, res)
		if err != nil {
			// handle error
		}
		return res
	}
	return nil
}
