package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/mysql"
	"time"
	"zeroagencytest/pkg/config"
	"zeroagencytest/pkg/utils/logging"
)

func New() (*reform.DB, error) {
	logger := logging.GetLogger()
	host := config.GetString("MYSQL_HOST")
	port := config.GetString("MYSQL_PORT")
	username := config.GetString("MYSQL_USER")
	pass := config.GetString("MYSQL_PASSWORD")
	dbName := config.GetString("MYSQL_DATABASE")
	logger.Infoln(host, port, username)
	sqlDB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?clientFoundRows=true&multiStatements=true", username, pass, host, port, dbName))
	if err != nil {
		logger.Fatal("SQL Connect", err)
		return nil, err
	}
	//if err = sqlDB.Ping(); err != nil {
	//	logger.Fatal("SQL Ping ", err)
	//	return nil, err
	//}
	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)
	db := reform.NewDB(sqlDB, mysql.Dialect, reform.NewPrintfLogger(logger.Infof))
	return db, nil
}
