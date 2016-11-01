package boois_utils

import (
	"testing"
	"fmt"
)

func TestValiRuleInfo_ParseRule(t *testing.T) {
	rule:=` field -i test%s123123 -n -l 1,4 -r "/mo\" -bi/gi" -t "-int|+float" -v aa|bbb`
	vali_rule_info:=ValiRuleInfo{}
	vali_rule_info.Rule=rule
	vali_rule_info.ParseRule()

	fmt.Println(vali_rule_info)
	/*
	Rule
	Field
	Info
	Regx
	Regx_Str
	Cmd_List map[string]string
	Type_Str
	Type_Arr []string
	Type
	Len_Str
	Len []int
	Required 
	Warning
	Values []string

	HasField 
	HasLen 
	HasRegx 
	HasInfo 
	HasWarning 
	HasRequired 
	HasType 
	HasValues */
	fmt.Println("Cmd_List: ",vali_rule_info.Cmd_List)
	fmt.Println("Rule: ",vali_rule_info.Rule)
	fmt.Println("Field: ",vali_rule_info.Field)
	fmt.Println("Info: ",vali_rule_info.Info)
	fmt.Println("Regx: ",vali_rule_info.Regx)
	fmt.Println("Regx_Str: ",vali_rule_info.Regx_Str)
	fmt.Println("Cmd_List: ",vali_rule_info.Cmd_List)
	fmt.Println("Type_Str: ",vali_rule_info.Type_Str)
	fmt.Println("Type_Arr: ",vali_rule_info.Type_Arr)
	fmt.Println("Type: ",vali_rule_info.Type)
	fmt.Println("Len_Str: ",vali_rule_info.Len_Str)
	fmt.Println("Len: ",vali_rule_info.Len)
	fmt.Println("Required: ",vali_rule_info.Required)
	fmt.Println("Warning: ",vali_rule_info.Warning)
	fmt.Println("Values: ",vali_rule_info.Values)
	fmt.Println("HasField: ",vali_rule_info.HasField)
	fmt.Println("HasLen: ",vali_rule_info.HasLen)
	fmt.Println("HasRegx: ",vali_rule_info.HasRegx)
	fmt.Println("HasInfo: ",vali_rule_info.HasInfo)
	fmt.Println("HasWarning: ",vali_rule_info.HasWarning)
	fmt.Println("HasRequired: ",vali_rule_info.HasRequired)
	fmt.Println("HasType: ",vali_rule_info.HasType)
	fmt.Println("HasValues: ",vali_rule_info.HasValues)
}
