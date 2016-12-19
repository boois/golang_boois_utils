package license

import (
	"strings"
	"net/http"

)

//软件启动时tarce
func StartTrace(data string)error {
	u := "http://lic.boois.cn/"
	client := &http.Client{}
	req, err := http.NewRequest("POST", u,strings.NewReader("data="+data))
	if err!=nil{
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("version", "1")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	return err
}
