package result

import "strconv"

type ReturnResult struct {
	Code   int `json:"code"`
	Msg    string `json:"msg"`
	Info   string `json:"info"`
	Debug  string `json:"debug"`
	Result interface{} `json:"result"`
}

var OK = ReturnResult{
	Code:0,
	Msg:"ok",
	Info:"操作成功",
	Debug:"",
	Result:nil,

}

func New(code int, msg string, info string, debug string, result interface{}) ReturnResult {
	return ReturnResult{
		Code:code,
		Msg:msg,
		Info:info,
		Debug:debug,
		Result:result,
	}
}
//实现了这个方法就可以作为函数执行错误的返回值
func (this *ReturnResult) Error() string {
	if this.Info != "" {
		return this.Info
	}
	return this.Msg + "(" + strconv.Itoa(this.Code) + ")"
}