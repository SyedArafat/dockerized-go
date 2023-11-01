package main

import (
	"database/sql"
	"export_nrt_report/internal/actions"
	"export_nrt_report/internal/database"
	"export_nrt_report/internal/repository"
	"fmt"
)

func main() {
	dbConn := initialize()
	defer func(dbConn *sql.DB) {
		err := dbConn.Close()
		if err != nil {

		}
	}(dbConn)
	fmt.Println("Project Initialisation Successful")
	actions.GetCommissionRecords()
}

func initialize() *sql.DB {
	conn, _ := database.ConnectSQL()
	repository.SetDatabase(conn.SQL)

	return conn.SQL
}
