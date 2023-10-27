package main

import (
	"fmt"
	"os"
)

func main() {
	env := os.Getenv("APP_ENV")
	fmt.Println(env)
	fmt.Println("Hello Worlds")
}
