package boois_conf_parser

import (
	"testing"
	"fmt"
	"encoding/json"
)

func TestParse(t *testing.T) {
	txt:=`
host=127.0.0.1
	[main]\r\n
	[main]
port=6666
[mysql]
mysql=root:root@tcp(127.0.0.1:3306)/test?charset=utf8&loc=Local
[oss]
app_key = A3kzBRzEvHNpj0F
app_sec = ShJJLCrQPgxt1EDJn8GFpyLYPEBPMw
end_point = oss-cn-hangzhou.aliyuncs.com
bucket = jiupin-private
prefix = jiumayun
[test]
host=0.0.0.0
	`
	res:=Parse(txt)
	json_str,err:=json.Marshal(res)
	fmt.Println(err)
	fmt.Println(string(json_str))
}
