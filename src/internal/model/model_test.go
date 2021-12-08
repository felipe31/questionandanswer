package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPersistance(t *testing.T) {
	MysqlConnect(testingHost)
	assert := assert.New(t)
	user := &User{Username: "test1"}

	user.ID = TAddUser(assert, user)
	TUserByID(assert, user.ID)
	TUserByUsername(assert, user.Username)
	TQuestionsByUserId(assert, user.ID)
	TQuestionPersistance(assert, user.ID)
	TRemoveUser(assert, user.ID)

}
