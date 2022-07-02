package controllers

import (
	"fmt"
	"net/http"
	"quiz2/config"
	"quiz2/helpers"
	"quiz2/models"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Initiate(c *gin.Context) {
	slug := helpers.GenerateSlug()

	type hostQuizInput struct {
		PlayerName     string `form:"playerName" binding:"required"`
		SubjectId      uint   `form:"subjectId" binding:"required"`
		QuestionAmount uint   `form:"questionAmount"`
	}

	var input hostQuizInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	questions := collectQuestions(int(input.SubjectId), int(input.QuestionAmount))

	quiz := models.Aquiz{
		QuizSlug:  slug,
		Questions: questions,
	}

	config.DB.Create(&quiz)

	if reflect.TypeOf(quiz.ID).String() == "uint" {
		result := models.Result{
			AquizId:    quiz.ID,
			PlayerName: input.PlayerName,
			PlayerSlug: helpers.GenerateSlug(),
			IsHost:     true,
		}
		config.DB.Create(&result)

		c.HTML(http.StatusCreated, "hostquiz.gohtml", gin.H{"data": slug})
	}
}

func Startquiz(c *gin.Context) {

	type startQuizInput struct {
		QuizSlug string `json:"quizSlug" binding:"required"`
	}

	var input startQuizInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var playerCount int64
	config.DB.Model(&models.Aquiz{}).Where("quiz_slug = ?", input.QuizSlug).Count(&playerCount)
	// set solo mode if player count is 1
	if playerCount == 1 {
		var quiz models.Aquiz
		config.DB.Model(&models.Aquiz{}).Select("id").Where("quiz_slug = ?", input.QuizSlug).First(&quiz)
		config.DB.Model(&models.Result{}).Where("aquiz_id = ?", quiz.ID).Update("is_host", false)
	}
	config.DB.Model(&models.Aquiz{}).Where("quiz_slug = ?", input.QuizSlug).Update("started", true)

	c.JSON(http.StatusOK, gin.H{"data": input.QuizSlug + " started with " + strconv.FormatInt(int64(playerCount), 10) + " players"})
}

func Joinquiz(c *gin.Context) {

	type joinInput struct {
		PlayerName string `json:"playerName" binding:"required"`
		QuizSlug   string `json:"quizSlug" binding:"required"`
	}

	var input joinInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var quiz models.Aquiz
	config.DB.Model(&models.Aquiz{}).Select("id").Where("quiz_slug = ?", input.QuizSlug).First(&quiz)

	result := models.Result{
		AquizId:    quiz.ID,
		PlayerName: input.PlayerName,
		PlayerSlug: helpers.GenerateSlug(),
		IsHost:     false,
	}
	config.DB.Create(&result)

	c.JSON(http.StatusCreated, gin.H{"data": input.QuizSlug})

}

func collectQuestions(subjectId, questionAmount int) string {

	if questionAmount == 0 {
		questionAmount = 20
	}

	var questions []models.Question
	config.DB.Select("id").Where(map[string]interface{}{"subject_id": subjectId}).Find(&questions)

	if questionAmount > len(questions) {
		questionAmount = len(questions)
	}

	var returnString string

	for _, question := range questions {
		questionId := strconv.FormatUint(uint64(question.ID), 10)
		addQuestion := fmt.Sprintf("%s,", questionId)
		returnString += addQuestion
	}

	return returnString
}
