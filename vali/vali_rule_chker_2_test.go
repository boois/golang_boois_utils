package vali

import (
	"testing"
)

func TestValiRuleChker2(t *testing.T) {
	vali_rule_chker:=ValiRuleChker{}

	vali_rule_info:=ValiRuleInfo{}

	//test start
	vali_rule_info.Rule=`name -n -t str`
	vali_rule_info.ParseRule()
	vali_rule_chker.vali_rule_info=vali_rule_info
	is_ok,msg:=vali_rule_chker.Chk("")
	if is_ok!=OK || msg!="验证通过"{
		t.Errorf("验证失败%s:%s",is_ok,msg)
	}









}
