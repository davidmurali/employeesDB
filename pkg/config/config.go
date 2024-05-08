package config

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var env = `{
    "MYSQL_USERNAME": "root",
    "MYSQL_PASSWORD": "admin",
    "MYSQL_DBNAME": "employees",
    "MYSQL_PROTOCOL": "tcp",
    "MYSQL_HOST": "localhost",
    "MYSQL_PORT": "3306"
}`

var appConfig map[string]string
var db *sql.DB

func Init() error {
	err := json.Unmarshal([]byte(env), &appConfig)
	if err != nil {
		return err
	}

	return nil
}

func GetDB() (db1 *sql.DB, err error) {
	if db == nil {
		mysqlCredentials := fmt.Sprintf(
			"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			appConfig["MYSQL_USERNAME"],
			appConfig["MYSQL_PASSWORD"],
			appConfig["MYSQL_PROTOCOL"],
			appConfig["MYSQL_HOST"],
			appConfig["MYSQL_PORT"],
			appConfig["MYSQL_DBNAME"],
		)
		db, err = sql.Open("mysql", mysqlCredentials)
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}

// func SelectDatabase(db *sql.DB, dbName string) error {
// 	query := fmt.Sprintf("USE %s", dbName)
// 	_, err := db.Exec(query)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
