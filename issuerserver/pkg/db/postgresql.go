package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"issuerserver/common"
	"issuerserver/config"
	"reflect"
	"sync"
)

const PG_NAME = "pg"

var once sync.Once = sync.Once{}

// It implements DB interface in https://github.com/iden3/go-merkletree-sql/blob/v1.0.2/db/sql/sql.go
type PostgreSqlStub struct {
}

var pgDb *PostgreSqlStub = &PostgreSqlStub{}

func GetPgDb() *PostgreSqlStub {
	return pgDb
}

func (db *PostgreSqlStub) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	db.ensureTables()
	conn, err := db.GetConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return conn.ExecContext(ctx, query, args...)
}

// Query a single row. By database/sql standard. Return sql.ErrorNoRows if no row found.
func (db *PostgreSqlStub) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	db.ensureTables()
	conn, err := db.GetConn()
	if err != nil {
		return err
	}
	defer conn.Close()
	val := reflect.ValueOf(dest).Elem() //dereference pointer
	numCols := val.NumField()
	columns := make([]interface{}, numCols)
	for i := 0; i < numCols; i++ {
		field := val.Field(i)
		columns[i] = field.Addr().Interface()
	}

	row := conn.QueryRowContext(ctx, query, args...)
	if row.Err() != nil {
		return row.Err()
	}
	err = row.Scan(columns...) //would return sql.SqlNoRows if not exists
	if err != nil {
		return err
	}
	return nil
}

// Select multiple rows.
func (db *PostgreSqlStub) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	db.ensureTables()
	conn, err := db.GetConn()
	if err != nil {
		return err
	}
	defer conn.Close()
	rows, err := conn.QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}

	destv := reflect.ValueOf(dest).Elem()
	elementType := destv.Type().Elem()
	fields := make([]interface{}, elementType.NumField())

	for rows.Next() {
		rowp := reflect.New(elementType) // create an object on heap
		rowv := rowp.Elem()              //Elem() dereference poointer

		for i := 0; i < rowv.NumField(); i++ {
			fields[i] = rowv.Field(i).Addr().Interface()
		}

		if err = rows.Scan(fields...); err != nil {
			return err
		}

		destv.Set(reflect.Append(destv, rowv))
	}

	return nil
}

func (db *PostgreSqlStub) ensureTables() {
	once.Do(func() {
		conn, err := db.GetConn()
		if err != nil {
			panic(err)
		}
		_, err = conn.Exec(common.CREATE_MTROOTS_TABLE)
		if err != nil {
			panic(err)
		}
		fmt.Println("create mt_roots existed")
		_, err = conn.Exec(common.CREATE_MTNODES_TABLE)
		if err != nil {
			panic(err)
		}
		fmt.Println("create mt_nodes existed")
	})
}

// TODO: use Connection pool
func (db *PostgreSqlStub) GetConn() (*sql.DB, error) {
	cfg := config.GetConfig()
	dbConfig := cfg.Database[PG_NAME]
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Server, dbConfig.Port, dbConfig.DbUser, dbConfig.DbPwd, dbConfig.Db)
	conn, err := sql.Open("postgres", datasource)
	return conn, err
}
