package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"quiz2/config"
	"quiz2/models"
)

func AddSubject(c *gin.Context) {
	type addSubjectInput struct {
		Name        string `form:"subjectName" binding:"required"`
		Description string `form:"description"`
	}

	var input addSubjectInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subject := models.Subject{
		Name:        input.Name,
		Description: input.Description,
	}

	config.DB.Create(&subject)

	var subjects []models.Subject
	config.DB.Find(&subjects)

	c.HTML(http.StatusOK, "addquiz.gohtml", gin.H{"subjects": subjects})
}

func GetSubjects(c *gin.Context) {
	var subjects []models.Subject
	config.DB.Find(&subjects)

	c.HTML(http.StatusOK, "initiate.gohtml", gin.H{"subjects": subjects})
}
