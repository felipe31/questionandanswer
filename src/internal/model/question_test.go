package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TQuestionPersistance(assert *assert.Assertions, userId int64) {
	question := &Question{Title: "Testing question title", Body: "Testing question body", UserId: userId}

	question.ID = TAddQuestion(assert, question)
	TUpdateQuestion(assert, question)
	TQuestionByTitle(assert, question.Title)
	TQuestionById(assert, question.ID)
	TAnswerPersistance(assert, question.ID, userId)
	TRemoveQuestion(assert, question.ID)
}

func TAddQuestion(assert *assert.Assertions, question *Question) int64 {
	id, err := AddQuestion(*question)
	assert.Nil(err, fmt.Errorf("Error while adding new question!"))
	assert.Greater(int(id), 0,
		fmt.Errorf("AddAnswer returned an unexpected id (%v) for the question: %v", id, question))
	return id
}

func TUpdateQuestion(assert *assert.Assertions, question *Question) {
	err := UpdateQuestion(*question)
	assert.Nil(err, fmt.Errorf("Error while updating question!"))
}

func TQuestionByTitle(assert *assert.Assertions, title string) {
	question, err := QuestionByTitle(title)
	message := "Error while getting question by title!"
	assert.Nil(err, fmt.Errorf(message))
	assert.NotNil(question, fmt.Errorf(message))
}

func TQuestionById(assert *assert.Assertions, id int64) {
	question, err := QuestionById(int(id))
	message := "Error while getting question by id!"
	assert.Nil(err, fmt.Errorf(message))
	assert.NotNil(question, fmt.Errorf(message))
}

func TRemoveQuestion(assert *assert.Assertions, id int64) {
	err := RemoveQuestion(int(id))
	assert.Nil(err, fmt.Errorf("Error while removing a question!"))
}

func TestGetAllQuestions(t *testing.T) {
	assert := assert.New(t)
	MysqlConnect(testingHost)
	_, err := GetAllQuestions()
	message := "Error while getting all questions!"
	assert.Nil(err, fmt.Errorf(message))
}
