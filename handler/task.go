package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/illiyyin/sagala-todo/database"
	"github.com/illiyyin/sagala-todo/model"
)

var formatLayoutWithoutHour = "2006-01-02"

func convertDate(date string) string {
	if date == "" {
		return ""
	}
	if strings.Contains(date, "T") {
		return date
	}
	parsedDate, err := time.Parse(formatLayoutWithoutHour, date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return ""
	}
	return parsedDate.Format(time.RFC3339)
}

func HandlerCreateTask() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var reqBody model.TaskRequest
		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			log.Println(err)
			return
		}

		if reqBody.Title == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
			return
		}

		task := &model.Task{
			Title:       reqBody.Title,
			Description: &reqBody.Description,
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
			ID:           task.ID,
			Title:        task.Title,
			Description:  *task.Description,
			ExpectedDate: convertDate(reqBody.ExpectedDate),
			StatusID:     task.StatusID,
			Status:       *resultTaskStatus,
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    resBody,
			"message": "Success create Task",
		})
	}

	return gin.HandlerFunc(fn)
}

func HandlerUpdateTask() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id := c.Param("id")
		var task model.Task

		if err := database.DB.Db.First(&task, "id = ?", id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}

		var reqBody model.TaskRequest
		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			log.Println(err)
			return
		}

		if reqBody.Title != "" {
			task.Title = reqBody.Title
		}
		if reqBody.Description != "" {
			task.Description = &reqBody.Description
		}
		if reqBody.ExpectedDate != "" {
			date := convertDate(reqBody.ExpectedDate)
			task.ExpectedDate = &date
		}
		var status model.TaskStatus
		if reqBody.StatusID != 0 {
			if err := database.DB.Db.First(&status, "id = ?", reqBody.StatusID).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Task Status Not found"})
				return
			}
			task.StatusID = reqBody.StatusID
		}

		if err := database.DB.Db.Save(&task).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		resultTaskStatus := &model.TaskStatusResponse{
			ID:         status.ID,
			StatusName: status.StatusName,
		}

		resBody := &model.TaskResponse{
			ID:           task.ID,
			Title:        task.Title,
			Description:  *task.Description,
			ExpectedDate: *task.ExpectedDate,
			StatusID:     reqBody.StatusID,
			Status:       *resultTaskStatus,
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    resBody,
			"message": "Success update Task",
		})
	}

	return gin.HandlerFunc(fn)
}

func HandlerGetTask() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id := c.Param("id")
		var task model.Task

		if err := database.DB.Db.Preload("Status").First(&task, "id = ?", id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}

		resultTaskStatus := &model.TaskStatusResponse{
			ID:         task.Status.ID,
			StatusName: task.Status.StatusName,
		}

		resBody := &model.TaskResponse{
			ID:          task.ID,
			Title:       task.Title,
			Description: *task.Description,
			StatusID:    task.StatusID,
			Status:      *resultTaskStatus,
		}
		if task.ExpectedDate == nil {
			resBody.ExpectedDate = ""
		} else {
			resBody.ExpectedDate = convertDate(*task.ExpectedDate)
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    resBody,
			"message": "Success get Task",
		})
	}
	return gin.HandlerFunc(fn)
}

func HandlerDeleteTask() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id := c.Param("id")
		var task model.Task

		if err := database.DB.Db.Delete(&task, "id = ?", id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Task deleted",
		})
	}
	return gin.HandlerFunc(fn)
}

func HandlerGetAllTask() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		statusID := c.Query("status_id")

		var status model.TaskStatus

		var tasks []model.Task
		if statusID != "" {
			if err := database.DB.Db.First(&status, "id = ?", statusID).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Status Task not found"})
				return
			}

			if err := database.DB.Db.Preload("Status").Where("status_id = ?", statusID).Find(&tasks).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
				return
			}
		} else {
			if err := database.DB.Db.Preload("Status").Find(&tasks).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
				return
			}
		}

		var resBody []model.TaskResponse
		for _, task := range tasks {
			resultTaskStatus := &model.TaskStatusResponse{
				ID:         task.Status.ID,
				StatusName: task.Status.StatusName,
			}
			taskTmp := model.TaskResponse{
				ID:           task.ID,
				Title:        task.Title,
				Description:  *task.Description,
				StatusID:     task.StatusID,
				Status:       *resultTaskStatus,
			}
			if task.ExpectedDate == nil {
				taskTmp.ExpectedDate = ""
			} else {
				taskTmp.ExpectedDate = convertDate(*task.ExpectedDate)
			}

			resBody = append(resBody, taskTmp)
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    resBody,
			"message": "Success get All Tasks",
		})
	}
	return gin.HandlerFunc(fn)
}
