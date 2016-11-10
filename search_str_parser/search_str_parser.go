package search_str_parser

import (
	"strings"
	"regexp"
	"fmt"
	"strconv"
)


//规则是:s.word.0=123123&s.word1.1=asdfasd&&s.word1=asdfasd&&&&&
//word只能是字母数字下划线组合
//s.表示是开头,word是指搜索字段名,.0表示不显示,.1和空表示默认显示
//s.word.0=foo：是一个搜索字段,字段名为word,默认不显示在ui中,.0表示为不显示,foo是搜索的值
//s.word.1=foo：是一个搜索字段,字段名为word,默认不显示在ui中,.1表示为显示在UI中,foo是搜索的值
//s.word=foo：是一个搜索字段,字段名为word,默认不显示在ui中,无显示后缀表示为默认显示在UI中,foo是搜索的值
//:param args_str:搜索字符串,也就是url的request.querystring部分
//:return:[('word', False, '123123'), ('word1', True, 'asdfasd'), ('word1', True, 'asdfasd')]
func SearchStrParse(args_str string) ([][]interface{}) {
	if args_str == "" {
		return [][]interface{}{}
	}
	if strings.Contains(args_str, "#") {
		args_str = strings.Split(args_str, "#")[0]
	}

	if strings.HasSuffix(args_str, "?") {
		args_str = "&" + args_str[1:]
	}
	arr := strings.Split(args_str, "?")
	args_str = "&" + strings.Trim(arr[len(arr) - 1], "&")
	fmt.Println("args_str:", args_str)
	r, _ := regexp.Compile(`[&]+s\.([^=\.]*)(?:(\.0|\.1)*)?[^=]*\s*=\s*([^&]*)`)
	lists := r.FindAllStringSubmatch(args_str, -1)
	res := [][]interface{}{}

	for _, one := range lists {

		//fmt.Println(len(one))
		//fmt.Println(one)
		res = append(res, []interface{}{
			one[1],
			one[2] == ".1",
			one[3],
		})
	}
	return res

}

const (
	EQUALS = 1 << 0
	NOT_EQUALS = 1 << 2
	NOT_LIKE = 1 << 3
	NOT_LIKE_END = 1 << 4
	NOT_LIKE_START = 1 << 5
	NOT_IN = 1 << 6
	LIKE = 1 << 7
	LIKE_START = 1 << 8
	LIKE_END = 1 << 9
	IN = 1 << 10
	LESS_THAN = 1 << 11
	LESS_EQUALS = 1 << 12
	GREATER_THAN = 1 << 13
	GREATER_EQUALS = 1 << 14
	INT_RANGE_G_L = 1 << 15
	INT_RANGE_GE_LE = 1 << 16
	INT_RANGE_G_LE = 1 << 17
	INT_RANGE_GE_L = 1 << 18
	DATE_BETWEEN = 1 << 19
)

type SearchConditionParser struct {

}

func _int_chker(val string) bool {
	_, err := strconv.Atoi(val)
	return err == nil
}

func _float_chker(val string) bool {
	_, err := strconv.ParseFloat(val, 64)
	return err == nil
}
func Parse(search_str string) (int, string, string, string, []interface{}, string) {
	// &s.field = abc => field = abc
	if !strings.Contains(search_str, "!") && !strings.Contains(search_str, "*") && !strings.Contains(search_str, "~") {
		return EQUALS, search_str, "{f} = '" + search_str + "'", "{f} = ?", []interface{}{search_str, }, "{f} 等于 {v}"
	}
	//# 以感叹号开头的
	//# region &s.field = !abc => field <> abc
	search_str = strings.Trim(search_str, " ")
	if strings.HasPrefix(search_str, "!") {
		search_val := search_str[1:]
		if !strings.Contains(search_val, "*") &&  !strings.Contains(search_val, "~") {
			return NOT_EQUALS, search_val, "{f} <> '" + search_val + "'",
				"{f} <> ?",
				[]interface{}{search_val},
				"{f} 不等于 {v}"
		}
	}
	//&s.field =!*abc* => field not like '%abc%'
	if strings.HasPrefix(search_str, "!*") && strings.HasSuffix(search_str, "*") {
		search_val := search_str[2:len(search_str) - 1]
		return NOT_LIKE, search_val, "{f} not like '%" + search_val + "%'",
			"{f} not like ?",
			[]interface{}{"%" + search_val + "%", },
			"{f} 不包含 {v}"
	}
	//&s.field = !*abc => field not like '%abc'
	if strings.HasPrefix(search_str, "!*") && strings.HasSuffix(search_str, "*") {
		search_val := search_str[2:]
		return NOT_LIKE_END, search_val, "{f} not like '%" + search_val + "'",
			"{f} not like ?",
			[]interface{}{"%" + search_val, },
			"{f} 不以 {v} 结尾"

	}
	//&s.field = !abc* => field not like 'abc%'
	if strings.HasPrefix(search_str, "!") && strings.HasSuffix(search_str, "*") {
		search_val := search_str[1:len(search_str) - 1]
		return NOT_LIKE_START, search_val, "{f} not like '" + search_val + "%'",
			"{f} not like ?",
			[]interface{}{search_val + "%", },
			"{f} 不以 {v} 开头"
	}
	//&s.field = !~1,2,3~ =>相当于not in(...),实际将转化为 field<>1 or field<>2 or filed<>3
	if strings.HasPrefix(search_str, "!~") && strings.HasSuffix(search_str, "~") {
		search_val := search_str[2:len(search_str) - 1]
		vals := []string{}
		sql := []string{}
		param_sql := []string{}
		mapping := []interface{}{}

		for _, val := range strings.Split(search_val, ",") {
			vals = append(vals, val)
			sql = append(sql, "{f} <> '" + val + "'")
			param_sql = append(param_sql, "{f} <> ?")
			mapping = append(mapping, val)
		}
		return NOT_IN, search_val, strings.Join(sql, " and "), strings.Join(param_sql, " and "), mapping, "{f} 不等于 " + strings.Join(vals, " 或 ")
	}
	//&s.field = *abc* => field like '%abc%'
	if strings.HasPrefix(search_str, "*") && strings.HasSuffix(search_str, "*") {
		search_val := search_str[1:len(search_str) - 1]
		return LIKE, search_val, "{f} like '%" + search_val + "%'",
			"{f} like ?",
			[]interface{}{"%" + search_val + "%", },
			"{f} 包含 {v}"
	}
	//&s.field = abc* => field like 'abc%'
	if strings.HasSuffix(search_str, "*") && !strings.HasPrefix(search_str, "*") && !strings.HasPrefix(search_str, "!*") && !strings.HasPrefix(search_str, "!") {
		search_val := search_str[:len(search_str) - 1]
		return LIKE_START, search_val, "{f} like '" + search_val + "%'",
			"{f} like ?",
			[]interface{}{search_val + "%", },
			"{f} 以 {v} 开头"
	}
	//&s.field = *abc => field like '%abc'
	if strings.HasPrefix(search_str, "*") && !strings.HasSuffix(search_str, "*") {
		search_val := search_str[1:]
		return LIKE_END, search_val,
			"{f} like '%" + search_val + "'",
			"{f} like ?",
			[]interface{}{"%" + search_val, },
			"{f} 以 {v} 结尾"
	}
	//&s.field = ~1~ or ~1,2,3~ => 相当于in(...),实际将转化为 field=1 or field=2 or filed=3
	if strings.HasPrefix(search_str, "~") && strings.HasSuffix(search_str, "~") {
		search_val := search_str[1:len(search_str) - 1]
		vals := []string{}
		sql := []string{}
		param_sql := []string{}
		mapping := []interface{}{}
		for _, val := range strings.Split(search_val, ",") {
			vals = append(vals, val)
			sql = append(sql, "{f} = '" + val + "'")
			param_sql = append(param_sql, "{f} = ?")
			mapping = append(mapping, val)
		}
		return IN, search_val, strings.Join(sql, " or "), strings.Join(param_sql, " or "), mapping, "{f} 等于 " + strings.Join(vals, " 或 ")
	}
	//&s.field = ~1 => field<1 小于
	if strings.HasPrefix(search_str, "~") && !strings.HasPrefix(search_str, "~=") && !strings.HasSuffix(search_str, "~") {
		search_val := search_str[1:]
		if _int_chker(search_val) || _float_chker(search_val) {
			return LESS_THAN, search_val,
				"{f} < " + search_val,
				"{f} < ?",
				[]interface{}{search_val, },
				"{f} 小于 {v}"
		}
	}
	//&s.field = ~=1 => field<=1 小于等于
	if strings.HasPrefix(search_str, "~=") && !strings.HasSuffix(search_str, "~") {
		search_val := search_str[2:]
		return LESS_EQUALS, search_val,
			"{f} <= " + search_val,
			"{f} <= ?",
			[]interface{}{search_val, },
			"{f} 小于等于 {v}"
	}
	//&s.field = 1~ => field>1 大于
	if strings.HasSuffix(search_str, "~") && !strings.HasSuffix(search_str, "=~") && !strings.HasPrefix(search_str, "~") && strings.HasPrefix(search_str, "!~") {
		search_val := search_str[:len(search_str) - 1]
		if _int_chker(search_val) || _float_chker(search_val) {
			return GREATER_THAN, search_val,
				"{f} > " + search_val,
				"{f} > ?",
				[]interface{}{search_val, },
				"{f} 大于 {v}"
		}

	}
	//&s.field = 1=~ => field>1 大于等于
	if strings.HasSuffix(search_str, "=~") {
		search_val := search_str[:len(search_str) - 2]
		if _int_chker(search_val) || _float_chker(search_val) {
			return GREATER_EQUALS,
				search_val,
				"{f} >= " + search_val,
				"{f} >= ?",
				[]interface{}{search_val, },
				"{f} 大于等于 {v}"
		}
	}
	//&s.field = 0~1 => field<0 and field>1
	if strings.Contains(search_str, "~") && !strings.HasPrefix(search_str, "~") && !strings.HasSuffix(search_str, "~") && !strings.Contains(search_str, "=") && !strings.Contains(search_str, "~~") {
		search_val_arr := strings.Split(search_str, "~")
		if len(search_val_arr) == 2 {
			if (_int_chker(search_val_arr[0]) || _float_chker(search_val_arr[0])) && (_int_chker(search_val_arr[1]) || _float_chker(search_val_arr[1])) {
				return INT_RANGE_G_L, search_val_arr[0] + "," + search_val_arr[1],
					"{f} > " + search_val_arr[0] + " and {f} < " + search_val_arr[1],
					"{f} > ? and {f} < ?",
					[]interface{}{search_val_arr[0], search_val_arr[1]},
					"{f} 大于 {v} 小于 {v1}"
			}
		}
	}
	//&s.field = 0=~=1 => field>=0 and field<=1
	if strings.Contains(search_str, "=~=") {
		search_val_arr := strings.Split(search_str, "=~=")
		if len(search_val_arr) == 2 {
			if (_int_chker(search_val_arr[0]) || _float_chker(search_val_arr[0])) && (_int_chker(search_val_arr[1]) || _float_chker(search_val_arr[1])) {
				return INT_RANGE_GE_LE,
					search_val_arr[0] + "," + search_val_arr[1],
					"{f} >= " + search_val_arr[0] + " and {f} <= " + search_val_arr[1],
					"{f} >= ? and {f} <= ?",
					[]interface{}{search_val_arr[0], search_val_arr[1]},
					"{f} 大于等于 {v} 小于等于 {v1}"
			}
		}
	}
	//&s.field = 0~=1 => field>0 and field=<1
	if strings.Contains(search_str, "~=") {
		search_val_arr := strings.Split(search_str, "~=")
		if len(search_val_arr) == 2 {
			if (_int_chker(search_val_arr[0]) || _float_chker(search_val_arr[0])) && (_int_chker(search_val_arr[1]) || _float_chker(search_val_arr[1])) {
				return INT_RANGE_G_LE,
					search_val_arr[0] + "," + search_val_arr[1],
					"{f} > " + search_val_arr[0] + " and {f} <= " + search_val_arr[1],
					"{f} > ? and {f} <= ?",
					[]interface{}{search_val_arr[0], search_val_arr[1]},
					"{f} 大于 {v} 小于等于 {v1}"
			}
		}
	}
	//&s.field = 0=~1 => field>=0 and field<1
	if strings.Contains(search_str, "=~") {
		search_val_arr := strings.Split(search_str, "=~")
		if len(search_val_arr) == 2 {
			if (_int_chker(search_val_arr[0]) || _float_chker(search_val_arr[0])) && (_int_chker(search_val_arr[1]) || _float_chker(search_val_arr[1])) {
				return INT_RANGE_GE_L,
					search_val_arr[0] + "," + search_val_arr[1],
					"{f} >= " + search_val_arr[0] + " and {f} < " + search_val_arr[1],
					"{f} >= ? and {f} < ?",
					[]interface{}{search_val_arr[0], search_val_arr[1]},
					"{f} 大于等于 {v} 小于 {v1}"
			}
		}
	}
	//&s.field = 1~~1 => 1 between 1 用于时间
	if strings.Contains(search_str, "~~") {
		search_val_arr := strings.Split(search_str, "~~")
		if len(search_val_arr) == 2 {
			date_regx := `^\s*([12]\d{3})-(0{0,1}[0-9]|1{0,1}[012])-([012]{0,1}\d|3[01])\s+([01]{0,1}[0-9]|2[0-3]):[0-5]{0,1}[0-9]:[0-5]{0,1}[0-9]\s*$`
			r, _ := regexp.Compile(date_regx)
			res1 := r.FindAllStringSubmatch(search_val_arr[0], -1)
			res2 := r.FindAllStringSubmatch(search_val_arr[1], -1)
			if res1 != nil && res2 != nil {
				return DATE_BETWEEN, search_val_arr[0] + "," + search_val_arr[1],
					"({f} between '" + search_val_arr[0] + "' and '" + search_val_arr[1] + "')",
					"({f} between ? and ?)",
					[]interface{}{search_val_arr[0], search_val_arr[1]},
					"{f} 从 {v} 到 {v1}"
			}
		}
	}

	return 0, "", "", "", []interface{}{}, ""
}

type SearchItemInfo struct {
	SqlStr    string
	WhereStr  string
	Info      string
	Field     string
	IsShow    bool
	Value     string
	Text      string
	SearchStr string
	Type      int
	SearchVal string
	Mapping   []interface{}
}

func (this *SearchItemInfo) Init(field string, is_show bool, value string) {
	this.Field = field
	this.IsShow = is_show
	this.Value = value
	sep_pos := strings.Index(this.Value, "|")
	if sep_pos == -1 {
		this.Text = field
		this.SearchStr = this.Value
	} else {

		fmt.Println("sep_pos", sep_pos)
		this.Text = this.Value[:sep_pos]
		this.SearchStr = strings.TrimLeft(this.Value[sep_pos:], "|")
		if strings.Contains(this.SearchStr, "|") {
			this.SearchStr = this.SearchStr[0:strings.LastIndex(this.SearchStr, "|")]
		}
	}
	this.Type, this.SearchVal, this.SqlStr, this.WhereStr, this.Mapping, this.Info = Parse(this.SearchStr)
	v := ""
	v1 := ""
	arr := []string{"", ""}
	if this.Type == INT_RANGE_G_L || this.Type == INT_RANGE_G_LE || this.Type == INT_RANGE_GE_L || this.Type == INT_RANGE_GE_LE || this.Type == DATE_BETWEEN {
		switch this.Type {
		case DATE_BETWEEN:
			arr = strings.Split(this.SearchStr, "~~")
		case INT_RANGE_G_L:
			arr = strings.Split(this.SearchStr, "~")
		case INT_RANGE_G_LE:
			arr = strings.Split(this.SearchStr, "~=")
		case INT_RANGE_GE_L:
			arr = strings.Split(this.SearchStr, "=~")
		case INT_RANGE_GE_LE:
			arr = strings.Split(this.SearchStr, "=~=")
		default:

		}
		v = arr[0]
		v1 = arr[1]
	} else {
		v = this.SearchVal
		v1 = ""
	}
	this.SqlStr = strings.Replace(this.SqlStr, "{f}", field, -1)
	this.SqlStr = strings.Replace(this.SqlStr, "{v}", v, -1)
	this.SqlStr = strings.Replace(this.SqlStr, "{v1}", v1, -1)

	this.WhereStr = strings.Replace(this.WhereStr, "{f}", field, -1)
	this.WhereStr = strings.Replace(this.WhereStr, "{v}", v, -1)
	this.WhereStr = strings.Replace(this.WhereStr, "{v1}", v1, -1)

	this.Info = strings.Replace(this.Info, "{f}", this.Text, -1)
	this.Info = strings.Replace(this.Info, "{v}", v, -1)
	this.Info = strings.Replace(this.Info, "{v1}", v1, -1)

}
//规则是:st.field=title|0&s.field1=title|1
//field只能是字母数字下划线组合
//st.表示是搜索条件,field是指参与排序的字段名,第一条|线的左边是titile(可省略),0表示asc正序排列,1表示desc倒序排列
//st.id=标识符|0：以字段id正序排列(asc),别名为:标识符
//sd.price=价格|1: 以字段price倒序排列(desc),别名为:价格
//sd.datetime=1: 以字段price倒序排列(desc),别名为:datetime
//如果有重复的值,以最后一个值为准
//:param args_str:搜索字符串,也就是url的request.querystring部分
//:return:[('field_name', 'title', 0), [('field_name', 'title', 0), [('field_name', 'title', 0)]
func SortStrParse(args_str string) [][]interface{} {
	if args_str == "" {
		return [][]interface{}{}
	}
	sharp_pos := strings.Index(args_str, "#")
	if sharp_pos != -1 {
		args_str = args_str[0:sharp_pos]
	}
	arr := strings.Split(args_str, "?")
	args_str = "&" + strings.Trim(arr[len(arr) - 1], "&")
	fmt.Println("args_str:", args_str)
	r, _ := regexp.Compile(`[&]+st\.([^=\.]*)\s*=\s*([^\|&]*)\|*([10])`)
	lists := r.FindAllStringSubmatch(args_str, -1)
	res := [][]interface{}{}
	//fmt.Println("one++++++++", lists)
	for _, one := range lists {

		//fmt.Println(len(one))
		fmt.Println("==SortStrParse==", one)
		res = append(res, []interface{}{
			strings.Trim(one[1], " "),
			one[2],
			one[3],
		})
	}
	return res
}

const (
	SORT_ASC = 0
	SORT_DESC = 1
)

type SortItemInfo struct {
	Field    string
	Text     string
	SortType int
	SortStr  string
}

func (this *SortItemInfo) Init(field string, title string, sort_type string) {
	this.Field = field
	this.Text = title
	if sort_type == "0" {
		this.SortType = SORT_ASC
		this.SortStr = this.Field + " ASC"
	} else {
		this.SortType = SORT_DESC
		this.SortStr = this.Field + " DESC"
	}

}

type SearchAdapter struct {
	SqlStr          string
	WhereStr        string
	SortStr         string
	SqlStrList      []string
	WhereStrList    []string
	SortStrList     []string
	SortFieldList   []string
	Sorts           []SortItemInfo
	Mapping         []interface{}
	HasAnyShow      bool
	SearchFieldList []string
	Items           map[string]SearchItemInfo
}

func inlist(val_to_chk string, vals []string) bool {
	for _, val := range vals {
		if val_to_chk == val {
			//命中返回true
			return true
		}
	}
	//没命中返回false
	return false
}
func (this *SearchAdapter) Init(url string, allow_fields []string) {
	this.Items = map[string]SearchItemInfo{}
	for _, one := range SearchStrParse(url) {
		word := one[0].(string)
		is_show := one[1].(bool)
		val := one[2].(string)
		if len(allow_fields) > 0 {
			if !inlist(word, allow_fields) {
				continue
			}
		}
		//word 只能是数字和字母组成
		r, _ := regexp.Compile("^[a-zA-Z0-9_]*$")
		if !r.MatchString(word) {
			continue
		}
		search_info := SearchItemInfo{}
		search_info.Init(word, is_show, val)
		this.SqlStrList = append(this.SqlStrList, search_info.SqlStr)
		this.WhereStrList = append(this.WhereStrList, search_info.WhereStr)
		this.Mapping = append(this.Mapping, search_info.Mapping...)
		this.HasAnyShow = is_show
		this.SearchFieldList = append(this.SearchFieldList, search_info.Field)
		this.Items[search_info.Field] = search_info
	}

	for _, row := range SortStrParse(url) {
		word := row[0].(string)
		title := row[1].(string)
		sort_type := row[2].(string)
		if len(allow_fields) > 0 {
			if !inlist(word, allow_fields) {
				continue
			}
		}
		//word 只能是数字和字母组成
		r, _ := regexp.Compile("^[a-zA-Z0-9_]*$")
		if !r.MatchString(word) {
			continue
		}
		sort_info := SortItemInfo{}
		sort_info.Init(word, title, sort_type)
		fmt.Println("sort_info",sort_info)
		this.SortStrList = append(this.SortStrList, sort_info.SortStr)
		this.SortFieldList = append(this.SortFieldList, sort_info.Field)
		this.Sorts = append(this.Sorts, sort_info)

	}
	if len(this.SqlStrList) > 0 {
		this.SqlStr = "(" + strings.Join(this.SqlStrList, ") and (") + ")"
	}
	if len(this.WhereStrList) > 0 {
		this.WhereStr = "(" + strings.Join(this.WhereStrList, ") and (") + ")"
	}
	this.SortStr = strings.Join(this.SortStrList, " , ")
}