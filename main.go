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

var currentUser = ""

func main() {
	router := gin.Default()

	{
		dbase.InitlizeDatabase()
	}

	// To Do App routers
	router.GET("/", info) // project info

	router.POST("/login", handleLogini)

	routerGroup := router.Group("/tasks", userAuthentication)
	{
		routerGroup.GET("", api.GetTasks)          // get all tasks
		routerGroup.POST("/add", api.AddTask)      // create task
		routerGroup.PUT("/:id", api.CompletedTask) // complete task
		routerGroup.DELETE("/:id", api.DeleteTask) // delete task
	}

	// Start the server
	router.Run(":8080")
}

func userAuthentication(ctx *gin.Context) {
	if currentUser == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		ctx.Abort()
		return
	}
	ctx.Next()
}

func handleLogini(ctx *gin.Context) {
	currentUser = "nqnlong"

	ctx.JSON(http.StatusOK, gin.H{
		"message": "login successfully",
	})
}
