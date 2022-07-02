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

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.gohtml", "")
	})

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Yes, we up."})
	})

	r.GET("/starthost", controllers.GetSubjects)

	//subject routes
	r.POST("/addsubject", controllers.AddSubject)

	//question routes
	r.GET("/allquestions", controllers.GetAllQuestions)
	r.POST("/addquestion", controllers.AddQuestion)

	//quiz routes
	r.POST("/hostquiz", controllers.Initiate)
	r.POST("/startquiz", controllers.Startquiz)
	r.POST("/joinquiz", controllers.Joinquiz)

	//admin routes
	r.DELETE("/resetdb", controllers.ResetDb)
	r.Run()
}
