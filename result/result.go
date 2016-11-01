package result

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