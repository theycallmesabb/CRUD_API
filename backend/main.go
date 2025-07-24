package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"sync"
)

type Todo struct {
	Id   int
	Task string
	Done bool
}

// created a slice of data-type : ToDo (to store the tasks)
var tasks []Todo

// Error handling function
func errHandle(c *gin.Context, err error) bool {
	if err != nil {
		c.JSON(400, gin.H{
			"Error": err,
		})
		return true
	}
	return false
}

// POST
func writeTask(c *gin.Context) {
	var task Todo
	err := c.ShouldBindJSON(&task)
	//error handling
	if errHandle(c, err) {
		return
	}

	tasks = append(tasks, task)
}

// DELETE
func deleteTask(c *gin.Context) {
	//defined the struct inside this function because, we don't need this anymore after this function ends.
	type upt struct {
		Id int
	}
	var del upt
	err := c.ShouldBindJSON(&del)

	//error handling
	if errHandle(c, err) {
		return
	}

	for i := range tasks {
		if tasks[i].Id == del.Id {
			//if task is done , delete from slice
			tasks = append(tasks[:i], tasks[i+1:]...)
			return
		}
	}
	c.JSON(404, gin.H{
		"message": "Task not found !!",
	})
}

// PUT
func updateTask(c *gin.Context) {
	var input Todo
	err := c.ShouldBindJSON(&input)

	//error handling
	if errHandle(c, err) {
		return
	}

	for i := range tasks {
		if tasks[i].Id == input.Id {
			//updating the task
			tasks[i] = input
			return
		}
	}
	c.JSON(404, gin.H{
		"message": "Task Id not found",
	})

}

// GET
func fetchTask(c *gin.Context) {
	c.JSON(200, tasks)
}

func main() {
	fmt.Println("Welcome to this server")
	//starting Gin
	r := gin.Default()

	r.Static("/static", "../frontend")
	r.StaticFile("/", "../frontend/index.html")

	r.POST("/add", writeTask)
	r.GET("/show", fetchTask)
	r.PUT("/update", updateTask)
	r.DELETE("/delete", deleteTask)

	r.Run(":8000")

}
