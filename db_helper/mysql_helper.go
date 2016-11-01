package db_helper

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"fmt"
)

func ExecuteNonQuery(cmdtxt string, mapping []interface{}, db_info string) (sql.Result, error) {
	mysql_db, err := sql.Open("mysql", db_info)
	defer mysql_db.Close()

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
	defer mysql_db.Close()

	if err != nil {
		return map[string]string{}, err
	}
	rows, err1 := mysql_db.Query(cmdtxt, mapping...)
	if err1 != nil {
		return map[string]string{}, err1
	}
	items := GetItems(rows)
	if (len(items) == 0) {
		return map[string]string{}, nil
	}
	return items[0], nil

}
func Each(cmdtxt string, mapping []interface{}, each_fn func(map[string]string), db_info string) error {
	mysql_db, err := sql.Open("mysql", db_info)
	defer mysql_db.Close()

	if err != nil {
		return err
	}
	rows, err1 := mysql_db.Query(cmdtxt, mapping...)
	if err1 != nil {
		return err1
	}
	//
	cols, _ := rows.Columns()
	scan_args := make([]interface{}, len(cols))
	row_vals := make([]string, len(cols))
	for i := range row_vals {
		scan_args[i] = &row_vals[i]
	}
	for rows.Next() {
		rows.Scan(scan_args...)
		record := map[string]string{}
		for i, val := range row_vals {
			record[cols[i]] = string(val)
		}
		//fmt.Println(record)
		each_fn(record)
	}
	return nil

}
func GetList(cmdtxt string, mapping []interface{}, db_info string) ([]map[string]string, error) {
	mysql_db, err := sql.Open("mysql", db_info)
	defer mysql_db.Close()

	if err != nil {
		return []map[string]string{}, err
	}
	rows, err1 := mysql_db.Query(cmdtxt, mapping...)
	if err1 != nil {
		return []map[string]string{}, err1
	}
	items := GetItems(rows)
	return items, nil
}

func Paging(tab_name string, fields string, where_str string, mapping []interface{}, sort_str string, page_size int, current_page int, each_fn func(map[string]string), counter bool, db_info string) (int, error) {
	rs_count := 0
	if counter {

		rs_count, err1 := Count(tab_name, where_str, mapping, db_info)
		if err1 != nil {
			_=rs_count

			return 0, err1
		}

	}
	//
	err3 := Each(fmt.Sprintf("select %s from %s where %s order by %s limit %d offset %d", fields, tab_name, where_str, sort_str, page_size, (current_page - 1) * page_size), mapping, each_fn, db_info)
	if err3 != nil {
		return 0, err3
	}
	return rs_count, nil

}
func Count(tab_name string, where_str string, mapping []interface{}, db_info string) (int, error) {
	rs_count := 0
	one, err1 := ExecuteScalar("select count(*) as num from " + tab_name + " where " + where_str, mapping, db_info)
	if err1 != nil {
		return 0, err1
	}
	rs_count, _ = strconv.Atoi(one["num"])
	return rs_count, nil
}

