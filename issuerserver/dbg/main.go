package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

//DO NOT USE THIS

func main() {
	cleanTablesForDbg()
}

func cleanTablesForDbg() {
	mysql1 := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		"root", "WMmpmHawul", "tidb.uunmqwe9i4u.clusters.tidb-cloud.com", 4000, "polygonid")
	mysql2 := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		"root", "WMmpmHawul", "tidb.uunmqwe9i4u.clusters.tidb-cloud.com", 4000, "hackthonw")

	//mysql1 := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
	//	"root", "12345678", "127.0.0.1", 3306, "hackthon")
	//mysql2 := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
	//	"root", "12345678", "127.0.0.1", 3306, "hackthonw")

	mysqlDb1, err := sql.Open("mysql", mysql1)
	if err != nil {
		panic(err)
	}
	mysqlDb2, err := sql.Open("mysql", mysql2)
	if err != nil {
		panic(err)
	}
	//
	_, err = mysqlDb1.Exec("drop table if exists t_did ;")
	if err != nil {
		panic(err)
	}
	_, err = mysqlDb1.Exec("drop table if exists t_claim;")
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

	pg1 := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"127.0.0.1", 5432, "centos", "q5c2lbobzl", "hackthons")
	pg2 := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"127.0.0.1", 5432, "centos", "q5c2lbobzl", "hackthonw")
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
