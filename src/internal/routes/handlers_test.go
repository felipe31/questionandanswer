package routes

import (
	"felipesoares/questionandanswer/internal/model"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// I was not able to figure how to test POST methods
func TestHandleGetAllUsers(t *testing.T) {
	t.Parallel()
	model.MysqlConnect("127.0.0.1")
	assert := assert.New(t)

	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r = mux.SetURLVars(r, make(map[string]string))

	HandleGetAllUsers(w, r)

	assert.Equal(http.StatusOK, w.Code)
}
