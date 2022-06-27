package controllers

import (
	"quiz2/config"

	"github.com/gin-gonic/gin"
)

func ResetDb(c *gin.Context) {
	config.DB.Exec("Delete FROM aquizzes")
	config.DB.Exec("Delete FROM results")
}
