package db_helper

import "database/sql"

func GetItems(rows * sql.Rows) [](map[string]string){
	items := [](map[string]string){}
	cols, _ := rows.Columns()
	scan_args:=make([]interface{}, len(cols))
	row_vals:=make([]string, len(cols))
	for i:=range row_vals{
		scan_args[i]=&row_vals[i]
	}
	for rows.Next() {
		rows.Scan(scan_args...)
		record:=map[string]string{}
		for i,val:=range row_vals{
			record[cols[i]]=string(val)
		}
		//fmt.Println(record)
		items=append(items,record)
	}
	return items
}
