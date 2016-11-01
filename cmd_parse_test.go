package boois_utils

import (
	"testing"
)

func TestCmd_parse(t *testing.T) {
//r:=[]rune("上官")
	//fmt.Print(string(r[0]))
	//field,maps:=Cmd_parse("python -h")
	//field,maps:=Cmd_parse(`-i test%s123123 -n -l 1,4 -r "/mo\" -bi/gi" -t "-int|+float"`)
	//field,maps:=Cmd_parse(`python -i love u u u -name u`)
	field,maps:=Cmd_parse(` field -i test%s123123 -n -l 1,4 -r "/mo\" -bi/gi" -t "-int|+float"`)
	t.Log(field,maps)
}
