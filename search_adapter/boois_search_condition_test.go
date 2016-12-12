// 萧鸣 boois@qq.com
package utils

import (
	"testing"
	"fmt"
)

func Test_BooisSearchStrParser(t *testing.T) {// 测试用例必须以Test开头
	ParseSearchStr("http://ccc.com?????s.dfadfasf&&s.asdfasf = asdf&s.asdfasdf = asdf=123|//123123//123123123","s.","st.")
	ParseSearchStr("&s.dfadfasf&&s.asdfasf = asdf&s.asdfasdf = asdf=123|//123123//123123123","s.","st.")
	ParseSearchStr("http://ccc.com?????s.dfadfasf&&s.asdfasf = asdf&s.asdfasdf = asdf=123|","s.","st.")
}
func Test_BooisSearchConditionAdapter(t *testing.T) {// 测试用例必须以Test开头
	p := BooisSearchConditionAdapter{}
	p.Parse("http://ccc.com?????s.field1.1=我去|*a|AA&&s.field1.0 =中文数字|0~=1|零,一&s.field2 = 测试|asdf=123|喵&st.time.0=0&st.id.1=时间正序|1","s.","st.")
	fmt.Println("this.SearchFieldList:",p.SearchFieldList)
	if p.SearchFieldList[0] != "field" { t.Error("出错!")}
	if p.SearchFieldList[1] != "field" { t.Error("出错!")}
	if p.SearchFieldList[2] != "field2" { t.Error("出错!")}
	fmt.Println("this.WhereStrList:",p.WhereStrList)
	if p.WhereStrList[0] != "field like ?" { t.Error("出错!")}
	if p.WhereStrList[1] != "field > ? and field <= ?" { t.Error("出错!")}
	if p.WhereStrList[2] != "field2 = ?" { t.Error("出错!")}
	fmt.Println("this.WhereStr:",p.WhereStr)
	if p.WhereStr != "(field like ?) or (field > ? and field <= ?) or (field2 = ?)" { t.Error("出错!")}
	fmt.Println("this.ValueMapping:",p.Mapping)
	if p.Mapping[0] != "%a" {t.Error("出错!")}
	if p.Mapping[1] != "0" {t.Error("出错!")}
	if p.Mapping[2] != "1" {t.Error("出错!")}
	if p.Mapping[3] != "asdf=123" {t.Error("出错!")}
	fmt.Println("this.HasAnyShow:",p.HasAnyShow)
	if p.HasAnyShow != true{t.Error("出错!")}
	fmt.Println("this.SortStr:",p.SortStr)
	if p.SortStr != "time desc,id asc" {t.Error("出错!")}
	fmt.Println("this.SortStrList:",p.SortStrList)
	if p.SortStrList[0] != "time desc" {t.Error("出错!")}
	if p.SortStrList[1] != "id asc" {t.Error("出错!")}
	fmt.Println("this.SortItems:",p.SortItems)
	if len(p.SortItems)!=2 {t.Error("出错!")}
	if p.SortItems[0].FieldName != "time" {t.Error("出错!")}
	if p.SortItems[0].Text != "time倒序" {t.Error("出错!")}
	if p.SortItems[0].IsAsc != false {t.Error("出错!")}
	if p.SortItems[0].SortStr != "time desc" {t.Error("出错!")}
	if p.SortItems[1].FieldName != "id" {t.Error("出错!")}
	if p.SortItems[1].Text != "时间正序" {t.Error("出错!")}
	if p.SortItems[1].IsAsc != true {t.Error("出错!")}
	if p.SortItems[1].SortStr != "id asc" {t.Error("出错!")}
	fmt.Println("this.SearchItems:",p.SearchItems)
	if len(p.SearchItems)!=2 {t.Error("出错!")}
	if p.SearchItems["field"].FieldName != "field" {t.Error("出错!")}
	if p.SearchItems["field"].IsShow != false {t.Error("出错!")}
	if p.SearchItems["field"].Text != "中文数字" {t.Error("出错!")}
	if p.SearchItems["field"].ValueRawStr != "中文数字|0~=1|零,一" {t.Error("出错!")}
	if p.SearchItems["field"].ConditionStr != "0~=1" {t.Error("出错!")}
	if p.SearchItems["field"].ConditionType != INT_RANGE_G_LE {t.Error("出错!")}
	if p.SearchItems["field"].SearchValue != "0,1" {t.Error("出错!")}
	if p.SearchItems["field"].WhereStr != "field > ? and field <= ?" {t.Error("出错!")}
	if p.SearchItems["field"].ConditionInfo != "中文数字 大于 零 小于等于 一" {t.Error("出错!")}
	if p.SearchItems["field"].VText1 != "零" {t.Error("出错!")}
	if p.SearchItems["field"].VText2 != "一" {t.Error("出错!")}
	if p.SearchItems["field"].Mapping[0] != "0" {t.Error("出错!")}
	if p.SearchItems["field"].Mapping[1] != "1" {t.Error("出错!")}
	if p.SearchItems["field"].VText2 != "一" {t.Error("出错!")}

	if p.SearchItems["field2"].FieldName != "field2" {t.Error("出错!")}
	if p.SearchItems["field2"].IsShow != true {t.Error("出错!")}
	if p.SearchItems["field2"].Text != "测试" {t.Error("出错!")}
	if p.SearchItems["field2"].ValueRawStr != "测试|asdf=123|喵" {t.Error("出错!")}
	if p.SearchItems["field2"].ConditionStr != "asdf=123" {t.Error("出错!")}
	if p.SearchItems["field2"].ConditionType != EQUALS {t.Error("出错!")}
	if p.SearchItems["field2"].SearchValue != "asdf=123" {t.Error("出错!")}
	if p.SearchItems["field2"].WhereStr != "field2 = ?" {t.Error("出错!")}
	if p.SearchItems["field2"].ConditionInfo != "测试 等于 喵" {t.Error("出错!")}
	if p.SearchItems["field2"].VText1 != "喵" {t.Error("出错!")}
	if p.SearchItems["field2"].Mapping[0] != "asdf=123" {t.Error("出错!")}
	if p.SearchItems["field2"].VText2 != "" {t.Error("出错!")}
}

func Test_Field(t *testing.T) {// 测试用例必须以Test开头
	// 带field的模式
    // 输入空值
	ss,sss := ParseSearchStr("","s.","st.")
	if len(ss) != 0 { t.Error("搜索结果必须为0")}
	if len(sss) != 0 { t.Error("排序结果必须为0")}
	ss,sss = ParseSearchStr("s_word.0=我了个去|123123&s_word1.1=asdfasd&&s_word1=asdfasd1&&&st_time.1&&","s_","st_")
	if len(ss) != 3 { t.Error("搜索结果必须为3")}
	if len(sss) != 1{ t.Error("排序结果必须为0")}

	s_info := BooisSearchItemInfo{}
	rs := BooisSearchStrParseResultInfo{FieldName:"word",IsShow:true,Value:"哈哈|*123123*|值1,值2"}
	s_info.Load(rs)
	if s_info.FieldName != "word" { t.Error("字段名必须为word")}
	if s_info.IsShow != true { t.Error("IsShow必须为true")}
	if s_info.Text != "哈哈" { t.Error("字段Text必须为哈哈")}
	if s_info.SearchValue != "123123" { t.Error("SearchValue必须为123123")}
	if s_info.ConditionStr != "*123123*" { t.Error("ConditionStr必须为*123123*")}
	if s_info.ValueRawStr != "哈哈|*123123*|值1,值2" { t.Error("ValueRawStr必须为 哈哈|*123123*|值1,值2")}
	if s_info.WhereStr != "word like ?" { t.Error("WhereStr必须为 word like ?")}
	if s_info.ConditionInfo != "哈哈 包含 值1" { t.Error("ConditionInfo必须为 哈哈 包含 值1")}
	if s_info.ConditionType != LIKE { t.Error("ConditionType必须为 LIKE")}
	if s_info.VText1 != "值1" { t.Error("VText1必须为 值1")}
	if s_info.VText2 != "值2" { t.Error("VText2必须为 值2")}
}

func Test_Condition(t *testing.T) {
	assertTupleEqual := func(info BooisSearchConditionParseResultInfo,
		ConditionType int,SearchValue string,WhereStr string,Mapping []interface{},Info string,
	){
		if info.ConditionType != ConditionType {
			t.Error(info.WhereStr+"检验错误:info.ConditionType:",info.ConditionType)
		}
		if info.ValueStr != SearchValue {
			t.Error(info.WhereStr+"检验错误:info.ValueStr:","|",info.ValueStr,"|",SearchValue)
		}
		if info.WhereStr != WhereStr {
			t.Error(info.WhereStr+"检验错误:info.WhereStr:","|",info.WhereStr,"|",WhereStr)
		}
		if info.Info != Info {
			t.Error(info.WhereStr+"检验错误:info.ConditionInfo:","|",info.Info,"|",Info)
		}
		if len(info.Mapping) != len(Mapping){
			t.Error(info.WhereStr+" mapping 长度错误 ",len(info.Mapping),len(Mapping))
		}
		for x:=range Mapping{
			has := false
			for j:= range info.Mapping {
				if Mapping[x].(string) == info.Mapping[j].(string) {has = true}
			}
			if !has {
				t.Error(info.WhereStr+" mapping 中没有相应的值")
			}
		}
	}
	// 有以下几种形式(用了四种判定符各自组合:!,*,~,=):
	// 1.以！开头的表示否
	// &s.field = !abc => field <> abc
	assertTupleEqual(ParseSearchCondition("!abc"),NOT_EQUALS, "abc", "%s <> ?", []interface{}{"abc"}, "%s 不等于 %s")
	// &s.field =!*abc* => field not like "%abc%"
	assertTupleEqual(ParseSearchCondition("!*abc*"),NOT_LIKE, "abc", "%s not like ?", []interface{}{"%abc%"}, "%s 不包含 %s")
	// &s.field = !*abc => field not like "%abc"
	assertTupleEqual(ParseSearchCondition("!*abc"), NOT_LIKE_END, "abc", "%s not like ?", []interface{}{"%abc"}, "%s 不以 %s 结尾")
	// &s.field = !abc* => field not like "abc%"
	assertTupleEqual(ParseSearchCondition("!abc*"),NOT_LIKE_START, "abc", "%s not like ?", []interface{}{"abc%"}, "%s 不以 %s 开头")
	// &s.field = !~1,2,3~ =>相当于not in(...),实际将转化为 field<>1 or field<>2 or filed<>3
	assertTupleEqual(ParseSearchCondition("!~1,2,3~"),NOT_IN, "1,2,3", "%s <> ? and %s <> ? and %s <> ?",[]interface{}{"1", "2", "3"},"%s 不等于 1 或 2 或 3")
	// 2.以波浪线为范围的
	// &s.field = 1~~1 => 1 between 1 用于时间
	assertTupleEqual(ParseSearchCondition("2014-1-1 1:1:1~~2014-1-1 1:1:1"),DATE_BETWEEN, "2014-1-1 1:1:1,2014-1-1 1:1:1","(%s between ? and ?)", []interface{}{"2014-1-1 1:1:1", "2014-1-1 1:1:1"}, "%s 从 %s 到 %s")
	// &s.field = ~1~ or ~1,2,3~ => 相当于in(...),实际将转化为 field=1 or field=2 or filed=3
	assertTupleEqual(ParseSearchCondition("~1~"),IN, "1",  "%s = ?", []interface{}{"1"}, "%s 等于 1")
	assertTupleEqual(ParseSearchCondition("~1,2,3~"),IN, "1,2,3",  "%s = ? or %s = ? or %s = ?", []interface{}{"1", "2", "3"},"%s 等于 1 或 2 或 3")
	// &s.field = ~1 => field<1 小于
	assertTupleEqual(ParseSearchCondition("~1"),LESS_THAN, "1", "%s < ?", []interface{}{"1"}, "%s 小于 %s")
	// assertTupleEqual(ConditionParser.parse("~a"),()
	// &s.field = ~=1 => field<=1 小于等于
	assertTupleEqual(ParseSearchCondition("~=1"),LESS_EQUALS, "1", "%s <= ?",[]interface{}{"1"}, "%s 小于等于 %s")
	// assertTupleEqual(ConditionParser.parse("~=a"),()
	// &s.field = 1~ => field>1 大于
	assertTupleEqual(ParseSearchCondition("1~"),GREATER_THAN, "1", "%s > ?", []interface{}{"1"}, "%s 大于 %s")
	// assertTupleEqual(ConditionParser.parse("a~"),()
	// &s.field = 1=~ => field>1 大于等于
	assertTupleEqual(ParseSearchCondition("1=~"),GREATER_EQUALS, "1", "%s >= ?", []interface{}{"1"}, "%s 大于等于 %s")
	// assertTupleEqual(ConditionParser.parse("a=~"),()
	// &s.field = 0~1 => field<0 and field>1
	assertTupleEqual(ParseSearchCondition("0~1"),INT_RANGE_G_L, "0,1", "%s > ? and %s < ?", []interface{}{"0", "1"},"%s 大于 %s 小于 %s")
	// assertTupleEqual(ConditionParser.parse("a~b"),()
	// &s.field = 0=~=1 => field>=0 and field<=1
	assertTupleEqual(ParseSearchCondition("0=~=1"),INT_RANGE_GE_LE, "0,1", "%s >= ? and %s <= ?", []interface{}{"0", "1"},"%s 大于等于 %s 小于等于 %s")
	// assertTupleEqual(ConditionParser.parse("a=~=b"),()
	// &s.field = 0~=1 => field>0 and field=<1
	assertTupleEqual(ParseSearchCondition("0~=1"),INT_RANGE_G_LE, "0,1", "%s > ? and %s <= ?", []interface{}{"0", "1"},"%s 大于 %s 小于等于 %s")
	// assertTupleEqual(ConditionParser.parse("a~=b"),()
	// &s.field = 0=~1 => field>=0 and field<1
	assertTupleEqual(ParseSearchCondition("0=~1"),INT_RANGE_GE_L, "0,1", "%s >= ? and %s < ?", []interface{}{"0", "1"},"%s 大于等于 %s 小于 %s")
	// assertTupleEqual(ConditionParser.parse("a=~b"),()
	// 3.带星号通配符的
	// &s.field = *abc* => field like "%abc%"
	assertTupleEqual(ParseSearchCondition("*abc*"),LIKE, "abc","%s like ?", []interface{}{"%abc%"}, "%s 包含 %s")
	// &s.field = abc* => field like "abc%"
	assertTupleEqual(ParseSearchCondition("abc*"),LIKE_START, "abc", "%s like ?", []interface{}{"abc%"}, "%s 以 %s 开头")
	// &s.field = *abc => field like "%abc"
	assertTupleEqual(ParseSearchCondition("*abc"),LIKE_END, "abc",  "%s like ?", []interface{}{"%abc"}, "%s 以 %s 结尾")
	// 4.什么都没有的
	// &s.field = abc  => field = abc
	assertTupleEqual(ParseSearchCondition("abc"),EQUALS, "abc", "%s = ?", []interface{}{"abc"}, "%s 等于 %s")
}

func Test_SearchInfo(t *testing.T) {
 		// 有以下几种形式(用了四种判定符各自组合:!,*,~,=):
        // 1.以！开头的表示否
        // region  &s.field = !abc => field <> abc
        info := BooisSearchItemInfo{}
		info.Load(BooisSearchStrParseResultInfo{FieldName:"name",IsShow:true,Value:"test title|!abc"})
	
        if info.ConditionType != NOT_EQUALS { t.Error("出错了")}
        if info.ConditionStr != "!abc" { t.Error("出错了")}
        if info.FieldName != "name" { t.Error("出错了")}
        if info.Text != "test title" { t.Error("出错了")}
        if info.IsShow!= true { t.Error("出错了")}
        if info.SearchValue!= "abc" { t.Error("出错了")}
        if info.WhereStr!= "name <> ?" { t.Error("出错了")}
        if len(info.Mapping) != 1 { t.Error("出错了")}
        if info.Mapping[0] != "abc" { t.Error("出错了")}
        if info.ConditionInfo != "test title 不等于 abc" { t.Error("出错了")}
        // endregion
        // region  &s.field =!*abc* => field not like '%abc%'
        info = BooisSearchItemInfo{}
		info.Load(BooisSearchStrParseResultInfo{FieldName:"name",IsShow:false,Value:"test 中文|!*abc*"})
        if info.ConditionType != NOT_LIKE { t.Error("出错了")}
        if info.ConditionStr != "!*abc*" { t.Error("出错了")}
        if info.FieldName != "name" { t.Error("出错了")}
        if info.Text != "test 中文" { t.Error("出错了")}
        if info.IsShow!= false { t.Error("出错了")}
        if info.SearchValue!= "abc" { t.Error("出错了")}
        if info.WhereStr!= "name not like ?" { t.Error("出错了")}
        if info.Mapping[0] != "%abc%" { t.Error("出错了")}
        if info.ConditionInfo != "test 中文 不包含 abc" { t.Error("出错了")}
        // endregion
        // region &s.field = !*abc => field not like '%abc'
        info = BooisSearchItemInfo{}
		info.Load(BooisSearchStrParseResultInfo{FieldName:"name",IsShow:false,Value:"test 中文|!*abc"})
        if info.ConditionType != NOT_LIKE_END { t.Error("出错了")}
        if info.ConditionStr != "!*abc" { t.Error("出错了")}
        if info.FieldName != "name" { t.Error("出错了")}
        if info.Text != "test 中文" { t.Error("出错了")}
        if info.IsShow!= false { t.Error("出错了")}
        if info.SearchValue!= "abc" { t.Error("出错了")}
        if info.WhereStr!= "name not like ?" { t.Error("出错了")}
        if info.Mapping[0] != "%abc" { t.Error("出错了")}
        if info.ConditionInfo != "test 中文 不以 abc 结尾" { t.Error("出错了")}
        // endregion
        // region &s.field = !abc* => field not like 'abc%'
        info = BooisSearchItemInfo{}
		info.Load(BooisSearchStrParseResultInfo{FieldName:"name",IsShow:false,Value:"test 中文|!abc*"})
        if info.ConditionType != NOT_LIKE_START { t.Error("出错了")}
        if info.ConditionStr != "!abc*" { t.Error("出错了")}
        if info.FieldName != "name" { t.Error("出错了")}
        if info.Text != "test 中文" { t.Error("出错了")}
        if info.IsShow!= false { t.Error("出错了")}
        if info.SearchValue!= "abc" { t.Error("出错了")}
        if info.WhereStr!= "name not like ?" { t.Error("出错了")}
        if info.Mapping[0] != "abc%" { t.Error("出错了")}
        if info.ConditionInfo != "test 中文 不以 abc 开头" { t.Error("出错了")}
        // endregion
        // region &s.field = !~1,2,3~ =>相当于not in(...),实际将转化为 field<>1 or field<>2 or filed<>3
        info = BooisSearchItemInfo{}
		info.Load(BooisSearchStrParseResultInfo{FieldName:"name",IsShow:false,Value:"号码|!~1,2,3~"})
        if info.ConditionType != NOT_IN { t.Error("出错了")}
        if info.ConditionStr != "!~1,2,3~" { t.Error("出错了")}
        if info.FieldName != "name" { t.Error("出错了")}
        if info.Text != "号码" { t.Error("出错了")}
        if info.IsShow!= false { t.Error("出错了")}
        if info.SearchValue!= "1,2,3" { t.Error("出错了")}
        if info.WhereStr!= "name <> ? and name <> ? and name <> ?" { t.Error("出错了")}
        if info.Mapping[0] != "1" { t.Error("出错了")}
        if info.Mapping[1] != "2" { t.Error("出错了")}
        if info.Mapping[2] != "3" { t.Error("出错了")}
        if info.ConditionInfo != "号码 不等于 1 或 2 或 3" { t.Error("出错了")}
        // endregion
        // 2.以波浪线为范围的
        // region &s.field = 1~~1 => 1 between 1 用于时间
        info = BooisSearchItemInfo{}
		info.Load(BooisSearchStrParseResultInfo{FieldName:"name",IsShow:false,Value:"号码|2015-1-1 1:1:1~~2015-1-1 1:1:2"})
        if info.ConditionType != DATE_BETWEEN { t.Error("出错了")}
        if info.ConditionStr != "2015-1-1 1:1:1~~2015-1-1 1:1:2" { t.Error("出错了")}
        if info.FieldName != "name" { t.Error("出错了")}
        if info.Text != "号码" { t.Error("出错了")}
        if info.IsShow!= false { t.Error("出错了")}
        if info.SearchValue!= "2015-1-1 1:1:1,2015-1-1 1:1:2" { t.Error("出错了")}
        if info.WhereStr!= "(name between ? and ?)" { t.Error("出错了")}
        if info.Mapping[0] != "2015-1-1 1:1:1" { t.Error("出错了")}
        if info.Mapping[1] != "2015-1-1 1:1:2" { t.Error("出错了")}
        if info.ConditionInfo != "号码 从 2015-1-1 1:1:1 到 2015-1-1 1:1:2" { t.Error("出错了")}
        // endregion
        // region &s.field = ~1~ or ~1,2,3~ => 相当于in(...),实际将转化为 field=1 or field=2 or filed=3
        info = BooisSearchItemInfo{}
		info.Load(BooisSearchStrParseResultInfo{FieldName:"name",IsShow:false,Value:"号码|~1,2,3~"})
        if info.ConditionType != IN { t.Error("出错了")}
        if info.ConditionStr != "~1,2,3~" { t.Error("出错了")}
        if info.FieldName != "name" { t.Error("出错了")}
        if info.Text != "号码" { t.Error("出错了")}
        if info.IsShow!= false { t.Error("出错了")}
        if info.SearchValue!= "1,2,3" { t.Error("出错了")}
        if info.WhereStr!= "name = ? or name = ? or name = ?" { t.Error("出错了")}
        if info.Mapping[0] != "1" { t.Error("出错了")}
        if info.Mapping[1] != "2" { t.Error("出错了")}
        if info.Mapping[2] != "3" { t.Error("出错了")}
        if info.ConditionInfo != "号码 等于 1 或 2 或 3" { t.Error("出错了")}
        // endregion
        // region &s.field = ~1 => field<1 小于
        info = BooisSearchItemInfo{}
		info.Load(BooisSearchStrParseResultInfo{FieldName:"name",IsShow:false,Value:"号码|~11"})
        if info.ConditionType != LESS_THAN { t.Error("出错了")}
        if info.ConditionStr != "~11" { t.Error("出错了")}
        if info.FieldName != "name" { t.Error("出错了")}
        if info.Text != "号码" { t.Error("出错了")}
        if info.IsShow!= false { t.Error("出错了")}
        if info.SearchValue!= "11" { t.Error("出错了")}
        if info.WhereStr!= "name < ?" { t.Error("出错了")}
        if info.Mapping[0] != "11" { t.Error("出错了")}
        if info.ConditionInfo != "号码 小于 11" { t.Error("出错了")}
        // endregion
        // self.assertTupleEqual(ConditionParser.parse("~a"),())
        // region &s.field = ~=1 => field<=1 小于等于
        info = BooisSearchItemInfo{}
		info.Load(BooisSearchStrParseResultInfo{FieldName:"name",IsShow:false,Value:"号码|~=1"})
        if info.ConditionType != LESS_EQUALS { t.Error("出错了")}
        if info.ConditionStr != "~=1" { t.Error("出错了")}
        if info.FieldName != "name" { t.Error("出错了")}
        if info.Text != "号码" { t.Error("出错了")}
        if info.IsShow!= false { t.Error("出错了")}
        if info.SearchValue!= "1" { t.Error("出错了")}
        if info.WhereStr!= "name <= ?" { t.Error("出错了")}
        if info.Mapping[0] != "1" { t.Error("出错了")}
        if info.ConditionInfo != "号码 小于等于 1" { t.Error("出错了")}
        // endregion
        // self.assertTupleEqual(ConditionParser.parse("~=a"),())
        // region &s.field = 1~ => field>1 大于
        info = BooisSearchItemInfo{}
		info.Load(BooisSearchStrParseResultInfo{FieldName:"name",IsShow:false,Value:"号码|1~"})
        if info.ConditionType != GREATER_THAN { t.Error("出错了")}
        if info.ConditionStr != "1~" { t.Error("出错了")}
        if info.FieldName != "name" { t.Error("出错了")}
        if info.Text != "号码" { t.Error("出错了")}
        if info.IsShow!= false { t.Error("出错了")}
        if info.SearchValue!= "1" { t.Error("出错了")}
        if info.WhereStr!= "name > ?" { t.Error("出错了")}
        if info.Mapping[0] != "1" { t.Error("出错了")}
        if info.ConditionInfo != "号码 大于 1" { t.Error("出错了")}
        // endregion
        // self.assertTupleEqual(ConditionParser.parse("a~"),())
        // region &s.field = 1=~ => field>1 大于等于
        info = BooisSearchItemInfo{}
		info.Load(BooisSearchStrParseResultInfo{FieldName:"name",IsShow:false,Value:"号码|1=~"})
        if info.ConditionType != GREATER_EQUALS { t.Error("出错了")}
        if info.ConditionStr != "1=~" { t.Error("出错了")}
        if info.FieldName != "name" { t.Error("出错了")}
        if info.Text != "号码" { t.Error("出错了")}
        if info.IsShow!= false { t.Error("出错了")}
        if info.SearchValue!= "1" { t.Error("出错了")}
        if info.WhereStr!= "name >= ?" { t.Error("出错了")}
        if info.Mapping[0] != "1" { t.Error("出错了")}
        if info.ConditionInfo != "号码 大于等于 1" { t.Error("出错了")}
        // endregion
        // self.assertTupleEqual(ConditionParser.parse("a=~"),())
        // region &s.field = 0~1 => field<0 and field>1
        info = BooisSearchItemInfo{}
		info.Load(BooisSearchStrParseResultInfo{FieldName:"name",IsShow:false,Value:"号码|0~1"})
        if info.ConditionType != INT_RANGE_G_L { t.Error("出错了")}
        if info.ConditionStr != "0~1" { t.Error("出错了")}
        if info.FieldName != "name" { t.Error("出错了")}
        if info.Text != "号码" { t.Error("出错了")}
        if info.IsShow!= false { t.Error("出错了")}
        if info.SearchValue!= "0,1" { t.Error("出错了")}
        if info.WhereStr!= "name > ? and name < ?" { t.Error("出错了")}
        if info.Mapping[0] != "0" { t.Error("出错了")}
        if info.Mapping[1] != "1" { t.Error("出错了")}
        if info.ConditionInfo != "号码 大于 0 小于 1" { t.Error("出错了")}
        // endregion
        // self.assertTupleEqual(ConditionParser.parse("a~b"),())
        // region &s.field = 0=~=1 => field>=0 and field<=1
        info = BooisSearchItemInfo{}
		info.Load(BooisSearchStrParseResultInfo{FieldName:"name",IsShow:false,Value:"号码|0=~=1"})
        if info.ConditionType != INT_RANGE_GE_LE { t.Error("出错了")}
        if info.ConditionStr != "0=~=1" { t.Error("出错了")}
        if info.FieldName != "name" { t.Error("出错了")}
        if info.Text != "号码" { t.Error("出错了")}
        if info.IsShow!= false { t.Error("出错了")}
        if info.SearchValue!= "0,1" { t.Error("出错了")}
        if info.WhereStr!= "name >= ? and name <= ?" { t.Error("出错了")}
        if info.Mapping[0] != "0" { t.Error("出错了")}
        if info.Mapping[1] != "1" { t.Error("出错了")}
        if info.ConditionInfo != "号码 大于等于 0 小于等于 1" { t.Error("出错了")}
        // endregion
        // self.assertTupleEqual(ConditionParser.parse("a=~=b"),())
        // region &s.field = 0~=1 => field>0 and field=<1
        info = BooisSearchItemInfo{}
		info.Load(BooisSearchStrParseResultInfo{FieldName:"name",IsShow:false,Value:"号码|0~=1"})
        if info.ConditionType != INT_RANGE_G_LE { t.Error("出错了")}
        if info.ConditionStr != "0~=1" { t.Error("出错了")}
        if info.FieldName != "name" { t.Error("出错了")}
        if info.Text != "号码" { t.Error("出错了")}
        if info.IsShow!= false { t.Error("出错了")}
        if info.SearchValue!= "0,1" { t.Error("出错了")}
        if info.WhereStr!= "name > ? and name <= ?" { t.Error("出错了")}
        if info.Mapping[0] != "0" { t.Error("出错了")}
        if info.Mapping[1] != "1" { t.Error("出错了")}
        if info.ConditionInfo != "号码 大于 0 小于等于 1" { t.Error("出错了")}
        // endregion
        // self.assertTupleEqual(ConditionParser.parse("a~=b"),())
        // region &s.field = 0=~1 => field>=0 and field<1
        info = BooisSearchItemInfo{}
		info.Load(BooisSearchStrParseResultInfo{FieldName:"name",IsShow:false,Value:"号码|0=~1"})
        if info.ConditionType != INT_RANGE_GE_L { t.Error("出错了")}
        if info.ConditionStr != "0=~1" { t.Error("出错了")}
        if info.FieldName != "name" { t.Error("出错了")}
        if info.Text != "号码" { t.Error("出错了")}
        if info.IsShow!= false { t.Error("出错了")}
        if info.SearchValue!= "0,1" { t.Error("出错了")}
        if info.WhereStr!= "name >= ? and name < ?" { t.Error("出错了")}
        if info.Mapping[0] != "0"{ t.Error("出错了")}
        if info.Mapping[1] != "1"{ t.Error("出错了")}
        if info.ConditionInfo != "号码 大于等于 0 小于 1" { t.Error("出错了")}
        // endregion
        // self.assertTupleEqual(ConditionParser.parse("a=~b"),())
        // 3.带星号通配符的
        // region &s.field = *abc* => field like "%abc%"
        info = BooisSearchItemInfo{}
		info.Load(BooisSearchStrParseResultInfo{FieldName:"name",IsShow:false,Value:"号码|*abc*"})
        if info.ConditionType != LIKE { t.Error("出错了")}
        if info.ConditionStr != "*abc*" { t.Error("出错了")}
        if info.FieldName != "name" { t.Error("出错了")}
        if info.Text != "号码" { t.Error("出错了")}
        if info.IsShow!= false { t.Error("出错了")}
        if info.SearchValue!= "abc" { t.Error("出错了")}
        if info.WhereStr!= "name like ?" { t.Error("出错了")}
        if info.Mapping[0] != "%abc%" { t.Error("出错了")}
        if info.ConditionInfo != "号码 包含 abc" { t.Error("出错了")}
        // endregion
        // region &s.field = abc* => field like "abc%"
        info = BooisSearchItemInfo{}
		info.Load(BooisSearchStrParseResultInfo{FieldName:"name",IsShow:false,Value:"号码|abc*"})
        if info.ConditionType != LIKE_START { t.Error("出错了")}
        if info.ConditionStr != "abc*" { t.Error("出错了")}
        if info.FieldName != "name" { t.Error("出错了")}
        if info.Text != "号码" { t.Error("出错了")}
        if info.IsShow!= false { t.Error("出错了")}
        if info.SearchValue!= "abc" { t.Error("出错了")}
        if info.WhereStr!= "name like ?" { t.Error("出错了")}
        if info.Mapping[0] != "abc%" { t.Error("出错了")}
        if info.ConditionInfo != "号码 以 abc 开头" { t.Error("出错了")}
        // endregion
        // region &s.field = *abc => field like "%abc"
        info = BooisSearchItemInfo{}
		info.Load(BooisSearchStrParseResultInfo{FieldName:"name",IsShow:false,Value:"号码|*abc"})
        if info.ConditionType != LIKE_END { t.Error("出错了")}
        if info.ConditionStr != "*abc" { t.Error("出错了")}
        if info.FieldName != "name" { t.Error("出错了")}
        if info.Text != "号码" { t.Error("出错了")}
        if info.IsShow!= false { t.Error("出错了")}
        if info.SearchValue!= "abc" { t.Error("出错了")}
        if info.WhereStr!= "name like ?" { t.Error("出错了")}
        if info.Mapping[0] != "%abc"{ t.Error("出错了")}
        if info.ConditionInfo != "号码 以 abc 结尾" { t.Error("出错了")}
        // endregion
        // 4.什么都没有的
        // region &s.field = abc  => field = abc
        info = BooisSearchItemInfo{}
		info.Load(BooisSearchStrParseResultInfo{FieldName:"name",IsShow:false,Value:"号码|abc"})
        if info.ConditionType != EQUALS { t.Error("出错了")}
        if info.ConditionStr != "abc" { t.Error("出错了")}
        if info.FieldName != "name" { t.Error("出错了")}
        if info.Text != "号码" { t.Error("出错了")}
        if info.IsShow!= false { t.Error("出错了")}
        if info.SearchValue!= "abc" { t.Error("出错了")}
        if info.WhereStr!= "name = ?" { t.Error("出错了")}
        if info.Mapping[0] != "abc" { t.Error("出错了")}
        if info.ConditionInfo != "号码 等于 abc" { t.Error("出错了")}
        // endregion
}
