package vali

import (
	"strings"
	"strconv"
	"git.boois.cn/d01/git_repo/boois_utils.git/cmd_parser"
)

var INT = 1 << 0  // 整数
var P_INT = 1 << 1  // 正整数 positive integer
var N_INT = 1 << 2  // 负整数 negative integer
var FLOAT = 1 << 3  // 小数
var P_FLOAT = 1 << 4  // 正小数 可以无小数点,兼容正整数
var N_FLOAT = 1 << 5  // 负小数 可以无小数点,兼容负整数
var BOOL = 1 << 6  // 布尔值 true|false|0|1 大小写不敏感
var INT_LIST = 1 << 7  // 数字数组 1,2,3,4,5
var STR_LIST = 1 << 8  // 字符串数组 a,b,c,d,e
var DICT = 1 << 9  // 字典 "a":"foo",b:"bar","c":"baz","d":"qux",必须用双引号将kv包裹起来
var JSON = 1 << 10  // 能通过 json.loads(foo)的字符串
var GUID_LIST = 1 << 11  // 32位无横线GUID,36位有横线GUID,..
var FLOAT_LIST = 1 << 12  // 小数组合成的列表
var ZERO = 1 << 13  // 零
var STR = 1 << 14 // 任意字符
var GUID = 1 << 15 // 32位无横线GUID,36位有横线GUID

var TYPE_LIST=map[string]int{
	"INT": INT,
        "P_INT": P_INT,
        "N_INT": N_INT,
        "FLOAT": FLOAT,
        "P_FLOAT": P_FLOAT,
        "N_FLOAT": N_FLOAT,
        "BOOL": BOOL,
        "INT_LIST": INT_LIST,
        "STR_LIST": STR_LIST,
        "DICT": DICT,
        "JSON": JSON,
        "GUID_LIST": GUID_LIST,
        "FLOAT_LIST": FLOAT_LIST,
        "ZERO": ZERO,
        "STR": STR,
        "GUID": GUID,
}

var REGX_TMP = map[string]string{
        // "mobi": r"^0?(13|14|15|18|17)[0-9]{9}$",  // 手机号 +86 13599999999
        "mobi": "^0?1[0-9]{10}$",  // 手机号 +86 13599999999
        "email": `^\w[-\w.+]*@([A-Za-z0-9][-A-Za-z0-9]*\.)+[A-Za-z]{2,14}$`,  // 邮箱 boois@qq.com
        "zip_code": `^\d{6}$`,  // 邮政编码 350001
        "phone": `^\d+$|^\d{4}[\s-]\d+$|^\+{0,1}86[\s-]\d{4}[\s-]\d+$`,
        "account": "^[a-zA-Z][0-9a-zA-Z_]*$",  // 账号 字母数字下划线
        "passwd": "^[a-zA-Z][0-9a-zA-Z_]*$",  // 密码 字母数字下划线
        "nickname": "^[\u4e00-\u9fa5_0-9a-zA-Z-]*$",  // 昵称 中文英文数字下划线短连接线,不能有其他特殊字符
        "qq": "^[1-9][0-9]{4,}$",  // QQ号 [1-9][0-9]{4,}
        "id_card": `^(\d{15}$|^\d{18}$|^\d{17}(\d|X|x))$`,  // 身份证
        "ip": `^((?:(?:25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))\.){3}(?:25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d))))$`,  // ip地址
        "url": `^((https|http|ftp|rtsp|mms)?:\/\/)[^\s]*`,  // 网络链接
        "zh": "^[\u4e00-\u9fa5]*$",  // 仅为中文,这里有个要点,中文匹配必须正则规则也使用u编码,否则无法匹配
}

type ValiRuleInfo struct {
	Rule string
	Field string
	Info string
	Regx string
	Regx_Str string
	Cmd_List map[string]string
	Type_Str string
	Type_Arr []string
	Type int
	Len_Str string
	Len []int
	Required bool
	Warning string
	Values []string

	HasField bool
	HasLen bool
	HasRegx bool
	HasInfo bool
	HasWarning bool
	HasRequired bool
	HasType bool
	HasValues bool



}
func (this *ValiRuleInfo) getLenVal(val string) int {
	i,err:=strconv.Atoi(val)
	if err==nil{
		return i
	}

	val_trim:=strings.Trim(val," ")
	if strings.HasPrefix(val_trim,"-") {
		ii,err2:=strconv.Atoi(val_trim)
		if err2==nil{
			return -ii
		}
	}
	return -1

}


func (this *ValiRuleInfo) ParseRule() {
	this.Field,this.Cmd_List=cmd_parser.Cmd_parse(this.Rule)
	this.Info=this.Cmd_List["i"]
	this.Type_Str=strings.Replace(strings.Replace(this.Cmd_List["t"],"+","P_",-1),"-", "N_",-1)
	this.Type_Arr=strings.Split(this.Type_Str,"|")
	this.Type=0

	for _,_type:=range this.Type_Arr{
		_type_upper:=strings.ToUpper(_type)
		val,exists:=TYPE_LIST[_type_upper]
		if exists{
			this.Type=this.Type | val
		}
	}
	this.Regx_Str=this.Cmd_List["r"]
	if strings.HasPrefix(this.Regx_Str,"/"){
		this.Regx=this.Regx_Str
	}else{
		regx_tmp,exists2:=REGX_TMP[strings.ToLower(this.Regx_Str)]
		if regx_tmp!="" && exists2{
			this.Regx=regx_tmp
		}else{
			this.Regx=""
		}
	}
	this.Len_Str=this.Cmd_List["l"]
	len_arr:=strings.Split(this.Len_Str,",")

	if len(len_arr)==1{
		this.Len=[]int {this.getLenVal(len_arr[0]),-1}
	}else if len(len_arr)>=2{
		this.Len=[]int {this.getLenVal(len_arr[0]),this.getLenVal(len_arr[1])}
	}

	this.Required=this.Cmd_List["n"]!=""
	this.Warning=this.Cmd_List["w"]
	values:=this.Cmd_List["v"]
	if values==""{
		this.Values=[]string{}
	}else{
		this.Values=strings.Split(values,"|")
	}

	//has
	this.HasField=this.Field!=""
	this.HasLen=this.Cmd_List["l"]!=""
	this.HasRegx=this.Cmd_List["r"]!=""
	this.HasInfo=this.Cmd_List["i"]!=""
	this.HasWarning=this.Cmd_List["w"]!=""
	this.HasRequired=this.Cmd_List["n"]!=""
	this.HasType=this.Cmd_List["t"]!=""
	this.HasValues=this.Cmd_List["v"]!=""



}
