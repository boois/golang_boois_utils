// 萧鸣 boois@qq.com
package utils

import (
	"strings"
	"fmt"
	"strconv"
	"regexp"
	"net/url"
)
/*
规则是:s.word.0=中文字段名|*搜索条件*|值1中文名,值2中文名&s.word1.1=asdfasd&&s.word1=asdfasd&&&&&st.id.0=描述(从高到低)
&st.time.0=0&st.id.0=时间倒序|1  .0=UI上不显示, 等号后面的值 0=desc,非0=asc
word只能是字母数字下划线组合
s.表示是开头,word是指搜索字段名,.0表示不显示,.1和空表示默认显示
s.word.0=foo：是一个搜索字段,字段名为word,默认不显示在ui中,.0表示为不显示,foo是搜索的值
s.word.1=foo：是一个搜索字段,字段名为word,默认不显示在ui中,.1表示为显示在UI中,foo是搜索的值
s.word=foo：是一个搜索字段,字段名为word,默认不显示在ui中,无显示后缀表示为默认显示在UI中,foo是搜索的值
:param condtion_str:搜索字符串,也就是url的request.querystring部分
:return:[('word', False, '123123'), ('word1', True, 'asdfasd'), ('word1', True, 'asdfasd')]
*/

const EQUALS = 1 << 0
const NOT_EQUALS = 1 << 2
const NOT_LIKE = 1 << 3
const NOT_LIKE_END = 1 << 4
const NOT_LIKE_START = 1 << 5
const NOT_IN = 1 << 6
const LIKE = 1 << 7
const LIKE_START = 1 << 8
const LIKE_END = 1 << 9
const IN = 1 << 10
const LESS_THAN = 1 << 11
const LESS_EQUALS = 1 << 12
const GREATER_THAN = 1 << 13
const GREATER_EQUALS = 1 << 14
const INT_RANGE_G_L = 1 << 15
const INT_RANGE_GE_LE = 1 << 16
const INT_RANGE_G_LE = 1 << 17
const INT_RANGE_GE_L = 1 << 18
const DATE_BETWEEN = 1 << 19

type BooisSearchStrParseResultInfo struct {
	FieldName string // 字段名称
	IsShow    bool   // 在UI上是否显示
	Value     string // 值
}

type BooisSortItemInfo struct {
	FieldName string // 字段名称
	Text string // 排序中文描述: 年龄从大到小
	IsAsc bool // 是否正序排列
	IsShow bool // 是否在UI中显示
	SortStr string // 排序字符串  field desc|field asc
}

type BooisSearchConditionParseResultInfo struct {
		ConditionType int
		ValueStr string
		WhereStr string
		Mapping []interface{}
		Info string
	}

func ParseSearchStr(condtion_str string, search_prefix string,sort_prefix string) ([]BooisSearchStrParseResultInfo,[]BooisSortItemInfo){
	if search_prefix == "" { search_prefix = "s." }
	if sort_prefix == "" { sort_prefix = "st." }
	search_result := []BooisSearchStrParseResultInfo{}
	sort_result := []BooisSortItemInfo{}
	if condtion_str == ""{
		return search_result,sort_result
	}else{
		// 处理地址路径
		condtion_str,_ = url.QueryUnescape(condtion_str)
		// 1.把#号之后的字符串全部去掉
		if idx := strings.Index(condtion_str,"#");idx != -1 {
			condtion_str = condtion_str[:strings.Index(condtion_str,"#")]
		}
		// 2.获取问号之后的内容
		condtion_str = condtion_str[strings.LastIndex(condtion_str,"?")+1:]
		// 3.将参数全部基于&拆解出来
		condition_arr := strings.Split(condtion_str,"&")
		for x := range condition_arr {
			str:= condition_arr[x]
			if str == ""{continue}
			if strings.HasPrefix(str, search_prefix) {
				idx := strings.Index(str, "=")
				if idx == -1 {continue}
				k := str[:idx]
				k = strings.TrimPrefix(k, search_prefix)
				k = strings.Trim(k, " ")
				if k == "" { continue }
				v := str[idx + 1:]
				v = strings.Trim(v, " ")
				if v == "" {continue}
				is_show := true
				if strings.HasSuffix(k, ".0") {
					is_show = false
				}
				k = strings.TrimRight(k,".0")
				k = strings.TrimRight(k,".1")
				search_result = append(search_result, BooisSearchStrParseResultInfo{FieldName:k, Value:v, IsShow:is_show})
			}
			if strings.HasPrefix(str, sort_prefix) {
				idx := strings.Index(str, "=")
				k := ""
				v := ""
				if idx == -1 {
					k = str
				}else{
					k = str[:idx]
					v = str[idx + 1:]
					v = strings.Trim(v, " ")
				}
				k = strings.TrimPrefix(k, sort_prefix)
				k = strings.Trim(k, " ")
				if k == "" { continue }
				is_show := true
				is_asc := true
				if strings.HasSuffix(k, ".0") {
					is_show = false
				}
				txt := ""
				if v != "" {
					__idx := strings.Index(v,"|")
					if __idx != -1 {
						txt= v[0:__idx]
						v = v[__idx+1:]
					}
				}
				if v == "0"{
					is_asc = false
				}
				k = strings.TrimRight(k,".0")
				k = strings.TrimRight(k,".1")
				sort_str := ""
				if is_asc {
					sort_str = fmt.Sprintf("%s asc",k)
					if txt == "" { txt = k+"正序"}
				}else{
					sort_str = fmt.Sprintf("%s desc",k)
					if txt == "" { txt = k+"倒序"}
				}
				sort_result = append(sort_result, BooisSortItemInfo{FieldName:k,IsShow:is_show,Text:txt, IsAsc:is_asc, SortStr:sort_str})
			}
		}
		return search_result,sort_result
	}
}

func ParseSearchCondition(condition_val_str string) BooisSearchConditionParseResultInfo{
	/*
    根据搜索字符串来解析出搜索条件
    :condtion_str:用来解析的condtion_str
    :return:(code, search_val, sql_str, where_str, mapping, info)
    有以下几种形式(用了四种判定符各自组合:!,*,~,=):
    1.以！开头的表示否
    &s.field = !abc => field <> abc
    &s.field =!*abc* => field not like '%abc%'
    &s.field = !*abc => field not like '%abc'
    &s.field = !abc* => field not like 'abc%'
    &s.field = !~1,2,3~ =>相当于not in(...),实际将转化为 field<>1 or field<>2 or filed<>3
    2.以波浪线为范围的
    &s.field = 2016-1-1 1:1:1~~2016-1-1 1:1:1 => field between '2016-1-1 1:1:1' && '2016-1-1 1:1:1' 用于时间
    &s.field = ~1~ or ~1,2,3~ => 相当于in(...),实际将转化为 field=1 or field=2 or filed=3
    &s.field = ~1 => field<1 小于
    &s.field = ~=1 => field<=1 小于等于
    &s.field = 1~ => field>1 大于
    &s.field = 1=~ => field>1 大于等于
    &s.field = 0~1 => field<0 && field>1
    &s.field = 0=~=1 => field>=0 && field<=1
    &s.field = 0~=1 => field>0 && field=<1
    &s.field = 0=~1 => field>=0 && field<1
    3.带星号通配符的
    &s.field = *abc* => field like '%abc%'
    &s.field = abc* => field like 'abc%'
    &s.field = *abc => field like '%abc'
    4.什么都没有的
    &s.field = abc  => field = abc
    #(\;|'|\sand\s|exec\s|\sinsert\s|select|delete|update|\scount|\*|\!|\~|\%|\schr\(|\smid\(|master|truncate|\schar|declare)
    */

	_int_chker := func (val string) bool {
		_, err := strconv.Atoi(val)
		return err == nil
	}
	_float_chker := func(val string) bool {
		_, err := strconv.ParseFloat(val, 64)
		return err == nil
	}

	result := BooisSearchConditionParseResultInfo{}
	// 先判断什么符号都没有的情况
	// &s.field = abc => field = abc
	if strings.Index(condition_val_str,"!") == -1 &&strings.Index(condition_val_str,"*") == -1 &&strings.Index(condition_val_str,"~") == -1{
		result.ConditionType = EQUALS
		result.ValueStr = condition_val_str
		result.WhereStr = "%s = ?"
		result.Mapping = []interface{}{condition_val_str}
		result.Info = "%s 等于 %s"
		return result
	}
	condition_val_str = strings.Trim(condition_val_str," ")
	// 以感叹号开头的
	// &s.field = !abc => field <> abc
	if strings.HasPrefix(condition_val_str,"!"){
		search_val := condition_val_str[1:]
		if strings.Index(search_val,"*") == -1 && strings.Index(search_val,"~") == -1{
			result.ConditionType = NOT_EQUALS
			result.ValueStr = search_val
			result.WhereStr = "%s <> ?"
			result.Mapping = []interface{}{search_val}
			result.Info = "%s 不等于 %s"
			return result
		}
	}
	//&s.field =!*abc* => field not like '%abc%'
	if strings.HasPrefix(condition_val_str,"!*") && strings.HasSuffix(condition_val_str,"*"){
		search_val := condition_val_str[2:len(condition_val_str)-1]
		result.ConditionType = NOT_LIKE
		result.ValueStr = search_val
		result.WhereStr = "%s not like ?"
		result.Mapping = []interface{}{"%" + search_val + "%"}
		result.Info = "%s 不包含 %s"
		return result
	}
	// &s.field = !*abc => field not like '%abc'
	if strings.HasPrefix(condition_val_str,"!*") && !strings.HasSuffix(condition_val_str,"*"){
		search_val := condition_val_str[2:]
		result.ConditionType = NOT_LIKE_END
		result.ValueStr = search_val
		result.WhereStr = "%s not like ?"
		result.Mapping = []interface{}{"%" + search_val}
		result.Info = "%s 不以 %s 结尾"
		return result
	}
	//&s.field = !abc* => field not like 'abc%'
	if strings.HasPrefix(condition_val_str,"!") && strings.HasSuffix(condition_val_str,"*"){
		search_val := condition_val_str[1:len(condition_val_str)-1]
		result.ConditionType = NOT_LIKE_START
		result.ValueStr = search_val
		result.WhereStr = "%s not like ?"
		result.Mapping = []interface{}{search_val + "%"}
		result.Info = "%s 不以 %s 开头"
		return result
	}
	//&s.field = !~1,2,3~ =>相当于not in(...),实际将转化为 field<>1 or field<>2 or filed<>3
	if strings.HasPrefix(condition_val_str,"!~") && strings.HasSuffix(condition_val_str,"~"){
		search_val := condition_val_str[2:len(condition_val_str)-1]
		vals := []string{}
		param_sql := []string{}
		mapping := []interface{}{}
		arr := strings.Split(search_val,",")
		for i := range arr{
			val := arr[i]
			vals = append(vals,val)
			param_sql = append(param_sql,"%s <> ?")

			mapping = append(mapping,val)
		}
		result.ConditionType = NOT_IN
		result.ValueStr = search_val
		result.WhereStr = strings.Join(param_sql," and ")
		result.Info = "%s 不等于 " + strings.Join(vals," 或 ")
		result.Mapping = mapping
		return result
	}
	// 带星号通配符的
	// &s.field = *abc* => field like '%abc%'
	if strings.HasPrefix(condition_val_str,"*") && strings.HasSuffix(condition_val_str,"*"){
		search_val := condition_val_str[1:len(condition_val_str)-1]
		result.ConditionType = LIKE
		result.ValueStr = search_val
		result.WhereStr = "%s like ?"
		result.Mapping = []interface{}{"%" + search_val + "%"}
		result.Info = "%s 包含 %s"
		return result
	}
	//&s.field = abc* => field like 'abc%'
	if strings.HasSuffix(condition_val_str,"*") &&! strings.HasPrefix(condition_val_str,"*") &&! strings.HasPrefix(condition_val_str,"!*") &&! strings.HasPrefix(condition_val_str,"!"){
		search_val := condition_val_str[:len(condition_val_str)-1]
		result.ConditionType = LIKE_START
		result.ValueStr = search_val
		result.WhereStr = "%s like ?"
		result.Mapping = []interface{}{search_val + "%"}
		result.Info = "%s 以 %s 开头"
		return result
	}
	// &s.field = *abc => field like '%abc'
	if strings.HasPrefix(condition_val_str,"*") && ! strings.HasSuffix(condition_val_str,"*"){
		search_val := condition_val_str[1:]
		result.ConditionType = LIKE_END
		result.ValueStr = search_val
		result.WhereStr = "%s like ?"
		result.Mapping = []interface{}{"%" + search_val}
		result.Info = "%s 以 %s 结尾"
		return result
	}
	// 以波浪线为范围的
	// &s.field = ~1~ or ~1,2,3~ => 相当于in(...),实际将转化为 field=1 or field=2 or filed=3
	if strings.HasPrefix(condition_val_str,"~") && strings.HasSuffix(condition_val_str,"~"){
		search_val := condition_val_str[1:len(condition_val_str)-1]
		vals := []string{}
		param_sql := []string{}
		mapping := []interface{}{}
		arr:=strings.Split(search_val,",")
		for i := range arr {
			val := arr[i]
			vals = append(vals,val)
			param_sql = append(param_sql,"%s = ?")
			mapping = append(mapping,val)
		}
		result.ConditionType = IN
		result.ValueStr = search_val
		result.WhereStr = strings.Join(param_sql," or ")
		result.Info = "%s 等于 " + strings.Join(vals," 或 ")
		result.Mapping = mapping
		return result
	}
	// &s.field = ~1 => field<1 小于
	if strings.HasPrefix(condition_val_str,"~") && ! strings.HasPrefix(condition_val_str,"~=") && ! strings.HasSuffix(condition_val_str,"~"){
		search_val := condition_val_str[1:]
		if _int_chker(search_val) || _float_chker(search_val){
			result.ConditionType = LESS_THAN
			result.ValueStr = search_val
			result.WhereStr = "%s < ?"
			result.Mapping = []interface{}{search_val}
			result.Info = "%s 小于 %s"
			return result
		}
	}
	//&s.field = ~=1 => field<=1 小于等于
	if strings.HasPrefix(condition_val_str,"~=") && ! strings.HasSuffix(condition_val_str,"~"){
		search_val := condition_val_str[2:]
		if _int_chker(search_val) || _float_chker(search_val){
			result.ConditionType = LESS_EQUALS
			result.ValueStr = search_val
			result.WhereStr = "%s <= ?"
			result.Mapping = []interface{}{search_val}
			result.Info = "%s 小于等于 %s"
			return result
		}
	}
	// &s.field = 1~ => field>1 大于
	if strings.HasSuffix(condition_val_str,"~") &&! strings.HasSuffix(condition_val_str,"=~") &&! strings.HasPrefix(condition_val_str,"~") &&! strings.HasPrefix(condition_val_str,"!~"){
		search_val := condition_val_str[0:len(condition_val_str)-1]
		if _int_chker(search_val) || _float_chker(search_val){
			result.ConditionType = GREATER_THAN
			result.ValueStr = search_val
			result.WhereStr = "%s > ?"
			result.Mapping = []interface{}{search_val}
			result.Info = "%s 大于 %s"
			return result
		}
	}
	// &s.field = 1=~ => field>1 大于等于
	if strings.HasSuffix(condition_val_str,"=~"){
		search_val := condition_val_str[:len(condition_val_str)-2]
		if _int_chker(search_val) ||  _float_chker(search_val){
			result.ConditionType = GREATER_EQUALS
			result.ValueStr = search_val
			result.WhereStr = "%s >= ?"
			result.Mapping = []interface{}{search_val}
			result.Info = "%s 大于等于 %s"
			return result
		}
	}
	// &s.field = 0~1 => field<0 && field>1
	if strings.Index(condition_val_str,"~") != -1 &&! strings.HasPrefix(condition_val_str,"~") &&! strings.HasSuffix(condition_val_str,"~") &&strings.Index(condition_val_str,"=") == -1 &&strings.Index(condition_val_str,"~~") == -1{
		search_val_arr := strings.Split(condition_val_str,"~")
		if len(search_val_arr) == 2 {
			if (_int_chker(search_val_arr[0]) ||
				_float_chker(search_val_arr[0])) &&(_int_chker(search_val_arr[1]) ||
				_float_chker(search_val_arr[1])) {
				result.ConditionType = INT_RANGE_G_L
				result.ValueStr = fmt.Sprintf("%s,%s", search_val_arr[0], search_val_arr[1])
				result.WhereStr = "%s > ? and %s < ?"
				result.Mapping = []interface{}{search_val_arr[0], search_val_arr[1]}
				result.Info = "%s 大于 %s 小于 %s"
				return result
			}
		}
	}
	// &s.field = 0=~=1 => field>=0 && field<=1
	if strings.Index(condition_val_str,"=~=") != -1{
		search_val_arr := strings.Split(condition_val_str,"=~=")
		if len(search_val_arr) == 2{
			if (_int_chker(search_val_arr[0]) ||
						_float_chker(search_val_arr[0])) &&(_int_chker(search_val_arr[1]) ||
							 _float_chker(search_val_arr[1])){
				result.ConditionType = INT_RANGE_GE_LE
				result.ValueStr = fmt.Sprintf("%s,%s" , search_val_arr[0], search_val_arr[1])
				result.WhereStr = "%s >= ? and %s <= ?"
				result.Mapping = []interface{}{search_val_arr[0], search_val_arr[1]}
				result.Info = "%s 大于等于 %s 小于等于 %s"
				return result
			}
		}
	}
	// &s.field = 0~=1 => field>0 && field=<1
	if strings.Index(condition_val_str,"~=") != -1{
		search_val_arr := strings.Split(condition_val_str,"~=")
		if len(search_val_arr) == 2 {
			if (_int_chker(search_val_arr[0]) ||
				_float_chker(search_val_arr[0])) &&(_int_chker(search_val_arr[1]) ||
				_float_chker(search_val_arr[1])) {
				result.ConditionType = INT_RANGE_G_LE
				result.ValueStr = fmt.Sprintf("%s,%s", search_val_arr[0], search_val_arr[1])
				result.WhereStr = "%s > ? and %s <= ?"
				result.Mapping = []interface{}{search_val_arr[0], search_val_arr[1]}
				result.Info = "%s 大于 %s 小于等于 %s"
				return result
			}
		}
	}
	// &s.field = 0=~1 => field>=0 && field<1
	if strings.Index(condition_val_str,"=~") != -1 {
		search_val_arr := strings.Split(condition_val_str,"=~")
		if len(search_val_arr) == 2 {
			if (_int_chker(search_val_arr[0]) ||
						_float_chker(search_val_arr[0])) &&(_int_chker(search_val_arr[1]) ||
							 _float_chker(search_val_arr[1])){
					result.ConditionType = INT_RANGE_GE_L
					result.ValueStr = fmt.Sprintf("%s,%s" , search_val_arr[0], search_val_arr[1])
					result.WhereStr = "%s >= ? and %s < ?"
					result.Mapping = []interface{}{search_val_arr[0], search_val_arr[1]}
					result.Info = "%s 大于等于 %s 小于 %s"
					return result
			}
		}
	}
	// &s.field = 1~~1 => 1 between 1 用于时间
	if strings.Index(condition_val_str,"~~") != -1{
		search_val_arr := strings.Split(condition_val_str,"~~")
		if len(search_val_arr) == 2{
			date_regx := "^\\s*([12]\\d{3})-(0{0,1}[0-9]|1{0,1}[012])-([012]{0,1}\\d|3[01])\\s+([01]{0,1}[0-9]|2[0-3]):[0-5]{0,1}[0-9]:[0-5]{0,1}[0-9]\\s*$"
			re := regexp.MustCompile(date_regx)
			if re.MatchString(search_val_arr[0]) && re.MatchString(search_val_arr[1]){
				result.ConditionType = DATE_BETWEEN
				result.ValueStr = fmt.Sprintf("%s,%s" , search_val_arr[0], search_val_arr[1])
				result.WhereStr = "(%s between ? and ?)"
				result.Mapping = []interface{}{search_val_arr[0], search_val_arr[1]}
				result.Info = "%s 从 %s 到 %s"
				return result
			}
		}
	}
	return result
}

type BooisSearchItemInfo struct {
	FieldName     string        // 字段名称
	IsShow        bool          // 是否显示
	Text          string        // 字段描述文字,通常表现形式为text|*abc*
	ValueRawStr   string        // url参数中的原始v值
	ConditionStr  string        // 参数中去掉了 text| 后的条件值
	ConditionType int           // 搜索类型
	SearchValue   string        // 真实的搜索值
	WhereStr      string        // 搜索字符串  field = ?
	Mapping       []interface{} // 向数据库传递的值
	ConditionInfo string        // 显示给UI的搜索条件信息
	VText1        string        // 第一个值的描述文字 字段中文名|*abc*|第一个值,第二个值
	VText2        string        // 第二个值的描述文字 字段中文名|*abc*|第一个值,第二个值
}

func (this *BooisSearchItemInfo) Load(info BooisSearchStrParseResultInfo)  {
	this.FieldName = info.FieldName
	this.IsShow = info.IsShow
	val,_ := url.QueryUnescape(info.Value) // 原始value字串中可能有中文,先解开
	this.ValueRawStr = val
	idx := strings.Index(this.ValueRawStr,"|") // 查找第一个出现的竖线
	if idx == -1 { // 没有竖线的话,也就是没有text|*abc*中的text值,所以text设为字段名称
		this.Text = this.FieldName
		this.ConditionStr = this.ValueRawStr
	}else{
		this.Text = this.ValueRawStr[0:idx] // 将第一条竖线左边的部分作为text
		if this.Text == "" {this.Text = this.FieldName} // 如果第一个竖线text为空则使用字段名称
		this.ConditionStr = this.ValueRawStr[idx+1:] // 第一条竖线右边的部分作为搜索条件
		if strings.Index(this.ConditionStr,"|")!=-1{//如果还存在竖线,只取竖线前面的作为搜索条件,并将竖线右边的作为值的描述
			// 字段中文名|*abc*|第一个值,第二个值
			vtext:= this.ConditionStr[strings.Index(this.ConditionStr,"|")+1:]
			idx := strings.Index(vtext,",")
			if idx == -1 {
				this.VText1 = vtext
			}else{
				this.VText1 = vtext[0:idx]
				this.VText2 = vtext[idx+1:]
			}
			this.ConditionStr = this.ConditionStr[0:strings.Index(this.ConditionStr,"|")]
		}
	}
	// 将搜索条件解析出来
	re:=ParseSearchCondition(this.ConditionStr)
	this.ConditionType = re.ConditionType
	this.Mapping = re.Mapping
	this.WhereStr = re.WhereStr
	this.SearchValue = re.ValueStr
	this.ConditionInfo = re.Info
	// 将wherestr和info中的空位填满
	v:=""
	v1:=""
	var arr []string
	if this.ConditionType == INT_RANGE_G_L ||
                        this.ConditionType == INT_RANGE_G_LE ||
                        this.ConditionType == INT_RANGE_GE_L ||
                        this.ConditionType == INT_RANGE_GE_LE ||
                        this.ConditionType == DATE_BETWEEN{

            if this.ConditionType == DATE_BETWEEN{
                arr = strings.Split(this.ConditionStr,"~~")
			}else if this.ConditionType == INT_RANGE_G_L{
                arr = strings.Split(this.ConditionStr,"~")
            }else if this.ConditionType == INT_RANGE_G_LE{
                arr = strings.Split(this.ConditionStr,"~=")
            }else if this.ConditionType == INT_RANGE_GE_L{
                arr = strings.Split(this.ConditionStr,"=~")
            }else if this.ConditionType == INT_RANGE_GE_LE {
				arr = strings.Split(this.ConditionStr,"=~=")
			}
            v = arr[0]
            v1 = arr[1]
     }else{
            v = this.SearchValue
            v1 = ""
	 }
	this.WhereStr = strings.Replace(this.WhereStr,"%s",this.FieldName,-1)
	if this.VText1 != "" {v = this.VText1}
	if this.ConditionType == NOT_IN || this.ConditionType == IN {
		this.ConditionInfo = fmt.Sprintf(this.ConditionInfo,this.Text)
	}else{
		if v1 != ""{
			if this.VText2 != "" {v1 = this.VText2}
			this.ConditionInfo = fmt.Sprintf(this.ConditionInfo,this.Text,v,v1)
		}else{
			this.ConditionInfo = fmt.Sprintf(this.ConditionInfo,this.Text,v)
		}
	}

}

type BooisSearchConditionAdapter struct {
	SearchFieldList []interface{}
	WhereStr        string
	WhereStrList    []string
	Mapping         []interface{}
	SearchItems     map[string]BooisSearchItemInfo
	HasAnyShow      bool //是否至少有一个显示的值
	SortStrList     []string
	SortStr         string
	SortItems       []BooisSortItemInfo
}

func (this *BooisSearchConditionAdapter) Parse(condtion_str,search_prefix,sort_prefix string) {
	// 先从url解出所有条件
	search_arr,sort_arr := ParseSearchStr(condtion_str,search_prefix,sort_prefix)
	this.SearchFieldList = []interface{}{}
	this.WhereStrList = []string{}
	this.Mapping = []interface{}{}
	this.HasAnyShow = false
	this.SearchItems = map[string]BooisSearchItemInfo{}
	for x := range search_arr {
		obj := search_arr[x]
		s_info := BooisSearchItemInfo{}
		s_info.Load(obj)
		this.SearchItems[obj.FieldName] = s_info
		this.SearchFieldList = append(this.SearchFieldList,obj.FieldName)
		this.WhereStrList = append(this.WhereStrList,s_info.WhereStr)
		this.Mapping = append(this.Mapping,s_info.Mapping...)
		if obj.IsShow {this.HasAnyShow = true}
	}
	for x := range sort_arr {
		obj := sort_arr[x]
		this.SortStrList = append(this.SortStrList,obj.SortStr)
		this.SortItems = append(this.SortItems,obj)
	}
	this.SortStr = strings.Join(this.SortStrList,",")
	if len(this.WhereStrList) == 0{
		this.WhereStr = ""
	}else{
		this.WhereStr = "("+strings.Join(this.WhereStrList,") and (")+")"
	}
}
