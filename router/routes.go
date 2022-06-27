package router

import (
	"net/http"
	"quiz2/config"
	"quiz2/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	config.ConnectDatabase()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Yes, we up."})
	})

	//subject routes
	r.GET("/allsubjects", controllers.GetSubjects)
	r.POST("/addsubject", controllers.AddSubject)

	//question routes
	r.GET("/allquestions", controllers.GetAllQuestions)
	r.POST("/addquestion", controllers.AddQuestion)

	//quiz routes
	r.POST("/hostquiz", controllers.Initiate)
	r.POST("/startquiz", controllers.Startquiz)

	//admin routes
	r.DELETE("/resetdb", controllers.ResetDb)
	r.Run()
}
