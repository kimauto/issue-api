package main

import "github.com/gin-gonic/gin"

func respondError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"error": message,
		"code":  code,
	})
}