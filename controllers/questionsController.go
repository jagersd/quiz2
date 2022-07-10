package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"quiz2/config"
	"quiz2/models"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllQuestions(c *gin.Context) {
	var questions []models.Question
	config.DB.Find(&questions)

	c.JSON(http.StatusOK, gin.H{"data": questions})
}

func GetAllSubjects(c *gin.Context) {
	var subjects []models.Subject
	config.DB.Find(&subjects)

	c.HTML(http.StatusOK, "addquiz.gohtml", gin.H{"subjects": subjects})
}

func AddQuestion(c *gin.Context) {
	type addQuestionInput struct {
		SubjectId       uint   `form:"subjectId" binding:"required"`
		Type            uint   `form:"questionType" binding:"required"`
		Attachment      string `form:"attachment"`
		Body            string `form:"questionBody" binding:"required"`
		Answer          string `form:"questionAnswer" binding:"required"`
		QuestionOption1 string `form:"questionOption1"`
		QuestionOption2 string `form:"questionOption2"`
		QuestionOption3 string `form:"questionOption3"`
		QuestionOption4 string `form:"questionOption4"`
		QuestionOption5 string `form:"questionOption5"`
		QuestionOption6 string `form:"questionOption6"`
	}

	var input addQuestionInput
	if err := c.ShouldBind(&input); err != nil {
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

	var subjects []models.Subject
	config.DB.Find(&subjects)

	if question.Type == 1 {
		options := []string{input.QuestionOption1,
			input.QuestionOption2,
			input.QuestionOption3,
			input.QuestionOption4,
			input.QuestionOption5,
			input.QuestionOption6,
			input.Answer}

		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(options), func(i, j int) { options[i], options[j] = options[j], options[i] })
		fmt.Println(options)

		for _, option := range options {
			if option != "" {
				addOption := models.Option{
					QuestionId: question.ID,
					Option:     option,
				}
				config.DB.Create(&addOption)
			}
		}
		c.HTML(http.StatusOK, "addquiz.gohtml", gin.H{"subjects": subjects})
	} else {
		c.HTML(http.StatusOK, "addquiz.gohtml", gin.H{"subjects": subjects})
	}

}
