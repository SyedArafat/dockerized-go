package database

import (
	"database/sql"
	"os"
	"strconv"
	"time"
)

// DB holds the database connection pool
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

var maxOpenDbConn, _ = strconv.Atoi(os.Getenv("MAX_OPEN_DB_CONNECTIONS"))
var maxIdleDbConn, _ = strconv.Atoi(os.Getenv("MAX_IDLE_DB_CONNECTIONS"))
var dbLifeTimeInMinute, _ = strconv.Atoi(os.Getenv("MAX_DB_LIFETIME_IN_MINUTE"))
var maxDbLifeTime = time.Duration(dbLifeTimeInMinute) * time.Minute

func ConnectSQL() (*DB, error) {
	d, err := NewDatabase(dsn())
	if err != nil {
		panic(err)
	}

	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifeTime)

	dbConn.SQL = d
	err = testDB(dbConn.SQL)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

func testDB(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}

	return nil
}

func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
