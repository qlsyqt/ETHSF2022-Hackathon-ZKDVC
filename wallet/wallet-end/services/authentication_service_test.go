package services

import (
	"database/sql"
	"fmt"
	"testing"
)

func cleanTablesForDbg() {
	mysql1 := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		"root", "12345678", "127.0.0.1", 3306, "hackthon")
	mysql2 := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		"root", "12345678", "127.0.0.1", 3306, "hackthonw")
	mysqlDb1, _ := sql.Open("mysql", mysql1)
	mysqlDb2, _ := sql.Open("mysql", mysql2)

	_, err := mysqlDb1.Exec("drop table if exists t_did ;")
	if err != nil {
		panic(err)
	}
	_, err = mysqlDb1.Exec("drop table if exists t_claim;")
	if err != nil {
		panic(err)
	}
	_, err = mysqlDb1.Exec("drop table if exists offer;")
	if err != nil {
		panic(err)
	}
	_, err = mysqlDb1.Exec("drop table if exists template;")
	if err != nil {
		panic(err)
	}

	_, err = mysqlDb2.Exec("drop table if exists t_did ;")
	if err != nil {
		panic(err)
	}
	_, err = mysqlDb2.Exec("drop table if exists t_claim;")
	if err != nil {
		panic(err)
	}
	_, err = mysqlDb2.Exec("drop table if exists offer;")
	if err != nil {
		panic(err)
	}
	_, err = mysqlDb2.Exec("drop table if exists template;")
	if err != nil {
		panic(err)
	}

	pg1 := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"127.0.0.1", 5432, "root", "12345678", "hackthon")
	pg2 := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"127.0.0.1", 5432, "root", "12345678", "hackthonw")
	pgdb1, _ := sql.Open("postgres", pg1)
	pgdb2, _ := sql.Open("postgres", pg2)

	_, err = pgdb1.Exec("drop table if exists mt_nodes;")
	if err != nil {
		panic(err)
	}
	_, err = pgdb1.Exec("drop table if exists mt_roots;")
	if err != nil {
		panic(err)
	}
	_, err = pgdb2.Exec("drop table if exists mt_nodes;")
	if err != nil {
		panic(err)
	}
	_, err = pgdb2.Exec("drop table if exists mt_roots;")
	if err != nil {
		panic(err)
	}
	fmt.Println("本地数据库清理完毕")
}

func TestAuthentication(t *testing.T) {
	cleanTablesForDbg()

}
