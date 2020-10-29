package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	MysqlDB *sql.DB
)

func InitMysql() error {
	MysqlDB, err := sql.Open("mysql", MySQLConnectURI)
	if err != nil {
		return err
	}

	MysqlDB.SetMaxOpenConns(MySQLMaxOpenConns)
	MysqlDB.SetMaxIdleConns(MySQLMaxIdleConns)

	err = MysqlDB.Ping()
	if err != nil {
		return err
	}

	return nil
}
