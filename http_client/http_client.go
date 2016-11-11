package http_client

import (
	"net/http"
	"strings"
	"io/ioutil"
)

func Post(url string, data map[string]string) ([]byte,error){
	data_str_arr := []string{}
	for k, v := range data {
		data_str_arr = append(data_str_arr, k + "=" + v)
	}
	data_str := strings.Join(data_str_arr, "&")
	resp, err := http.Post(url,"application/x-www-form-urlencoded",strings.NewReader(data_str))
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