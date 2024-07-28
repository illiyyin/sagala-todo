package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/illiyyin/sagala-todo/database"
	"github.com/illiyyin/sagala-todo/model"
)

func HandlerCreateTask() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var reqBody model.TaskRequest
		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			log.Println(err)
			return
		}

		task := &model.Task{
			Title:       reqBody.Title,
			Description: reqBody.Description,
			StatusID:    reqBody.StatusID,
		}

		status := &model.TaskStatus{
			ID: reqBody.StatusID,
		}
		if err := database.DB.Db.First(&status).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Status id not found", "message": err.Error()})
			fmt.Println(err)
			return
		}

		resultTaskStatus := &model.TaskStatusResponse{
			ID:         status.ID,
			StatusName: status.StatusName,
		}

		if err := database.DB.Db.Create(&task).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Task Status", "message": err.Error()})
			fmt.Println(err)
			return
		}


		resBody := &model.TaskResponse{
			ID:          int(task.ID),
			Title:       task.Title,
			Description: task.Description,
			StatusID:    task.StatusID,
			Status:      *resultTaskStatus,
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    resBody,
			"message": "Success create Task Status",
		})
	}

	return gin.HandlerFunc(fn)
}
