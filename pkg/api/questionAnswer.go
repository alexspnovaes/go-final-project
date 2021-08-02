package api

import (
	"final_project/internal/data"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/questions", data.GetQuestions)
	r.GET("/questions/:id", data.GetQuestionById)
	r.GET("/users/:id/questions", data.GetQuestionsByUserId)
	r.POST("/questions", data.InsertQuestion)
	r.PUT("/questions/:id", data.EditQuestion)
	r.DELETE("/questions/:id", data.DeleteQuestion)
	log.Fatal(http.ListenAndServe(":10000", r))
}
