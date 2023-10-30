package main

import (
	"export_nrt_report/packages/db"
	"fmt"
)

func main() {
	fmt.Println("Hello Arafat")
	initialize()
}

func initialize() {
	conn := db.MySqlConnection()
	defer conn.Close()
}
