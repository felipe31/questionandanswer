package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TAddUser(assert *assert.Assertions, user *User) int64 {
	id, err := AddUser(*user)
	assert.Nil(err, fmt.Errorf("Error while adding new user!"))
	assert.Greater(int(id), 0,
		fmt.Errorf("AddUser returned an unexpected id (%v) for the user: %v", id, user))
	return id
}

func TUserByID(assert *assert.Assertions, id int64) {
	users, err := UserByID(int(id))
	message := "Error while getting user by id!"
	assert.Nil(err, fmt.Errorf(message))
	assert.NotNil(users, fmt.Errorf(message))
}

func TUserByUsername(assert *assert.Assertions, username string) {
	users, err := UserByUsername(username)
	message := "Error while getting user by username!"
	assert.Nil(err, fmt.Errorf(message))
	assert.NotNil(users, fmt.Errorf(message))
}

func TQuestionsByUserId(assert *assert.Assertions, id int64) {
	_, err := QuestionsByUserId(int(id))
	message := "Error while getting questions by user id!"
	assert.Nil(err, fmt.Errorf(message))
}

func TRemoveUser(assert *assert.Assertions, id int64) {
	err := RemoveUser(int(id))
	assert.Nil(err, fmt.Errorf("Error while removing an user!"))
}

func TestGetAllUsers(t *testing.T) {
	MysqlConnect(testingHost)
	assert := assert.New(t)
	_, err := GetAllUsers()
	message := "Error while getting all users!"
	assert.Nil(err, fmt.Errorf(message))
}
