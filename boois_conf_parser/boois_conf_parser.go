package boois_conf_parser

import (
	"strings"
)

func Parse(txt string) map[string]map[string]string {

	res:=map[string]map[string]string{}
	lines:=strings.Split(txt,"\n")
	currrent_section:="main"
	for _,line :=range lines{
		//line_trim:=strings.Trim(line," ")
		line_trim:=strings.TrimSpace(line)
		p("行",line_trim)
		if len(line_trim)==0{
			continue
		}
		if strings.HasPrefix(line_trim,"#"){
			continue
		}
	//fmt.Println("line",line_trim)

		if strings.HasPrefix(line_trim,"[") && strings.HasSuffix(line_trim,"]"){
			currrent_section=line_trim[1:len(line_trim)-1]
		}else{
			i:=strings.Index(line_trim,"=")
			if i==-1{
				p("没有等号",line_trim)
				continue
			}
			key:=strings.TrimSpace(line_trim[:i])
			p("key",key)
			val:=strings.TrimSpace(line_trim[i+1:])
			p("val",val)
			//初始化
			if res[currrent_section]==nil || len(res[currrent_section])==0{
				res[currrent_section]=map[string]string{}
			}
			res[currrent_section][key]=val
		}
	}
	return res
}
func p(v ...interface{})  {

	//fmt.Println(v...)
}
