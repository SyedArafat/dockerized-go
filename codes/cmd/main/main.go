package main

import (
	"database/sql"
	"export_nrt_report/internal/controller"
	"export_nrt_report/internal/database"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	dbConn := initialize()
	defer dbConn.Close()
	a := "Initialisation Successful"
	fmt.Println(a)
	controller.GetCommissionRecords()
}

func initialize() *sql.DB {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}
	conn, _ := database.ConnectSQL()

	controller.SetDatabase(conn.SQL)

	return conn.SQL
}
