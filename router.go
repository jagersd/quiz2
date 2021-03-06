package main

import (
	"net/http"
	"quiz2/config"
	"quiz2/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.Static("/static", "ui/static")
	r.LoadHTMLGlob("ui/html/**/*.gohtml")
	config.Roachconn(false)

	/*
		/
		/ Following routes are meant for html requests and responses
		/
	*/

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.gohtml", "")
	})

	r.GET("/initiate", controllers.GetSubjects)
	r.POST("/hostquiz", controllers.Initiate)
	r.POST("/hostroutine", controllers.HostRoutine)

	r.GET("/join", func(c *gin.Context) {
		c.HTML(http.StatusOK, "joinquiz.gohtml", "")
	})

	r.POST("/joined", controllers.Joinquiz)
	r.POST("/participantroutine", controllers.ParticipantRoutine)

	r.GET("/addquiz", controllers.GetAllSubjects)
	r.POST("/addquiz", controllers.AddQuestion)

	/*
		/
		/ Following routes are meant for API/json requests and responses
		/
	*/
	r.GET("/waitingroom/:quizSlug/:playerSlug", controllers.Waitingroom)
	r.POST("/startquiz", controllers.Startquiz)

	r.GET("/liveresults/:quizId/:quizSlug/:stage", controllers.GetLiveResults)

	r.GET("/readytoreveal/:quizId/:stage", controllers.RevealNextQuestion)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Yes, we up."})
	})

	//subject routes
	r.POST("/addsubject", controllers.AddSubject)

	//question routes
	r.GET("/allquestions", controllers.GetAllQuestions)

	//admin routes
	r.DELETE("/resetdb", controllers.ResetDb)
	r.Run()
}
