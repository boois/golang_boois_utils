package boois_utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ExecuteNonQuery(cmdtxt string, mapping []interface{}, db_info string) (sql.Result, error) {
	mysql_db, err := sql.Open("mysql", db_info)
	if err != nil {
		return nil, err
	}
	ret, err1 := mysql_db.Exec(cmdtxt, mapping...)
	if err1 != nil {
		return nil, err1
	}
	return ret, nil

}
func ExecuteScalar(cmdtxt string, mapping []interface{}, db_info string) (map[string]string, error) {
	mysql_db, err := sql.Open("mysql", db_info)
	if err != nil {
		return map[string]string{}, err
	}
	rows, err1 := mysql_db.Query(cmdtxt, mapping...)
	if err != nil {
		return map[string]string{}, err1
	}
	items := get_items(rows)
	if (len(items) == 0) {
		return map[string]string{}, nil
	}
	return items[0], nil

}
func Each(cmdtxt string,each_fn func(map[string]string), mapping []interface{}, db_info string) error {
	mysql_db, err := sql.Open("mysql", db_info)
	if err != nil {
		return map[string]string{}, err
	}
	rows, err1 := mysql_db.Query(cmdtxt, mapping...)
	if err != nil {
		return map[string]string{}, err1
	}
	//
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
		each_fn(record)
	}

}
func GetList(cmdtxt string, mapping []interface{}, db_info string) ([]map[string]string, error) {
	mysql_db, err := sql.Open("mysql", db_info)
	if err != nil {
		return map[string]string{}, err
	}
	rows, err1 := mysql_db.Query(cmdtxt, mapping...)
	if err != nil {
		return map[string]string{}, err1
	}
	items := get_items(rows)
	return items, nil
}