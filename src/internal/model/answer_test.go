package model

import (
	"fmt"

	"github.com/stretchr/testify/assert"
)

func TAnswerPersistance(assert *assert.Assertions, questionId, userId int64) {
	answer := &Answer{Body: "Answer test body", UserId: userId, QuestionId: questionId}

	answer.ID = TAddAnswer(assert, answer)
	TAnswerByQuestionId(assert, answer.QuestionId)
	TUpdateAnswer(assert, answer)
	TRemoveAnswer(assert, answer.ID)
}

func TAddAnswer(assert *assert.Assertions, answer *Answer) int64 {
	id, err := AddAnswer(*answer)
	assert.Nil(err, fmt.Errorf("Error while adding new answer!"))
	assert.Greater(int(id), 0,
		fmt.Errorf("AddAnswer returned an unexpected id (%v) for the answer: %v", id, answer))
	return id
}

func TUpdateAnswer(assert *assert.Assertions, answer *Answer) {
	err := UpdateAnswer(*answer)
	assert.Nil(err, fmt.Errorf("Error while updating answer!"))
}

func TAnswerByQuestionId(assert *assert.Assertions, questionId int64) {
	answer, err := AnswerByQuestionId(int(questionId))
	message := "Error while getting answer by question id!"
	assert.Nil(err, fmt.Errorf(message))
	assert.NotNil(answer, fmt.Errorf(message))
}

func TRemoveAnswer(assert *assert.Assertions, id int64) {
	err := RemoveAnswer(int(id))
	assert.Nil(err, fmt.Errorf("Error while removing an answer!"))
}
