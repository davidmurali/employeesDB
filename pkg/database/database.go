package database

import (
	"context"
	"database/sql"
	"employeesDB/pkg/constants"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

var databases map[string]*sql.DB
var databaseMutex = sync.Mutex{}

func init() {
	databases = make(map[string]*sql.DB)
}

func openDBConnection(dbConnectStr string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbConnectStr+"?parseTime=true&multiStatements=true&timeout=30s")
	if err != nil {
		logrus.Error(constants.DBOpenError + ": " + err.Error())
		return nil, err
	}
	db.SetConnMaxLifetime(1 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(7)
	db.SetConnMaxIdleTime(10 * time.Second)
	return db, err
}

func GetDb(dbUser string, dbPass string, dbProtocol string, dbHost string, dbPort int, dbName string) (db *sql.DB, err error) {
	dbConnectStr := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", dbUser, dbPass, dbProtocol, dbHost, dbPort, dbName)

	databaseMutex.Lock()
	defer databaseMutex.Unlock()

	db, exists := databases[dbConnectStr]
	if !exists {
		db, err = openDBConnection(dbConnectStr)
		if err != nil {
			return nil, err
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		logrus.Error(constants.DBConnectError + ": " + err.Error())
		db.Close()
		delete(databases, dbConnectStr)
		return nil, err
	}
	databases[dbConnectStr] = db
	return db, nil
}

func CloseDB(db *sql.DB) {
	db.Close()
	for k, v := range databases {
		if v == db {
			delete(databases, k)
			return
		}
	}
}
