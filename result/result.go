package result

import "strconv"

type ReturnResult struct {
	Code   int `json:"code"`
	Msg    string `json:"msg"`
	Info   string `json:"info"`
	Debug  string `json:"debug"`
	Result interface{} `json:"result"`
}
var OK=ReturnResult{
	Code:0,
	Msg:"ok",
	Info:"操作成功",
	Debug:"",
	Result:nil,

}

func (this *ReturnResult) New(code int,msg string,info string,debug string,result interface{}) ReturnResult {
	this.Code=code
	this.Msg=msg
	this.Info=info
	this.Debug=debug
	this.Result=result
	return this
}
//实现了这个方法就可以作为函数执行错误的返回值
func (this *ReturnResult) Error() string {
	if this.Info!=""{
		return this.Info
	}
	return this.Msg+"("+strconv.Itoa(this.Code)+")"
}