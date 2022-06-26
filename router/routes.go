package router

import (
	"net/http"
	"quiz2/controllers"
	"quiz2/models"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	models.ConnectDatabase()

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

	r.Run()
}
