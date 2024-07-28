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
	router.GET("/task/:id", handler.HandlerGetTask())
	router.GET("/tasks", handler.HandlerGetAllTask())
	router.POST("/task", handler.HandlerCreateTask())
	router.PATCH("/task/:id", handler.HandlerUpdateTask())

	router.Run(":3000")
}