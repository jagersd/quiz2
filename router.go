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
	r.LoadHTMLGlob("ui/*.gohtml")
	config.ConnectDatabase()

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

	/*
		/
		/ Following routes are meant for API/json requests and responses
		/
	*/

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Yes, we up."})
	})

	//subject routes
	r.POST("/addsubject", controllers.AddSubject)

	//question routes
	r.GET("/allquestions", controllers.GetAllQuestions)
	r.POST("/addquestion", controllers.AddQuestion)

	//quiz routes

	r.POST("/startquiz", controllers.Startquiz)
	r.POST("/joinquiz", controllers.Joinquiz)

	//admin routes
	r.DELETE("/resetdb", controllers.ResetDb)
	r.Run()
}
