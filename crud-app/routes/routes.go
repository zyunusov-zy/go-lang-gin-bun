package routes

import (
	"crud-app/handlers"
	"crud-app/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	taskService := services.NewInMemoryTaskService()
	taskHandler := handlers.NewTaskHandler(taskService)

	r.POST("/tasks", taskHandler.CreateTask)
	r.GET("/tasks", taskHandler.GetTasks)
	r.GET("/tasks/:id", taskHandler.GetTaskById)
	r.PUT("/tasks/:id", taskHandler.UpdateTask)
	r.DELETE("/tasks/:id", taskHandler.DeleteTask)
}