package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"wallet-end/config"
)

const MYSQL_NAME = "mysql"

type MySqlStub struct {
}

func NewMySqlDb() *MySqlStub {
	return &MySqlStub{}
}

func (db *MySqlStub) GetConn() (*sql.DB, error) {
	datasource := db.GetDsn()
	conn, err := sql.Open("mysql", datasource)
	return conn, err
}

func (db *MySqlStub) GetDsn() string {
	cfg := config.GetConfig()
	dbConfig := cfg.Database[MYSQL_NAME]
	datasource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		dbConfig.DbUser, dbConfig.DbPwd, dbConfig.Server, dbConfig.Port, dbConfig.Db)
	return datasource
}
