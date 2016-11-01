package boois_utils

import (
	"strconv"
	"strings"
	"fmt"
	"regexp"
	"errors"
)

const REQUIRED_ERR = 1 //必选项错误
const LEN_ERR = 2 //长度错误
const TYPE_ERR = 4 //类型错误
const REGX_ERR = 8 //正则错误
const VAL_ERR = 16 //可选值错误
const OK = 32 //成功

func p_int_chker(val string) bool {
	val_trim := strings.Trim(val, " ")
	if zero_chker(val) {
		return false
	}
	i, err := strconv.Atoi(val_trim)
	if err == nil {
		if i <= 0 {
			return false
		}
	} else {
		return false
	}
	return true

}
func n_int_chker(val string) bool {
	val_trim := strings.Trim(val, " ")

	if zero_chker(val) {
		return false
	}
	i, err := strconv.Atoi(val_trim)
	if err == nil {
		if i >= 0 {
			return false
		}
	} else {
		if !strings.HasPrefix(val_trim, "-") {
			return false
		}
		arr := []rune(val_trim)
		sub_str := string(arr[1:])
		_, err2 := strconv.Atoi(sub_str)
		if err2 != nil {
			return false
		}
	}
	return true

}
func int_chker(val string) bool {
	val_trim := strings.Trim(val, " ")

	if zero_chker(val_trim) {
		return false
	}
	_, err := strconv.Atoi(val_trim)
	if err != nil {
		return false
	}
	return true

}
func p_float_chker(val string) bool {
	val_trim := strings.Trim(val, " ")
	if zero_chker(val_trim) {
		return false
	}
	//如果是整数 返回false
	_,err1:=strconv.Atoi(val)
	if err1==nil{
		return false
	}
	//是否为小数
	f, err := strconv.ParseFloat(val_trim, 64)
	if err == nil {
		return f > 0
	}
	return false

}
func n_float_chker(val string) bool {
	val_trim := strings.Trim(val, " ")
	if zero_chker(val_trim) {
		return false
	}
	//如果是整数 返回false
	_,err1:=strconv.Atoi(val)
	if err1==nil{
		return false
	}
	f, err := strconv.ParseFloat(val_trim, 64)
	if err == nil {
		return f < 0
	}
	return false

}
func float_chker(val string) bool {
	val_trim := strings.Trim(val, " ")
	if zero_chker(val_trim) {
		return true
	}
	return p_float_chker(val_trim) || n_float_chker(val_trim)

}
func bool_chker(val string) bool {
	val_trim := strings.Trim(val, " ")
	vals := []string{
		"true", "false", "1", "0",
	}
	for _, v := range vals {
		if strings.ToLower(val_trim) == v {
			return true
		}
	}
	return false

}
func int_list_chker(val string) bool {
	val_trim := strings.Trim(val, " ")
	if val_trim == "" {
		return false
	}
	if !strings.Contains(val_trim, ",") {
		return false
	}
	arr := strings.Split(val_trim, ",")
	for _, v := range arr {
		if strings.Replace(v, " ", "", -1) == "" {
			return false
		}
		if !int_chker(v) {
			return false
		}
	}
	return true

}
func str_list_chker(val string) bool {
	val_trim := strings.Trim(val, " ")
	if val_trim == "" {
		return false
	}
	if !strings.Contains(val_trim, ",") {
		return false
	}
	arr := strings.Split(val_trim, ",")
	for _, v := range arr {
		if strings.Replace(v, " ", "", -1) == "" {
			return false
		}
		if !str_chker(v) {
			return false
		}
	}
	return true

}
func guid_chker(val string) bool {
	val_trim := strings.Trim(val, " ")
	if val_trim == "" {
		return false
	}
	str_len := len(val_trim)

	if str_len != 32  && str_len != 36 {
		return false
	}
	if str_len == 36 {
		group := strings.Split(val_trim, "-")
		if len(group) != 5 {
			return false
		} else {
			if len(group[0]) != 8 || len(group[1]) != 4 || len(group[2]) != 4 || len(group[3]) != 4 || len(group[4]) != 12 {
				return false
			}
		}
	}
	val_trim_arr := strings.Split(val_trim, "")
	for _, v := range val_trim_arr {
		hit := false
		for _, o := range []string{"a", "b", "c", "d", "e", "f", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", } {
			if v == o {
				hit = true
				break
			}
		}
		if !hit {
			return false
		}
	}
	return true

}
func guid_list_chker(val string) bool {
	val_trim := strings.Trim(val, " ")
	if val_trim == "" {
		return false
	}
	if !strings.Contains(val_trim, ",") {
		return false
	}
	val_trim_arr := strings.Split(val_trim, ",")
	for _, v := range val_trim_arr {
		if strings.Replace(v, " ", "", -1) == "" {
			return false
		}
		if !guid_chker(v) {
			return false
		}
	}
	return true

}
func str_chker(val string) bool {
	return true

}
func float_list_chker(val string) bool {
	val_trim := strings.Trim(val, " ")
	if val_trim == "" {
		return false
	}
	if !strings.Contains(val_trim, ",") {
		return false
	}
	val_trim_arr := strings.Split(val_trim, ",")
	for _, v := range val_trim_arr {
		if strings.Replace(v, " ", "", -1) == "" {
			return false
		}
		if !float_chker(v) {
			return false
		}
	}
	return true

}
func zero_chker(val string) bool {
	return strings.Trim(val, " ") == "0"
}

type ValiRuleChker struct {
	vali_rule_info ValiRuleInfo
	IsValidated    bool
}

func (this *ValiRuleChker)Chk(val string) (int, string) {
	val_rune := []rune(val)
	str_len := len(val_rune)

	if this.vali_rule_info.HasRequired && val == "" {
		this.IsValidated = false
		return REQUIRED_ERR, this.vali_rule_info.Field + "必须输入值"
	}
	if this.vali_rule_info.HasLen {
		min_len := this.vali_rule_info.Len[0]
		max_len := this.vali_rule_info.Len[1]
		if min_len != -1 {
			if str_len < min_len {
				this.IsValidated = false
				return LEN_ERR, fmt.Sprintf("%s的长度不能小于%d", this.vali_rule_info.Field, min_len)
			}
		}
		if max_len != -1 {
			if str_len > max_len {
				this.IsValidated = false
				return LEN_ERR, fmt.Sprintf("%s的长度不能大于%d", this.vali_rule_info.Field, max_len)
			}
		}
	}
	if this.vali_rule_info.HasValues && len(this.vali_rule_info.Values) > 0 {
		for _, v := range val_rune {
			hit := false
			for _, allow_str := range this.vali_rule_info.Values {
				if string(v) == allow_str {
					hit = true
					break
				}
			}
			if !hit {
				this.IsValidated = false
				allow_strs := strings.Join(this.vali_rule_info.Values, "或")
				return VAL_ERR, fmt.Sprintf("%s的值为必须为%s", this.vali_rule_info.Field, allow_strs)
			}
		}
	}
	is_passed := false
	type_list := []string{}
	for rule_type, _ := range TYPE_LIST {

		if TYPE_LIST[rule_type] & this.vali_rule_info.Type != 0 {
			type_list = append(type_list, rule_type)

			fmt.Println("type_list", type_list)

			func_name := fmt.Sprintf("%s_chker", strings.ToLower(rule_type))
			func1, err1 := getChker(func_name)
			fmt.Println("func1", func1)
			fmt.Println("err1", err1)

			if err1 != nil {
				//未找到方法  不通过
			} else {
				ret := func1(val)
				fmt.Println("func1_ret", ret)

				if ret {
					is_passed = true
					break
				}
			}

		}
	}
	if !is_passed && len(type_list) != 0 {
		this.IsValidated = false
		return TYPE_ERR, fmt.Sprintf("%s的类型必须是%s", this.vali_rule_info.Field, strings.Join(type_list, "或"))
	}

	if this.vali_rule_info.HasRegx && this.vali_rule_info.Regx != "" {
		regx := this.vali_rule_info.Regx
		has_gen := strings.HasPrefix(regx, "/")
		if has_gen {
			regx = string([]rune(regx)[1:])
		}
		regx_flags := []string{
			"/", "/i", "/g", "/m", "/ig", "/gi", "/mi", "/im", "/gm", "/mg", "/igm", "/gim", "/img",
			"/mig", "/gmi", "/mgi",
		}
		end_flag := ""
		if has_gen {
			for _, flag := range regx_flags {
				if strings.HasSuffix(regx, flag) {
					end_flag = flag
					regx = string([]rune(regx)[:len(flag)])
					break
				}
			}
		}
		re_flags := 0
		_ = re_flags
		if strings.Contains(end_flag, "g") {
			//todo
		}
		if strings.Contains(end_flag, "i") {
			//todo
		}
		if strings.Contains(end_flag, "m") {
			//todo
		}


		reg := regexp.MustCompilePOSIX(regx)//最长匹配
		is_match := reg.Match([]byte(val))
		if !is_match {
			this.IsValidated = false
			return REGX_ERR, fmt.Sprintf("%s没有通过正则%s检查", this.vali_rule_info.Field, this.vali_rule_info.Regx)
		}
	}
	this.IsValidated = true
	return OK, "验证通过"

}

func getChker(chkType string) (func(val string) bool, error) {
	m := map[string]func(val string) bool{
		"p_int_chker":p_int_chker,
		"n_int_chker":n_int_chker,
		"int_chker":int_chker,
		"p_float_chker":p_float_chker,
		"n_float_chker":n_float_chker,
		"float_chker":float_chker,
		"bool_chker":bool_chker,
		"int_list_chker":int_list_chker,
		"str_list_chker":str_list_chker,
		"guid_chker":guid_chker,
		"guid_list_chker":guid_list_chker,
		"str_chker":str_chker,
		"float_list_chker":float_list_chker,
		"zero_chker":zero_chker,
	}
	if res, ok := m[chkType]; ok {
		return res, nil
	}

	return func(val string) bool {
		return false
	}, errors.New("未找到相应的检测方法")

}




//
//Call(FuncMap, "say", "hello")

