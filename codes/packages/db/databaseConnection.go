package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

import "export_nrt_report/utils"

type Connection struct {
	Username string
	Password string
	Hostname string
	Dbname   string
	Port     string
}

func getConnectionString() Connection {
	return Connection{
		Username: utils.Config("DB_USER"),
		Password: utils.Config("DB_PASSWORD"),
		Hostname: utils.Config("DB_HOST"),
		Dbname:   utils.Config("DB_NAME"),
		Port:     utils.Config("DB_HOST_PORT"),
	}
}

func dsn() string {
	connection := getConnectionString()
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", connection.Username, connection.Password, connection.Hostname, connection.Port, connection.Dbname)
}

func MySqlConnection() *sql.DB {
	db, err := sql.Open("mysql", dsn())
	if err != nil {
		log.Fatalf("Errors in SQL Open %s\n\n", err)
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	err = db.Ping()
	if err != nil {
		log.Fatalf("Errors %s when pinging database\n\n", err)
	}

	log.Printf("Connected to DB %s successfully\n", "dbname")

	return db
}
