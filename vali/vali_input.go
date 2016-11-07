package vali

import (
	"net/http"
	"fmt"
	"git.boois.cn/d01/git_repo/boois_utils.git/result"
)
//按照规则验证post的值  返回result
func Validete(r *http.Request,rule_strs ... string)result.ReturnResult {
	for _,rule_str :=range rule_strs{
		vali_rule_chker := ValiRuleChker{}
		vali_rule_info := ValiRuleInfo{}
		vali_rule_info.Rule = rule_str
		vali_rule_info.ParseRule()
		vali_rule_chker.vali_rule_info = vali_rule_info
		fmt.Println("field:",vali_rule_info.Field)
		fmt.Println("field val:",r.Form.Get(vali_rule_info.Field))
		is_ok, msg:=vali_rule_chker.Chk(r.Form.Get(vali_rule_info.Field))
		if !vali_rule_chker.IsValidated{
			return result.ReturnResult{
				Code   :is_ok,
				Msg    :"input_val_err",
				Info   :msg,
				Debug  :"",
				Result :nil,
			}
		}
	}
	//成功
	return result.ReturnResult{
		Code   :0,
		Msg    :"ok",
		Info   :"操作成功",
		Debug  :"",
		Result :nil,
	}
}
func ValideteVal(val_and_rules ... []string) result.ReturnResult{
	for _,val_and_rule :=range val_and_rules{
		vali_rule_chker := ValiRuleChker{}
		vali_rule_info := ValiRuleInfo{}
		vali_rule_info.Rule = val_and_rule[1]
		vali_rule_info.ParseRule()
		vali_rule_chker.vali_rule_info = vali_rule_info
		fmt.Println("field:",vali_rule_info.Field)
		fmt.Println("field val:",val_and_rule[0])
		code, msg:=vali_rule_chker.Chk(val_and_rule[0])
		if !vali_rule_chker.IsValidated{
			return result.ReturnResult{
				Code   :code,
				Msg    :"input_val_err",
				Info   :msg,
				Debug  :"",
				Result :nil,
			}
		}
	}
	//成功
	return result.ReturnResult{
		Code   :0,
		Msg    :"ok",
		Info   :"操作成功",
		Debug  :"",
		Result :nil,
	}

}

