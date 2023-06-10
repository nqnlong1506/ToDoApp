package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nqnlong/todoapp/api"
	dbase "github.com/nqnlong/todoapp/database"
)

func info(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"detail": "To-Do-App Project",
	})
}

func main() {
	router := gin.Default()

	{
		dbase.InitlizeDatabase()
	}

	// To Do App routers
	router.GET("/", info) // project info

	router.GET("/tasks", api.GetTasks) // get all tasks

	router.POST("/tasks/add", api.AddTask) // create task

	router.PUT("/tasks/:id", api.CompletedTask) // complete task

	router.DELETE("/tasks/:id", api.DeleteTask) // delete task

	// Start the server
	router.Run(":80")
}
