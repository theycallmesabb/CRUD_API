package main

import (
	"github.com/gin-gonic/gin"
)

type Todo struct {
	Id   int
	Task string
	Done bool
}

var tasks []Todo

func main() {

	r := gin.Default()

	r.POST("/", workS)
	r.GET("/g", WorkD)
	r.POST("/d", doneF)
	r.Run()

}
func workS(c *gin.Context) {
	var task Todo

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "Something went wrong",
		})
		return
	}

	tasks = append(tasks, task)

}
func WorkD(c *gin.Context) {
	c.JSON(200, tasks)
}

func doneF(c *gin.Context) {
	type req struct {
		Id int
	}
	var data req
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	for i := range tasks {
		if tasks[i].Id == data.Id {
			tasks[i].Done = true
			c.JSON(200, gin.H{
				"Task":    "Updated",
				"Message": "Task is done",
			})

			return
		}
	}
	c.JSON(404, gin.H{"error": "Task not found"})
}
