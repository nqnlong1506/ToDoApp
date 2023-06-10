package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	dbase "github.com/nqnlong/todoapp/database"
	"github.com/nqnlong/todoapp/model"
)

func IndexHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello world from indexHandler",
	})
}

func GetTasks(ctx *gin.Context) {
	rows, err := dbase.DB.Query("SELECT * FROM tasks")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var list []model.Task

	// Iterate over the query results
	for rows.Next() {
		var (
			id       int
			name     string
			complete bool
		)
		// Process the row data
		rows.Scan(&id, &name, &complete)

		task := model.Task{
			Task_id:   id,
			Task_name: name,
			Completed: complete}

		list = append(list, task)
	}

	// Check for any errors during iteration
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "get all tasks from database",
		"data":    list,
	})
}

func AddTask(ctx *gin.Context) {

	var task map[string]string

	body, err := ctx.GetRawData()
	if err != nil {
		// Handle the error
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	err = json.Unmarshal(body, &task)
	if err != nil {
		// Handle the error
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	query := fmt.Sprintf(`INSERT INTO tasks (task_name) VALUES ('%s')`, task["task"])
	_, err = dbase.DB.Exec(query)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "add task successfully",
	})
}
