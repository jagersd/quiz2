package controllers

import (
	"net/http"
	"quiz2/config"

	"github.com/gin-gonic/gin"
)

func ResetDb(c *gin.Context) {
	config.DB.Exec("TRUNCATE aquizzes")
	config.DB.Exec("TRUNCATE results")
	c.JSON(http.StatusOK, gin.H{"data": "db clean"})
}
