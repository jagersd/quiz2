package controllers

import (
	"fmt"
	"net/http"
	"quiz2/config"
	"quiz2/models"
	"strconv"
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

	//break routine if quiz completed
	if input.Stage == len(strings.Split(currentQuiz.Questions, ",")) {
		finalResult := getFinalResult(currentQuiz.ID)
		c.HTML(http.StatusOK, "results.gohtml", gin.H{"totals": finalResult})
		return
	}

	questionId := strings.Split(currentQuiz.Questions, ",")[input.Stage-1]
	currentQuestion, options := getQuestion(questionId)

	questionAmount := fmt.Sprintf("(%v / %v)", input.Stage, len(strings.Split(currentQuiz.Questions, ",")))

	config.DB.Model(&models.Result{}).
		Where("player_slug = ? AND aquiz_id = ?", input.PlayerSlug, currentQuiz.ID).
		Update(fmt.Sprintf("result%v", input.Stage), 1)

	c.HTML(http.StatusOK, "hostview.gohtml", gin.H{
		"question":       currentQuestion.Body,
		"answer":         currentQuestion.Answer,
		"questionAmount": questionAmount,
		"options":        options,
		"hostSlug":       input.PlayerSlug,
		"quizSlug":       input.QuizSlug,
		"quizId":         currentQuiz.ID,
		"stage":          input.Stage + 1})
}

func GetLiveResults(c *gin.Context) {

	type apiReponse struct {
		PlayerName string
		Result     *uint8
		Total      uint
	}

	var liveResults []apiReponse

	if err := config.DB.Table("results").Select("player_name", fmt.Sprintf("result%v as Result", c.Param("stage")), "total").
		Where("aquiz_id = ? AND is_host = ?", c.Param("quizId"), 0).
		Order("total desc").
		Find(&liveResults).Error; err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": liveResults})

}

func ParticipantRoutine(c *gin.Context) {
	var input RoutineRequest
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var currentQuiz models.Aquiz
	config.DB.Model(&models.Aquiz{}).Where("quiz_slug = ?", input.QuizSlug).First(&currentQuiz)

	lastResult := false

	if input.Stage != 1 {
		questionId := strings.Split(currentQuiz.Questions, ",")[input.Stage-2]
		previousQuestion, _ := getQuestion(questionId)

		if strings.EqualFold(previousQuestion.Answer, input.SubmittedAnswer) {
			lastResult = true
		}

		processResult(currentQuiz.ID, input.PlayerSlug, input.Stage, lastResult)
	}

	//break routine if quiz ends
	if input.Stage == len(strings.Split(currentQuiz.Questions, ",")) {
		finalResult := getFinalResult(currentQuiz.ID)
		c.HTML(http.StatusOK, "results.gohtml", gin.H{"totals": finalResult})
		return
	}

	currentQuestion, options := getQuestion(strings.Split(currentQuiz.Questions, ",")[input.Stage-1])

	c.HTML(http.StatusOK, "participantview.gohtml", gin.H{
		"playerSlug": input.PlayerSlug,
		"lastResult": lastResult,
		"question":   currentQuestion.Body,
		"type":       currentQuestion.Type,
		"options":    options,
		"quizSlug":   input.QuizSlug,
		"quizId":     currentQuiz.ID,
		"stage":      input.Stage + 1,
	})
}

func RevealNextQuestion(c *gin.Context) {

	var result int
	var response bool

	quizId, err := strconv.Atoi(c.Param("quizId"))
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err})
	}

	if err := config.DB.Table("results").Select(fmt.Sprintf("ifnull (result%v , 0)", c.Param("stage"))).
		Where("aquiz_id = ? AND is_host = ?", quizId, 1).
		Take(&result).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if result == 1 {
		response = true
	}

	c.JSON(http.StatusOK, gin.H{"result": response})
}

func getQuestion(questionId string) (models.Question, []string) {
	var currentQuestion models.Question
	config.DB.Model(&models.Question{}).Where("id = ?", questionId).First(&currentQuestion)

	var options []string
	config.DB.Model(&models.Option{}).Select("option").Where("question_id = ?", currentQuestion.ID).Find(&options)

	return currentQuestion, options
}

func processResult(currentQuiz uint, playerSlug string, stage int, lastResult bool) {

	type processor struct {
		Result *bool
		Total  uint
	}

	//check whether response already provided
	var checker processor

	if err := config.DB.Table("results").Select(fmt.Sprintf("result%v as Result", stage-1), "total as Total").
		Where("aquiz_id = ? AND player_slug = ?", currentQuiz, playerSlug).
		Find(&checker).Error; err != nil {
		return
	}

	fmt.Printf("checker: %v \n lastresult:%v", checker, lastResult)

	if lastResult {
		checker.Total += 1
	}

	if checker.Result == nil {
		config.DB.Model(&models.Result{}).
			Where("player_slug = ? AND aquiz_id = ?", playerSlug, currentQuiz).
			Updates(map[string]interface{}{fmt.Sprintf("result%v", stage-1): lastResult, "total": checker.Total})
	}

}

func getFinalResult(quizId uint) []models.Result {

	var results []models.Result

	if err := config.DB.Model(&models.Result{}).
		Where("aquiz_id = ? AND is_host = ?", quizId, 0).
		Order("total").
		Find(&results).Error; err != nil {
		fmt.Print(err)
	}

	return results
}
