// 萧鸣 boois@qq.com
package boois_account

import (
	"testing"
	"fmt"
)

func Test_DecodeAccountStr(t *testing.T)  { // 测试用例必须以Test开头
	ad := BooisAccountAdapter{}
	ad.Load(`
	account1 123123 小明1 0 auth1:1,auth2:2,auth3:3 额外字段1 额外字段2 额外字段3 额外字段4
	account2 123123 小明2 0 auth1,auth2:,:3,,,,,auth9:9 额外字段1 额外字段2 额外字段3 额外字段4
	account3 123123 小明3 0 auth1:1,auth2:2,auth3:3 额外字段1 额外字段2 额外字段3 额外字段4
	account4 123123 小明4 0 auth1:1,auth2:2,auth3:3 额外字段1 额外字段2 额外字段3 额外字段4
	account5 123123 小明5 0 auth1:1,auth2:2,auth3:3 额外字段1 额外字段2 额外字段3 额外字段4
	account6 123123 小明6 0 auth1:1,auth2:2,auth3:3 额外字段1 额外字段2 额外字段3 额外字段4
	`)
	if len(ad.AccountMap) != 6 { t.Errorf("本应该为6个账号,当前为%d个",len(ad.AccountMap))}
	if !ad.HasAccount("account1") {
		t.Error("应该存在账号account1")
	}else{
		info := ad.GetAccountInfo("account1")
		if (info.Account != "account1") {t.Error("账号account1出错")}
		if (info.Psw != "123123") {t.Error("密码出错")}
		if (info.Nickname != "小明1") {t.Error("昵称出错")}
		if (len(info.AuthMap) != 3) {t.Error("AuthMap数量错误")}
		if (len(info.OtherDataArr) != 4) {t.Error("OtherDataArr数量错误")}
		if (info.GetOtherData(0) != "额外字段1") {{t.Error(fmt.Sprintf("GetOtherData(%d)错误",0))}}
		if (info.GetOtherData(1) != "额外字段2") {{t.Error(fmt.Sprintf("GetOtherData(%d)错误",1))}}
		if (info.GetOtherData(2) != "额外字段3") {{t.Error(fmt.Sprintf("GetOtherData(%d)错误",2))}}
		if (info.GetOtherData(3) != "额外字段4") {{t.Error(fmt.Sprintf("GetOtherData(%d)错误",3))}}

	}
	if !ad.HasAccount("account2") {t.Error("应该存在账号account2")}else{
		info := ad.GetAccountInfo("account2")
		if(len(info.AuthMap) !=1){t.Error("AuthMap数量错误")}
		if(info.GetAuth("auth9")!=9){t.Error("AuthMap获取auth9错误")}
		if(info.GetAuth("auth1")!=-1){t.Error("AuthMap获取auth错误")}
		if(info.GetOtherData(311)!=""){t.Error("GetOtherData错误")}

	}
	if !ad.HasAccount("account3") {t.Error("应该存在账号account3")}
	if !ad.HasAccount("account4") {t.Error("应该存在账号account4")}
	if !ad.HasAccount("account5") {t.Error("应该存在账号account5")}
	if !ad.HasAccount("account6") {t.Error("应该存在账号account6")}
	if ad.HasAccount("account7") { t.Error("account7")}

}
