package search_str_parser

import (
	"testing"
	"fmt"
)

func TestSearchStrParse(t *testing.T) {
	//SearchStrParse("s.word.0=123123&s.word1.1=asdfasd&&s.word1=asdfasd&&&&&")
	//SearchStrParse("?s.word.0=123123&s.word1.1=asdfasd&&s.word1=asdfasd&&&&&")
	res := SearchStrParse("&s.word.0=123123&s.word1.1=asdfasd&&s.word1=asdfasd&&&&&")
	fmt.Println(res)
	s:="abc上官def"
	fmt.Println(s[1:])
	fmt.Println(s[1:len(s)-1])
	//fmt.Println(s[1:-1])
}
func TestSearchAdapter_Init(t *testing.T) {
	sa:=SearchAdapter{}
	sa.Init("s.firm_id=9pin&s.page=1&st.id=1&s.create_date=2016-11-12 11:12:13~~2017-11-12 11:12:13",[]string{"firm_id","id","create_date"})
	fmt.Println("mapping",sa.Mapping)
	fmt.Println("whereStr",sa.WhereStr)
	fmt.Println("whereStrlist",sa.WhereStrList)
	fmt.Println("sorts type:",sa.Sorts[0].SortType)
	fmt.Println("sorts field:",sa.Sorts[0].Field)
}
