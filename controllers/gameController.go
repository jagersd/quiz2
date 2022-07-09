package controllers

import (
	"fmt"
	"net/http"
	"quiz2/config"
	"quiz2/models"
	"strings"

	"github.com/gin-gonic/gin"
)

type RoutineRequest struct {
	PlayerSlug      string `form:"playerName" binding:"required"`
	QuizSlug        string `form:"quizCode" binding:"required"`
	SubmittedAnswer string `form:"answer"`
	Stage           int    `form:"stage" binding:"required"`
}

func HostRoutine(c *gin.Context) {
	var input RoutineRequest
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var currentQuiz models.Aquiz
	config.DB.Model(&models.Aquiz{}).Where("quiz_slug = ?", input.QuizSlug).
		First(&currentQuiz)

	questionId := strings.Split(currentQuiz.Questions, ",")[input.Stage-1]
	currentQuestion, options := getQuestion(questionId)

	config.DB.Model(&models.Result{}).
		Where("player_slug = ? AND aquiz_id = ?", input.PlayerSlug, currentQuiz.ID).
		Update(fmt.Sprintf("result%v", input.Stage), 1)

	c.HTML(http.StatusOK, "hostview.gohtml", gin.H{
		"question": currentQuestion.Body,
		"answer":   currentQuestion.Answer,
		"options":  options,
		"hostSlug": input.PlayerSlug,
		"quizSlug": input.QuizSlug,
		"quizId":   currentQuiz.ID,
		"stage":    input.Stage + 1})
}

func GetLiveResults(c *gin.Context) {

	type apiReponse struct {
		PlayerName string
		Result     uint8
	}

	var liveResults []apiReponse

	fmt.Printf("%T \n", c.Param("stage"))

	if err := config.DB.Table("results").Select("player_name", fmt.Sprintf("result%v as Result", c.Param("stage"))).
		Where("aquiz_id = ? AND is_host = ?", c.Param("quizId"), 0).
		Find(&liveResults).Error; err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err})

		return
	}

	c.JSON(http.StatusOK, gin.H{"playername": liveResults})

}

func ParticipantRoutine(c *gin.Context) {
	var input RoutineRequest
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func getQuestion(questionId string) (models.Question, []string) {
	var currentQuestion models.Question
	config.DB.Model(&models.Question{}).Where("id = ?", questionId).First(&currentQuestion)

	var options []string
	config.DB.Model(&models.Option{}).Select("option").Where("question_id = ?", currentQuestion.ID).Find(&options)

	return currentQuestion, options
}
