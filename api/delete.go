package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	dbase "github.com/nqnlong/todoapp/database"
)

func DeleteTask(ctx *gin.Context) {
	task_id := ctx.Param("id")

	if !isNumber(task_id) {
		ctx.JSON(500, gin.H{"error": "task id needed"})
		return
	}

	query := fmt.Sprintf(`DELETE FROM tasks WHERE task_id = %s`, task_id)
	_, err := dbase.DB.Exec(query)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		log.Fatal(err)
	}

	query = `SELECT COUNT(*) from tasks`
	var count int
	err = dbase.DB.QueryRow(query).Scan(&count)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	query = fmt.Sprintf(`SELECT setval('tasks_task_id_seq', %d)`, count)
	_, err = dbase.DB.Exec(query)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete task successfully",
	})
}
