package cmd_parser

import (
	"strings"
)



func Cmd_parse(cmd_str string)(cmd string,args map[string]string){

	if !strings.Contains(cmd_str,"-"){
		if !strings.HasPrefix(cmd_str," "){
			return cmd_str,	map[string]string{}
		}else{
			cmd_arr:=strings.Split(cmd_str," ")
			return cmd_arr[0],map[string]string{
				cmd_arr[0]:cmd_arr[1],
			}
		}
	}
	cmd_store:=map[string]string{}
	last_char:=""
	is_start:=false
	is_cmd_start:=false
	field:=[]string{}
	cmd_bag:=[]string{}
	cmd_bag_is_open:=false
	val_bag:=[]string{}
	val_bg_is_open:=false
	qout_start:=false
	is_cmd_char:=false
	str_len:=len([]rune(cmd_str))

	cmd_str_rune:=[]rune(cmd_str)
	for i:=0;i<str_len;i++ {
		str_in_i:=string(cmd_str_rune[i])
		is_cmd_char=(last_char==" " || i==0) && str_in_i=="-" && !qout_start
		if qout_start{
			if last_char!="\\" && str_in_i=="\"" {
				qout_start = false
			}

		}else if last_char!="\\" && str_in_i=="\""{
			qout_start=true
		}
		if !is_start && string(cmd_str_rune)!=" "{
			is_start=true
		}
		if is_start && !is_cmd_start && str_in_i!=" " && !is_cmd_char{
			field=append(field,str_in_i)
		}
		if !is_cmd_start && is_cmd_char{
			is_cmd_start=true
		}
		if !cmd_bag_is_open && val_bg_is_open{
			if (str_in_i!=" " || (str_in_i==" " && qout_start)) && !is_cmd_char && str_in_i!="\""{
				val_bag=append(val_bag,str_in_i)
			}else if(str_in_i=="\"" && last_char=="\\"){
				val_bag=append(val_bag,str_in_i)
			}else if (!qout_start){
				val_bg_is_open=false
			}
		}
		if cmd_bag_is_open{
			if(str_in_i!=" " && !is_cmd_char){
				cmd_bag=append(cmd_bag,str_in_i)
			}else{
				cmd_bag_is_open=false
			}
		}

		if(is_cmd_char){
			if(len(cmd_bag)>0){
				cmd_store[strings.Join(cmd_bag,"")]=strings.Join(val_bag,"")

			}
			cmd_bag=[]string{}
			cmd_bag_is_open=true
			val_bag=[]string{}
			val_bg_is_open=true
		}
		if (i==str_len-1){
			if(len(cmd_bag)>0){
				cmd_store[strings.Join(cmd_bag,"")]=strings.Join(val_bag,"")
			}
			cmd_bag=[]string{}
			cmd_bag_is_open=false
			val_bag=[]string{}
			val_bg_is_open=false
		}
		last_char=str_in_i

	}
	if qout_start{
		//
	}
	return strings.Join(field,""),cmd_store
}
