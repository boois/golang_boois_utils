package http_client

import (
	"net/http"
	"strings"
	"io/ioutil"
	"net/url"
)

func Post(u string, data map[string]string) ([]byte,error){
	vals := url.Values{}
	for k, v := range data {
		vals.Set(k, v)
	}
	resp, err := http.PostForm(u,vals)
	if err != nil {
		return []byte{},err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{},err
	}
	return body,nil
}
func Get(url string) ([]byte,error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{},err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{},err
	}
	return body,nil
}
func PostJson(url string,json_str string) ([]byte,error) {
	resp, err := http.Post(url,"application/json",strings.NewReader(json_str))
	if err != nil {
		return []byte{},err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{},err
	}
	return body,nil
}