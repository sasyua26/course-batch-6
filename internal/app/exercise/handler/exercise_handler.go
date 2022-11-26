package handler

import (
	"exercise/internal/app/domain"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Define gorm db
type ExerciseHandler struct {
	db *gorm.DB
}

// Define ExerciseHandler
func NewExerciseHandler(db *gorm.DB) *ExerciseHandler {
	return &ExerciseHandler{db: db}
}

// Function to create new exercise
func (eh ExerciseHandler) CreateExercise(c *gin.Context) {
	var createExercise domain.CreateExercise
	if err := c.ShouldBind(&createExercise); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid body",
		})
	}

	exer, err := domain.NewExercise(createExercise.Title, createExercise.Description)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	if err := eh.db.Create(exer).Error; err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]string{
		"message": "New exercise successfully created",
	})
}

// Function to display all question in an exercise based of exercise id parameter :id
func (eh ExerciseHandler) GetExerciseByID(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid exercise id",
		})
		return
	}

	var exercise domain.Exercise
	err = eh.db.Where("id = ?", id).Preload("Question").Take(&exercise).Error
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": "exercise not found",
		})
		return
	}

	c.JSON(http.StatusOK, exercise)
}

// Function to calculate score based on user answer
func (eh ExerciseHandler) GetScore(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid exercise id",
		})
		return
	}

	var exercise domain.Exercise
	err = eh.db.Where("id = ?", id).Preload("Question").Take(&exercise).Error
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": "exercise not found",
		})
		return
	}

	userID := c.Request.Context().Value("user_id").(int)

	var answers []domain.Answer
	err = eh.db.Where("exercise_id = ? AND user_id = ?", id, userID).Take(&answers).Error
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": "no answer yet",
		})
		return
	}

	mapQA := make(map[int]domain.Answer)
	for _, answer := range answers {
		mapQA[answer.QuestionID] = answer
	}

	var score int
	for _, question := range exercise.Question {
		if strings.EqualFold(question.CorrectAnswer, mapQA[question.ID].Answer) {
			score += question.Score
		}
	}
	c.JSON(http.StatusOK, map[string]int{
		"score": score,
	})
}

// Function to create new question with exercise id as :id parameter
func (eh ExerciseHandler) CreateQuestion(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid exercise id",
		})
		return
	}

	var exercise domain.Exercise
	err = eh.db.Where("id = ?", id).Preload("Question").Take(&exercise).Error
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": "exercise not found",
		})
		return
	}

	var createQuestion domain.CreateQuestion
	if err := c.ShouldBind(&createQuestion); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid body",
		})
	}

	creatorID := c.Request.Context().Value("user_id").(int)

	question, err := domain.NewQuestion(id, createQuestion.Score, creatorID, createQuestion.Body, createQuestion.OptionA, createQuestion.OptionB, createQuestion.OptionC, createQuestion.OptionD, createQuestion.CorrectAnswer)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	if err := eh.db.Create(question).Error; err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]string{
		"message": "New question successfully created",
	})
}

// Function to create new answer with exercise id as :id and question id as :qid parameter
func (eh ExerciseHandler) CreateAnswer(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid id",
		})
		return
	}

	var exercise domain.Exercise
	err = eh.db.Where("id = ?", id).Preload("Question").Take(&exercise).Error
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": "exercise not found",
		})
		return
	}

	qidString := c.Param("qid")
	qid, err := strconv.Atoi(qidString)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid question id",
		})
		return
	}

	var question domain.Question
	err = eh.db.Where("id = ? AND exercise_id = ?", qid, id).Take(&question).Error
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": "question not found",
		})
		return
	}

	var createAnswer domain.CreateAnswer
	if err := c.ShouldBind(&createAnswer); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid body",
		})
	}

	userID := c.Request.Context().Value("user_id").(int)

	answer, err := domain.NewAnswer(id, qid, userID, createAnswer.Answer)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	if err := eh.db.Create(answer).Error; err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]string{
		"message": "New answer successfully created",
	})
}
