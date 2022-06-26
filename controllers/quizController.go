package controllers

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"

	"quiz2/helpers"
	"quiz2/models"
)

func Initiate(c *gin.Context) {
	slug := helpers.GenerateSlug()

	type startQuizInput struct {
		PlayerName     string `json:"playerName" binding:"required"`
		SubjectId      uint   `json:"subjectId" binding:"required"`
		QuestionAmount uint   `json:"questionAmount"`
	}

	var input startQuizInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	questions := collectQuestions(int(input.SubjectId), int(input.QuestionAmount))

	quiz := models.Aquiz{
		QuizSlug:  slug,
		Questions: questions,
	}

	models.DB.Create(&quiz)

	if reflect.TypeOf(quiz.ID).String() == "uint" {
		result := models.Result{
			AquizId:    quiz.ID,
			PlayerName: input.PlayerName,
			PlayerSlug: helpers.GenerateSlug(),
			IsHost:     true,
		}
		models.DB.Create(&result)

		c.JSON(http.StatusOK, gin.H{"data": slug, "questions": questions})
	}
}

func collectQuestions(subjectId, questionAmount int) string {

	if questionAmount == 0 {
		questionAmount = 20
	}

	var questions []models.Question
	models.DB.Where(map[string]interface{}{"subject_id": subjectId}).Find(&questions)

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
