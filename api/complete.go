package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	dbase "github.com/nqnlong/todoapp/database"
)

func CompletedTask(ctx *gin.Context) {
	// var task map[string]string

	task_id := ctx.Param("id")

	if !isNumber(task_id) {
		ctx.JSON(500, gin.H{"error": "task id needed"})
		return
	}

	// body, err := ctx.GetRawData()
	// if err != nil {
	// 	// Handle the error
	// 	ctx.JSON(500, gin.H{"error": err.Error()})
	// 	return
	// }

	// err = json.Unmarshal(body, &task)
	// if err != nil {
	// 	// Handle the error
	// 	ctx.JSON(500, gin.H{"error": err.Error()})
	// 	return
	// }

	query := fmt.Sprintf(`UPDATE tasks SET completed = true WHERE task_id = %s`, task_id)
	_, err := dbase.DB.Exec(query)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "complete task successfully",
	})
}
