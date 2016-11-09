package http_client

import (
	"testing"
	"fmt"
)

func TestPost(t *testing.T) {
	//http -f http://v0.client.sub-sys.ework.net/api/user/require/get \
//app_key=3b43209b3ca04ea1b7b14db28a209bc8 \
//user_id=33bf6059cba54c929f22a825b0091b19 \
//user_token=ef7e4713eb18417fa048480e7fde4044 \
//require_id=2654733dadc3483cb6dc37266ef8285f
	url:="http://v0.client.sub-sys.ework.net/api/user/require/get"
	data:=map[string]string{
		"app_key":"3b43209b3ca04ea1b7b14db28a209bc8",
		"user_id":"33bf6059cba54c929f22a825b0091b19",
		"user_token":"ef7e4713eb18417fa048480e7fde4044",
		"require_id":"2654733dadc3483cb6dc37266ef8285f",
	}
	body,err:=Post(url,data)
	fmt.Println(string(body))
	fmt.Println(err)
}
func TestGet(t *testing.T) {
	url:="http://counter.api.jiumayun.9pin.cn/api/app/counter/get_and_plus?app_key=279a38d18b4e40fc9054c7a64eed3af5&counter_name=test1111&delta=1&callback=cb"
	body,err:=Get(url)
	fmt.Println(string(body))
	fmt.Println(err)
}
