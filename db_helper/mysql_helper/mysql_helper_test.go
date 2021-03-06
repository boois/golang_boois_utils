package mysql_helper

import (
	"testing"
	"fmt"
	"time"
	"database/sql"
	"strconv"
)

var db_info = "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&loc=Local"

func TestExecuteNonQuery(t *testing.T) {
	ret, err := ExecuteNonQuery(`INSERT INTO visit (ip, user_key, create_date, url, code,price,num)VALUES (?, ?, ?, ?, ?,?,?);`, []interface{}{"1.2.3.4", "aaa", time.Now(), "url", "231231", 13.2, 3}, db_info)

	fmt.Println(err)
	fmt.Println(ret.RowsAffected())

}
func BenchmarkExecuteNonQuery(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ret, err := ExecuteNonQuery(`INSERT INTO visit (ip, user_key, create_date, url, code,price,num)VALUES (?, ?, ?, ?, ?,?,?);`, []interface{}{"1.2.3.4", "aaa", time.Now(), "url", "231231", 13.2, 3}, db_info)

		fmt.Println(err)
		fmt.Println(ret.RowsAffected())
	}
}
func TestExecuteScalar(t *testing.T) {
	one, err := ExecuteScalar("select * from visit where id=?", []interface{}{1}, db_info)
	fmt.Println(err)
	fmt.Println(one)
	fmt.Println(one["id"])
	fmt.Println(one["create_date"])
}
func BenchmarkExecuteScalar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		one, err := ExecuteScalar("select * from visit where id=?", []interface{}{1}, db_info)
		fmt.Println(err)
		fmt.Println(one)
		fmt.Println(one["id"])
		fmt.Println(one["create_date"])
	}
}

func TestBind(t *testing.T) {
	one, err := ExecuteScalar("select ? from ? where id=?", []interface{}{1, "*", "visit"}, db_info)
	fmt.Println(err)
	fmt.Println(one)
	fmt.Println(one["id"])
	fmt.Println(one["create_date"])
}
func TestEach(t *testing.T) {
	each_func := func(row map[string]sql.RawBytes) {
		fmt.Println("-------row-----")
		fmt.Println(row)

	}
	err := Each("select * from visit", []interface{}{}, each_func, db_info)
	fmt.Println(err)
}
func TestGetList(t *testing.T) {
	rows, err := GetList("select * from visit ", []interface{}{}, db_info)
	fmt.Println(err)
	fmt.Println(rows)
}
func BenchmarkGetList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rows, err := GetList("select * from visit ", []interface{}{}, db_info)
		fmt.Println(err)
		fmt.Println(rows)
	}
}
func TestPaging(t *testing.T) {
	each_func := func(row map[string]sql.RawBytes) {
		fmt.Println("-------row-----")
		fmt.Println(strconv.ParseFloat(string(row["price"]), 64))

	}
	rs_count, err := Paging("visit", "id,ip,code,price", "user_key=?", []interface{}{"aaa"}, "id desc", 2, 1, each_func, true, db_info)
	fmt.Println(rs_count)
	fmt.Println(err)
}
func BenchmarkPaging(b *testing.B) {
	for i := 0; i < b.N; i++ {
		each_func := func(row map[string]sql.RawBytes) {
			fmt.Println("-------row-----")
			fmt.Println(strconv.ParseFloat(string(row["price"]), 64))

		}
		rs_count, err := Paging("visit", "id,ip,code,price", "user_key=?", []interface{}{"aaa"}, "id desc", 2, 1, each_func, true, db_info)
		fmt.Println(rs_count)
		fmt.Println(err)
	}
}
func TestCount(t *testing.T) {
	rs_count, err := Count("visit", "user_key=?", []interface{}{"aaa"}, db_info)
	fmt.Println(rs_count)
	fmt.Println(err)
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//use b.N for looping
		rs_count, err := Count("visit", "user_key=?", []interface{}{"aaa"}, db_info)
		fmt.Println(rs_count)
		fmt.Println(err)
	}
}
