package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"quiz2/config"
	"quiz2/models"
)

func AddSubject(c *gin.Context) {
	type addSubjectInput struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
	}

	var input addSubjectInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subject := models.Subject{
		Name:        input.Name,
		Description: input.Description,
	}

	config.DB.Create(&subject)
	c.JSON(http.StatusOK, gin.H{"data": subject})
}

func GetSubjects(c *gin.Context) {
	var subjects []models.Subject
	config.DB.Find(&subjects)

	c.HTML(http.StatusOK, "starthost.gohtml", gin.H{"subjects": subjects})
}
