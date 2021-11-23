package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Answer related handlers
func HandleNewAnswer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleNewAnswer")

	var newAnswer Answer

	reqBody, err := ioutil.ReadAll(r.Body)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Error while processing answer body!", err.Error())
		return
	}
	json.Unmarshal(reqBody, &newAnswer)
	_, err = QuestionById(int(newAnswer.QuestionId))
	if err != nil {
		WriteError(w, http.StatusNotFound, "Question not found.", err.Error())
		return
	}
	_, err = AnswerByQuestionId(int(newAnswer.QuestionId))
	if err == nil {
		WriteError(w, http.StatusConflict, "This question already has an answer. Try updating it.", "")
		return
	}

	newAnswer.ID, err = AddAnswer(newAnswer)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Error while creating answer! ", err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newAnswer)
}

func HandleUpdateAnswer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleUpdateAnswer")
	var updatedAnswer Answer

	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Error while processing answer body!", err.Error())
		return
	}
	json.Unmarshal(reqBody, &updatedAnswer)

	questionId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		WriteError(w, http.StatusNotFound, "Invalid question id!", err.Error())
		return
	}
	if int(updatedAnswer.QuestionId) != questionId {
		WriteError(w, http.StatusBadRequest, "ID from URL does not match ID from body!", "")
		return
	}

	_, err = QuestionById(int(updatedAnswer.QuestionId))
	if err != nil {
		WriteError(w, http.StatusConflict, "Question not found!", err.Error())
		return
	}
	_, err = AnswerByQuestionId(int(updatedAnswer.QuestionId))
	if err != nil {
		WriteError(w, http.StatusConflict, "Answer not found!", err.Error())
		return
	}
	err = UpdateAnswer(updatedAnswer)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Error while updating answer! ", err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(updatedAnswer)
}
