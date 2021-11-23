package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HandleNewQuestion(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleNewQuestion")
	var newQuestion Question

	reqBody, err := ioutil.ReadAll(r.Body)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Error while processing question body!", err.Error())
		return
	}
	json.Unmarshal(reqBody, &newQuestion)

	_, err = QuestionByTitle(newQuestion.Title)
	if err == nil {
		WriteError(w, http.StatusConflict, "Title already registered", "")
		return
	}

	newQuestion.ID, err = AddQuestion(newQuestion)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Error while creating question! ", err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newQuestion)
}

func HandleUpdateQuestion(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleUpdateQuestion")
	var updatedQuestion Question

	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Error while processing question body!", err.Error())
		return
	}
	json.Unmarshal(reqBody, &updatedQuestion)

	questionId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		WriteError(w, http.StatusNotFound, "Invalid question id!", err.Error())
		return
	}
	if int(updatedQuestion.ID) != questionId {
		WriteError(w, http.StatusBadRequest, "ID from URL does not match ID from body!", "")
		return
	}

	_, err = QuestionByTitle(updatedQuestion.Title)
	if err != nil {
		_, err = QuestionById(int(updatedQuestion.ID))
		if err != nil {
			WriteError(w, http.StatusConflict, "Question not found!", err.Error())
			return
		}
	}

	err = UpdateQuestion(updatedQuestion)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Error while updating question! ", err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(updatedQuestion)
}

func HandleDeleteQuestion(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleDeleteQuestion")

	questionIdVar := mux.Vars(r)["id"]
	w.Header().Set("Content-Type", "application/json")
	questionId, err := strconv.Atoi(questionIdVar)
	if err != nil {
		WriteError(w, http.StatusNotFound, "Invalid question id!", err.Error())
		return
	}

	question, err := QuestionById(questionId)
	if err != nil {
		WriteError(w, http.StatusNotFound, "Question to be deleted not found!", err.Error())
		return
	}
	err = RemoveQuestion(questionId)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Error while removing question!", err.Error())
		return
	}
	answer, err := AnswerByQuestionId(questionId)
	if err == nil {
		err = RemoveAnswer(int(answer.ID))
		if err != nil {
			WriteError(w, http.StatusInternalServerError, "Error while removing question's answer!", err.Error())
			return
		}
	}
	json.NewEncoder(w).Encode(question)

}

func HandleGetQuestionAndAnswer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleGetQuestionAndAnswer")
	questionIdVar := mux.Vars(r)["id"]
	w.Header().Set("Content-Type", "application/json")
	questionId, err := strconv.Atoi(questionIdVar)
	if err != nil {
		WriteError(w, http.StatusBadGateway, "Invalid question id!", err.Error())
		return
	}
	question, err := QuestionById(questionId)
	if err != nil {
		WriteError(w, http.StatusNotFound, "Question not found!", err.Error())
		return
	}

	answer, err := AnswerByQuestionId(questionId)
	if err != nil {
		fmt.Println(err.Error())
	}

	json.NewEncoder(w).Encode(QnA{question, answer})

}

func HandleGetAllQuestions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleGetAllQuestions")
	w.Header().Set("Content-Type", "application/json")
	questions, err := GetAllQuestions()
	if err != nil {
		WriteError(w, http.StatusNotFound, "Error while retrieving questions!", err.Error())
		return
	}
	if questions == nil {
		json.NewEncoder(w).Encode([]Question{})
		return
	}

	json.NewEncoder(w).Encode(questions)
}
