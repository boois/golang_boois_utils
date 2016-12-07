// 萧鸣 boois@qq.com
package boois_account

import (
	"strings"
	"strconv"
)

// # 账号 密码 昵称 0 a:1,b:1,c:1,d:0,e:1 其他1 其他2 其他3 其他4 其他5
type BooisAccountInfo struct {
	Account string
	Psw string
	Nickname string
	IsLock bool
	AuthStr string
	AuthMap map[string]int
	OtherDataStr string
	OtherDataArr []string
}
func (this *BooisAccountInfo) GetAuth(key string) int {
	if this.AuthMap == nil {return -1}
	if v,ok := this.AuthMap[key];ok {
		return v
	}
	return -1
}
// 根据位置获取一个额外数据
func (this *BooisAccountInfo) GetOtherData(pos int) string {
	if this.OtherDataArr == nil {return ""}
	if len(this.OtherDataArr) > pos {
		return this.OtherDataArr[pos]
	}else{
		return ""
	}
}
func trim(val string) string {
	return strings.Trim(strings.Trim(val,"\t")," ")
}
func (this *BooisAccountInfo) Load(raw string) {
	raw = trim(raw)
	if raw == "" {return}
	if raw[0:1] == "#" {return}// 如果以# 开头则不处理
	arr := strings.Split(raw," ")
	if len(arr) == 0 {return}
	if len(arr) > 0 {this.Account = trim(arr[0])}
	if len(arr) > 1 {this.Psw = trim(arr[1])}
	if len(arr) > 2 {this.Nickname = trim(arr[2])}
	if len(arr) > 3 {
		i,_ := strconv.Atoi(trim(arr[3]))
		this.IsLock = i!=0
	}
	if len(arr) > 4 {
		this.AuthMap = map[string]int {}
		auth_str := arr[4]
		auth_arr := strings.Split(auth_str,",")
		for x := range auth_arr {
			auth_ele_str := auth_arr[x]
			auth_ele_arr := strings.Split(auth_ele_str,":")
			if len(auth_ele_arr) <2 {continue}
			if auth_ele_arr[0] == "" {continue}
			if auth_ele_arr[1] == "" {continue}
			i,_ := strconv.Atoi(trim(auth_ele_arr[1]))
			this.AuthMap[auth_ele_arr[0]]=i
		}
	}
	if len(arr)>5 {
		this.OtherDataArr = []string{}
		for ii:=5;ii<len(arr);ii++{
			this.OtherDataArr = append(this.OtherDataArr,arr[ii])
		}
	}

}

type BooisAccountAdapter struct {
	AccountMap map[string]BooisAccountInfo
}

func (this *BooisAccountAdapter) Load(content string){
	this.AccountMap = map[string]BooisAccountInfo{}
	con_arr := strings.Split(content,"\n")
	//println(len(con_arr))
	for x := range con_arr {
		line := trim(con_arr[x])
		if line == "" {continue}
		info := BooisAccountInfo{}
		info.Load(line)
		if info.Account == "" {continue}
		this.AccountMap[info.Account] = info
	}
}

func (this *BooisAccountAdapter) HasAccount(account string) bool{
	if this.AccountMap == nil {return false}
	if _,ok := this.AccountMap[account];ok {
		return true
	}else{
		return false
	}
}

func (this *BooisAccountAdapter) GetAccountInfo(account string) BooisAccountInfo{
	if this.AccountMap == nil {return BooisAccountInfo{}}
	if v,ok := this.AccountMap[account];ok {
		return v
	}else{
		return BooisAccountInfo{}
	}
}