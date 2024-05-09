package config

import (
	"encoding/json"
	"os"

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

func Init() error {
	err := json.Unmarshal([]byte(env), &appConfig)
	if err != nil {
		return err
	}

	return nil
}

func GetMySQLDBUsername() string {
	username := os.Getenv("MYSQL_USERNAME")
	if username == "" {
		return appConfig["MYSQL_USERNAME"]
	}
	return username
}
func GetMySQLDBPassword() string {
	pwd := os.Getenv("MYSQL_PASSWORD")
	if pwd == "" {
		return appConfig["MYSQL_PASSWORD"]
	}
	return pwd
}

func GetMySQLDBName() string {
	dbname := os.Getenv("MYSQL_DBNAME")
	if dbname == "" {
		return appConfig["MYSQL_DBNAME"]
	}
	return dbname
}
func GetMySQLDBProtocol() string {
	protocol := os.Getenv("MYSQL_PROTOCOL")
	if protocol == "" {
		return appConfig["MYSQL_PROTOCOL"]
	}
	return protocol
}
func GetMySQLDBHost() string {
	user := os.Getenv("MYSQL_HOST")
	if user == "" {
		return appConfig["MYSQL_HOST"]
	}
	return user
}
func GetMySQLDBPort() string {
	port := os.Getenv("MYSQL_PORT")
	if port == "" {
		return appConfig["MYSQL_PORT"]
	}
	return port
}
