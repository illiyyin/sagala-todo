package main

import (
	"fmt"

	"github.com/illiyyin/sagala-todo/database"
)

func main ()  {
	fmt.Println("go running")
	database.ConnectDB()
}