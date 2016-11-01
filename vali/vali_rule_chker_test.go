package vali

import (
	"testing"
)

func TestValiRuleChker(t *testing.T) {
	vali_rule_chker:=ValiRuleChker{}

	vali_rule_info:=ValiRuleInfo{}

	//test start
	vali_rule_info.Rule=`name -t str`
	vali_rule_info.ParseRule()
	vali_rule_chker.vali_rule_info=vali_rule_info
	is_ok,msg:=vali_rule_chker.Chk("aaa")
	if is_ok!=OK || msg!="验证通过"{
		t.Errorf("验证失败%s:%s",is_ok,msg)
	}
	//----

	vali_rule_info.Rule=`name -t int`
	vali_rule_info.ParseRule()
	vali_rule_chker.vali_rule_info=vali_rule_info
	is_ok,msg=vali_rule_chker.Chk("1")
	if is_ok!=OK || msg!="验证通过"{
		t.Errorf("验证失败%s:%s",is_ok,msg)
	}
	//t.Skip("skip")
	//
	vali_rule_info.Rule=`name -t int`
	vali_rule_info.ParseRule()
	vali_rule_chker.vali_rule_info=vali_rule_info
	is_ok,msg=vali_rule_chker.Chk("-1")
	if is_ok!=OK || msg!="验证通过"{
		t.Errorf("验证失败%s:%s",is_ok,msg)
	}
	//
	vali_rule_info.Rule=`name -t +int`
	vali_rule_info.ParseRule()
	vali_rule_chker.vali_rule_info=vali_rule_info
	is_ok,msg=vali_rule_chker.Chk("1")
	if is_ok!=OK || msg!="验证通过"{
		t.Errorf("验证失败%s:%s",is_ok,msg)
	}
	//
	vali_rule_info.Rule=`name -t -int`
	vali_rule_info.ParseRule()
	vali_rule_chker.vali_rule_info=vali_rule_info
	is_ok,msg=vali_rule_chker.Chk("-1")
	if is_ok!=OK || msg!="验证通过"{
		t.Errorf("验证失败%s:%s",is_ok,msg)
	}
	//
	vali_rule_info.Rule=`name -t zero`
	vali_rule_info.ParseRule()
	vali_rule_chker.vali_rule_info=vali_rule_info
	is_ok,msg=vali_rule_chker.Chk("0")
	if is_ok!=OK || msg!="验证通过"{
		t.Errorf("验证失败%s:%s",is_ok,msg)
	}
	//float
	vali_rule_info.Rule=`name -t float`
	vali_rule_info.ParseRule()
	vali_rule_chker.vali_rule_info=vali_rule_info
	is_ok,msg=vali_rule_chker.Chk("1.3")
	if is_ok!=OK || msg!="验证通过"{
		t.Errorf("验证失败%s:%s",is_ok,msg)
	}
	//
	vali_rule_info.Rule=`name -t float`
	vali_rule_info.ParseRule()
	vali_rule_chker.vali_rule_info=vali_rule_info
	is_ok,msg=vali_rule_chker.Chk("-1.3")
	if is_ok!=OK || msg!="验证通过"{
		t.Errorf("验证失败%s:%s",is_ok,msg)
	}
	//
	vali_rule_info.Rule=`name -t +flaot`
	vali_rule_info.ParseRule()
	vali_rule_chker.vali_rule_info=vali_rule_info
	is_ok,msg=vali_rule_chker.Chk("1.3")
	if is_ok!=OK || msg!="验证通过"{
		t.Errorf("验证失败%s:%s",is_ok,msg)
	}
	//
	vali_rule_info.Rule=`name -t -float`
	vali_rule_info.ParseRule()
	vali_rule_chker.vali_rule_info=vali_rule_info
	is_ok,msg=vali_rule_chker.Chk("-1.3")
	if is_ok!=OK || msg!="验证通过"{
		t.Errorf("验证失败%s:%s",is_ok,msg)
	}

	//guid
	vali_rule_info.Rule=`name -t guid`
	vali_rule_info.ParseRule()
	vali_rule_chker.vali_rule_info=vali_rule_info
	is_ok,msg=vali_rule_chker.Chk("7ea995ac499a495e8d21b3a32a5f0a5e")
	if is_ok!=OK || msg!="验证通过"{
		t.Errorf("验证失败%s:%s",is_ok,msg)
	}
	//
	vali_rule_info.Rule=`name -t guid_list`
	vali_rule_info.ParseRule()
	vali_rule_chker.vali_rule_info=vali_rule_info
	is_ok,msg=vali_rule_chker.Chk("7ea995ac499a495e8d21b3a32a5f0a5e,7ea995ac499a495e8d21b3a32a5f0a5e")
	if is_ok!=OK || msg!="验证通过"{
		t.Errorf("验证失败%s:%s",is_ok,msg)
	}
	//guid guid_list

	vali_rule_info.Rule=`name -t guid|guid_list`
	vali_rule_info.ParseRule()
	vali_rule_chker.vali_rule_info=vali_rule_info
	is_ok,msg=vali_rule_chker.Chk("7ea995ac499a495e8d21b3a32a5f0a5e")
	if is_ok!=OK || msg!="验证通过"{
		t.Errorf("验证失败%s:%s",is_ok,msg)
	}
	vali_rule_info.Rule=`name -t guid|guid_list`
	vali_rule_info.ParseRule()
	vali_rule_chker.vali_rule_info=vali_rule_info
	is_ok,msg=vali_rule_chker.Chk("7ea995ac499a495e8d21b3a32a5f0a5e,7ea995ac499a495e8d21b3a32a5f0a5e")
	if is_ok!=OK || msg!="验证通过"{
		t.Errorf("验证失败%s:%s",is_ok,msg)
	}
	//v
	vali_rule_info.Rule=`name -v a|b|C`
	vali_rule_info.ParseRule()
	vali_rule_chker.vali_rule_info=vali_rule_info
	is_ok,msg=vali_rule_chker.Chk("a")
	if is_ok!=OK || msg!="验证通过"{
		t.Errorf("验证失败%s:%s",is_ok,msg)
	}
	//len
	vali_rule_info.Rule=`name -l 1,3`
	vali_rule_info.ParseRule()
	vali_rule_chker.vali_rule_info=vali_rule_info
	is_ok,msg=vali_rule_chker.Chk("a")
	if is_ok!=OK || msg!="验证通过"{
		t.Errorf("验证失败%s:%s",is_ok,msg)
	}
	vali_rule_info.Rule=`name -l 1,3`
	vali_rule_info.ParseRule()
	vali_rule_chker.vali_rule_info=vali_rule_info
	is_ok,msg=vali_rule_chker.Chk("aa")
	if is_ok!=OK || msg!="验证通过"{
		t.Errorf("验证失败%s:%s",is_ok,msg)
	}
	vali_rule_info.Rule=`name -l 1,3`
	vali_rule_info.ParseRule()
	vali_rule_chker.vali_rule_info=vali_rule_info
	is_ok,msg=vali_rule_chker.Chk("aaa")
	if is_ok!=OK || msg!="验证通过"{
		t.Errorf("验证失败%s:%s",is_ok,msg)
	}
	vali_rule_info.Rule=`name -l 1,`
	vali_rule_info.ParseRule()
	vali_rule_chker.vali_rule_info=vali_rule_info
	is_ok,msg=vali_rule_chker.Chk("aaaa")
	if is_ok!=OK || msg!="验证通过"{
		t.Errorf("验证失败%s:%s",is_ok,msg)
	}

	//chk fail









}
