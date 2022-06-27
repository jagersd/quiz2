package controllers

import (
	"net/http"
	"quiz2/config"
	"quiz2/models"

	"github.com/gin-gonic/gin"
)

func GetAllQuestions(c *gin.Context) {
	var questions []models.Question
	config.DB.Find(&questions)

	c.JSON(http.StatusOK, gin.H{"data": questions})
}

func AddQuestion(c *gin.Context) {
	type addQuestionInput struct {
		SubjectId  uint     `json:"subjectId" binding:"required"`
		Type       uint     `json:"type" binding:"required"`
		Attachment string   `json:"attachment"`
		Body       string   `json:"body" binding:"required"`
		Answer     string   `json:"answer" binding:"required"`
		Options    []string `json:"options"`
	}

	var input addQuestionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	question := models.Question{
		SubjectId:  input.SubjectId,
		Type:       input.Type,
		Attachment: input.Attachment,
		Body:       input.Body,
		Answer:     input.Answer,
	}

	config.DB.Create(&question)

	if question.Type == 1 {
		answerAsOption := models.Option{
			QuestionId: question.ID,
			Option:     input.Answer,
		}
		config.DB.Create(&answerAsOption)

		for _, option := range input.Options {
			addOption := models.Option{
				QuestionId: question.ID,
				Option:     option,
			}
			config.DB.Create(&addOption)
		}
		c.JSON(http.StatusOK, gin.H{"data": question, "options": input.Options})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": question})
	}

}
