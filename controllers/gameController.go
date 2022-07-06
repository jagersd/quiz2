package controllers

import (
	"net/http"
	"quiz2/config"
	"quiz2/models"
	"strings"

	"github.com/gin-gonic/gin"
)

type HostRequest struct {
	PlayerName string `form:"playerName" binding:"required"`
	QuizSlug   string `form:"quizCode" binding:"required"`
	Stage      int    `form:"stage" binding:"required"`
}

func HostRoutine(c *gin.Context) {
	var input HostRequest
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var currentQuiz models.Aquiz
	config.DB.Model(&models.Aquiz{}).Where("quiz_slug = ?", input.QuizSlug).First(&currentQuiz)

	questionId := strings.Split(currentQuiz.Questions, ",")[input.Stage-1]

	currentQuestion, options := getQuestion(questionId)

	c.HTML(http.StatusOK, "hostview.gohtml", gin.H{"question": currentQuestion.Body,
		"answer":  currentQuestion.Answer,
		"options": options})
}

func getQuestion(questionId string) (models.Question, []string) {
	var currentQuestion models.Question
	config.DB.Model(&models.Question{}).Where("id = ?", questionId).First(&currentQuestion)

	var options []string
	config.DB.Model(&models.Option{}).Select("option").Where("question_id = ?", currentQuestion.ID).Find(&options)

	return currentQuestion, options
}
