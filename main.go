package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/illiyyin/sagala-todo/database"
	"github.com/illiyyin/sagala-todo/docs"
	"github.com/illiyyin/sagala-todo/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Sagala Todo API
//	@version		1.0

func main() {
	fmt.Println("go running")
	database.ConnectDB()

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	router.POST("/task-status", handler.HandlerCreateTaskStatus())
	router.GET("/task/:id", handler.HandlerGetTask())
	router.GET("/tasks", handler.HandlerGetAllTask())
	router.POST("/task", handler.HandlerCreateTask())
	router.PATCH("/task/:id", handler.HandlerUpdateTask())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":3000")
}
