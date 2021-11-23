package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// I was not able to figure how to test POST methods
func TestHandleGetAllUsers(t *testing.T) {
	t.Parallel()
	MysqlConnect(testingHost)
	assert := assert.New(t)

	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r = mux.SetURLVars(r, make(map[string]string))

	HandleGetAllUsers(w, r)

	assert.Equal(http.StatusOK, w.Code)
}
