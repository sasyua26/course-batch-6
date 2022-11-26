package domain

import (
	"errors"
	"time"
)

// Exercise struct for db
type Exercise struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Question    []Question `json:"question"`
}

// Question struct for db
type Question struct {
	ID            int       `json:"id"`
	ExerciseID    int       `json:"exercise_id"`
	Body          string    `json:"body"`
	OptionA       string    `json:"option_a"`
	OptionB       string    `json:"option_b"`
	OptionC       string    `json:"option_c"`
	OptionD       string    `json:"option_d"`
	CorrectAnswer string    `json:"-"`
	Score         int       `json:"score"`
	CreatorID     int       `json:"-"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// Answer struct for db
type Answer struct {
	ID         int       `json:"id"`
	ExerciseID int       `json:"exercise_id"`
	QuestionID int       `json:"question_id"`
	UserID     int       `json:"user_id"`
	Answer     string    `json:"answer"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// CreateExercise struct to create new exercise
type CreateExercise struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// CreateQuestion struct to create new question
type CreateQuestion struct {
	Body          string `json:"body"`
	OptionA       string `json:"option_a"`
	OptionB       string `json:"option_b"`
	OptionC       string `json:"option_c"`
	OptionD       string `json:"option_d"`
	CorrectAnswer string `json:"correct_answer"`
	Score         int    `json:"score"`
}

// CreateAnswer struct to create new answer
type CreateAnswer struct {
	Answer string `json:"answer"`
}

// Function to check new exercise input
func NewExercise(title, desc string) (*Exercise, error) {
	if title == "" {
		return nil, errors.New("Title is required")
	}

	if desc == "" {
		return nil, errors.New("Description is required")
	}

	return &Exercise{
		Title:       title,
		Description: desc,
	}, nil
}

// Function to check new question input
func NewQuestion(exerciseId, score, creatorID int, body, a, b, c, d, correctAnswer string) (*Question, error) {
	if score == 0 {
		return nil, errors.New("Score shoud be more than 0")
	}

	if body == "" {
		return nil, errors.New("Body is required")
	}

	if a == "" {
		return nil, errors.New("Option A is required")
	}

	if b == "" {
		return nil, errors.New("Option B is required")
	}

	if c == "" {
		return nil, errors.New("Option C is required")
	}

	if d == "" {
		return nil, errors.New("Option D is required")
	}

	if correctAnswer == "" {
		return nil, errors.New("Correct Answer is required")
	}

	return &Question{
		ExerciseID:    exerciseId,
		Body:          body,
		OptionA:       a,
		OptionB:       b,
		OptionC:       c,
		OptionD:       d,
		CorrectAnswer: correctAnswer,
		Score:         score,
		CreatorID:     creatorID,
	}, nil
}

// Function to check new answer input
func NewAnswer(exerciseId, questionID, userID int, answer string) (*Answer, error) {
	if answer == "" {
		return nil, errors.New("Answer is required")
	}

	return &Answer{
		ExerciseID: exerciseId,
		QuestionID: questionID,
		UserID:     userID,
		Answer:     answer,
	}, nil
}
