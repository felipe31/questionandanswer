package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	err := MysqlConnect("mysql")
	if err != nil {
		fmt.Printf("Error while connecting to the database!\n%v\n", err)
		return
	}

	// Set question related handlers
	router.HandleFunc("/question", HandleNewQuestion).Methods(http.MethodPost)
	router.HandleFunc("/question/{id}", HandleUpdateQuestion).Methods(http.MethodPut)
	router.HandleFunc("/question/{id}", HandleDeleteQuestion).Methods(http.MethodDelete)
	router.HandleFunc("/question/{id}", HandleGetQuestionAndAnswer).Methods(http.MethodGet)
	router.HandleFunc("/questions", HandleGetAllQuestions).Methods(http.MethodGet)

	// Set answer related handlers
	router.HandleFunc("/question/{id}/answer", HandleNewAnswer).Methods(http.MethodPost)
	router.HandleFunc("/question/{id}/answer", HandleUpdateAnswer).Methods(http.MethodPut)

	// Set user related handlers
	router.HandleFunc("/user", HandleNewUser).Methods(http.MethodPost)
	router.HandleFunc("/user/{id-username}", HandleDeleteUser).Methods(http.MethodDelete)
	router.HandleFunc("/user/{id-username}", HandleGetUser).Methods(http.MethodGet)
	router.HandleFunc("/user/{id-username}/questions", HandleGetUserQuestions).Methods(http.MethodGet)
	router.HandleFunc("/users", HandleGetAllUsers).Methods(http.MethodGet)

	router.NotFoundHandler = http.HandlerFunc(notFound)
	http.ListenAndServe(":8080", router)

}

func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Println("notFound")
	w.Header().Set("Content-Type", "application/json")
	WriteError(w, http.StatusNotFound, "Page not found!", r.URL.Path)
}
