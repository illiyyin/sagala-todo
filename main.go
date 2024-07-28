package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/illiyyin/sagala-todo/database"
	"github.com/illiyyin/sagala-todo/handler"
)

func main ()  {
	fmt.Println("go running")
	database.ConnectDB()

	router := gin.Default()

	router.POST("/task-status", handler.HandlerCreateTaskStatus())

	router.Run(":3000")
}