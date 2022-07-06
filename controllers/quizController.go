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

		c.HTML(http.StatusCreated, "hostquiz.gohtml", gin.H{
			"quizSlug": slug,
			"hostSlug": result.PlayerSlug,
		})
	}
}

func Startquiz(c *gin.Context) {

	type startQuizInput struct {
		PlayerSlug string `json:"playerSlug" binding:"required"`
		QuizSlug   string `json:"quizSlug" binding:"required"`
	}

	var input startQuizInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.PlayerSlug != input.QuizSlug {
		c.JSON(http.StatusExpectationFailed, gin.H{"error": "incorrect party initiated quiz"})
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
		PlayerName string `form:"playerName" binding:"required"`
		QuizSlug   string `form:"quizSlug" binding:"required"`
	}

	var input joinInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var quiz models.Aquiz
	config.DB.Model(&models.Aquiz{}).Select("id").Where("quiz_slug = ?", input.QuizSlug).First(&quiz)

	newResultRecord := models.Result{
		AquizId:    quiz.ID,
		PlayerName: input.PlayerName,
		PlayerSlug: helpers.GenerateSlug(),
		IsHost:     false,
	}

	if quiz.ID != 0 {
		config.DB.Create(&newResultRecord)
	}

	c.HTML(http.StatusCreated, "joined.gohtml", gin.H{"quizCode": input.QuizSlug, "playerSlug": newResultRecord.PlayerSlug})
}

func Waitingroom(c *gin.Context) {

	var quizState models.Aquiz
	if err := config.DB.Where("quiz_slug = ?", c.Param("quizSlug")).Find(&quizState).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Quiz not found"})
		return
	}

	var players []string
	if quizState.ID != 0 {
		config.DB.Model(&models.Result{}).Select("player_name").Where("aquiz_id =?", quizState.ID).Find(&players)
	}

	c.JSON(http.StatusOK, gin.H{"quiz": quizState, "players": players})
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
