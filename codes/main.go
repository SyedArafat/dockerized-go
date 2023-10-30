package main

import (
	"export_nrt_report/packages/db"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Hello Arafat")
	initialize()
	conn, err := amqp091.Dial(
		"amqp://guest:guest@common-rabbitmq:5672",
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()
}

func initialize() {
	conn := db.MySqlConnection()
	defer conn.Close()
}
