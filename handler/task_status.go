package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/illiyyin/sagala-todo/database"
	"github.com/illiyyin/sagala-todo/model"
)

func HandlerCreateTaskStatus() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var reqBody model.TaskStatusRequest
		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			log.Println(err)
			return
		}

		taskStatus := &model.TaskStatus{
			StatusName: reqBody.StatusName,
		}

		if err := database.DB.Db.Create(&taskStatus).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Task Status", "message": err.Error()})
			fmt.Println(err)
			return
		}

		resBody := &model.TaskStatusResponse{
			ID:         int(taskStatus.ID),
			StatusName: reqBody.StatusName,
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    resBody,
			"message": "Success create Task Status",
		})
	}

	return gin.HandlerFunc(fn)
}
