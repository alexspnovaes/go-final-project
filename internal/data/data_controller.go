package data

import (
	"context"
	model "final_project/internal/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

//TODO: do some validations on the inserts/updates to see if there is a value on the data model

const dbName = "questionsAnswers"
const collection = "question"

var questionCollection = db().Database(dbName).Collection(collection) //TODO: the connectionstring could be on a settings

func InsertQuestion(c *gin.Context) {
	var q model.Question
	c.BindJSON(&q)
	newQuestion := model.Question{
		Id:       guuid.New().String(),
		User:     q.User,
		Question: q.Question,
		Answer: model.Answer{
			Id:        guuid.New().String(),
			User:      q.User,
			Text:      q.Answer.Text,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err := questionCollection.InsertOne(context.TODO(), newQuestion)
	if err != nil {
		log.Printf("Error while inserting new question into db, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": newQuestion.Id,
	})
}

func GetQuestions(c *gin.Context) {
	questions := []model.Question{}
	cursor, err := questionCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Printf("Error while getting all questions, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	// Iterate through the returned cursor.
	for cursor.Next(context.TODO()) {
		var q model.Question
		cursor.Decode(&q)
		questions = append(questions, q)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   questions,
	})
}

func GetQuestionById(c *gin.Context) {
	id := c.Param("id")
	q := model.Question{}
	err := questionCollection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&q)
	if err != nil {
		log.Printf("Error while getting a single question, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Question not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   q,
	})
}

func GetQuestionsByUserId(c *gin.Context) {
	id := c.Param("id")
	questions := []model.Question{}
	cursor, err := questionCollection.Find(context.TODO(), bson.M{"user": id})
	if err != nil {
		log.Printf("Error while getting questions by user, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Question not found",
		})
		return
	}
	if err != nil {
		log.Printf("Error while getting questions by user, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	// Iterate through the returned cursor.
	for cursor.Next(context.TODO()) {
		var q model.Question
		cursor.Decode(&q)
		questions = append(questions, q)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   questions,
	})
}

func EditQuestion(c *gin.Context) {
	id := c.Param("id")
	var question model.Question
	question.UpdatedAt = time.Now()
	question.Answer.UpdatedAt = time.Now()
	c.BindJSON(&question)
	newData := bson.M{
		"$set": bson.M{
			"user":     question.User,
			"question": question.Question,
			"answer":   question.Answer,
		},
	}
	_, err := questionCollection.UpdateOne(context.TODO(), bson.M{"id": id}, newData)
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
	})
}

func DeleteQuestion(c *gin.Context) {
	id := c.Param("id")
	//TODO: do a logic to see if the question exists
	_, err := questionCollection.DeleteOne(context.TODO(), bson.M{"id": id})
	if err != nil {
		log.Printf("Error while deleting a single question, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}
