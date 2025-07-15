package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	InitData()
	r := gin.Default()

	r.POST("/issue", CreateIssue)
	r.GET("/issues", GetIssues)
	r.GET("/issue/:id", GetIssueByID)
	r.PATCH("/issue/:id", UpdateIssue)

	r.Run(":8080")
}
