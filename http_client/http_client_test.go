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
