package mysql_helper

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
func ExecuteScalar(cmdtxt string, mapping []interface{}, db_info string) (map[string]sql.RawBytes, error) {
	mysql_db, err := sql.Open("mysql", db_info)
	defer mysql_db.Close()

	if err != nil {
		return map[string]sql.RawBytes{}, err
	}
	rows, err1 := mysql_db.Query(cmdtxt, mapping...)
	if err1 != nil {
		return map[string]sql.RawBytes{}, err1
	}
	defer rows.Close()
	item := map[string]sql.RawBytes{}
	eachRow(rows, func(record map[string]sql.RawBytes) bool {
		item = record
		return false//停止遍历
	})
	return item, nil

}
func Each(cmdtxt string, mapping []interface{}, each_fn func(map[string]sql.RawBytes), db_info string) error {
	mysql_db, err := sql.Open("mysql", db_info)
	defer mysql_db.Close()

	if err != nil {
		return err
	}
	rows, err1 := mysql_db.Query(cmdtxt, mapping...)
	if err1 != nil {
		return err1
	}
	defer rows.Close()
	eachRow(rows, func(record map[string]sql.RawBytes) bool {
		each_fn(record)
		return true
	})

	return nil

}
func GetList(cmdtxt string, mapping []interface{}, db_info string) ([]map[string]sql.RawBytes, error) {
	mysql_db, err := sql.Open("mysql", db_info)
	defer mysql_db.Close()

	if err != nil {
		return []map[string]sql.RawBytes{}, err
	}
	rows, err1 := mysql_db.Query(cmdtxt, mapping...)
	if err1 != nil {
		return []map[string]sql.RawBytes{}, err1
	}
	defer rows.Close()
	items := []map[string]sql.RawBytes{}
	eachRow(rows, func(record map[string]sql.RawBytes) bool {
		items = append(items, record)
		return true
	})
	return items, nil
}

func Paging(tab_name string, fields string, where_str string, mapping []interface{}, sort_str string, page_size int, current_page int, each_fn func(map[string]sql.RawBytes), counter bool, db_info string) (int, error) {
	rs_count := 0
	if counter {
		cnt, err := Count(tab_name, where_str, mapping, db_info)
		if err != nil {
			return 0, err
		}
		rs_count = cnt
	}
	if where_str != "" {
		where_str = " where "+where_str
	}
	if sort_str != "" {
		sort_str = " order by "+sort_str
	}
	cmd := fmt.Sprintf("select %s from %s %s %s limit %d offset %d", fields, tab_name, where_str, sort_str, page_size, (current_page - 1) * page_size)

	err := Each(cmd, mapping, each_fn, db_info)
	if err != nil {
		return 0, err
	}
	return rs_count, nil

}
func Count(tab_name string, where_str string, mapping []interface{}, db_info string) (int, error) {
	rs_count := 0
	if where_str != "" {
		where_str = " where "+where_str
	}
	one, err := ExecuteScalar(fmt.Sprintf("select count(*) as num from %s%s",tab_name ,where_str) , mapping, db_info)
	if err != nil {
		return 0, err
	}
	rs_count, _ = strconv.Atoi(string(one["num"]))
	return rs_count, nil
}

func eachRow(rows *sql.Rows, each_fn func(record map[string]sql.RawBytes) bool) {
	cols, _ := rows.Columns()
	col_size := len(cols)
	for rows.Next() {
		//数据容器开始 务必把这段代码放在for之中  数据容器不可复用
		scan_args := make([]interface{}, col_size)
		row_vals := make([]sql.RawBytes, col_size)
		for i := range row_vals {
			scan_args[i] = &row_vals[i]
		}
		//数据容器结束

		rows.Scan(scan_args...)
		record := map[string]sql.RawBytes{}
		for i, val := range row_vals {
			record[cols[i]] = []byte(val)
		}
		is_go_on := each_fn(record)
		//each_fn return false 的话就停止遍历
		if !is_go_on {
			break
		}
	}
}

