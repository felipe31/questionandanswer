package main

import (
	"felipesoares/questionandanswer/internal/model"
	"felipesoares/questionandanswer/internal/routes"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	err := model.MysqlConnect("127.0.0.1")
	if err != nil {
		fmt.Printf("Error while connecting to the database!\n%v\n", err)
		return
	}

	// Set question related handlers
	router.HandleFunc("/question", routes.HandleNewQuestion).Methods(http.MethodPost)
	router.HandleFunc("/question/{id}", routes.HandleUpdateQuestion).Methods(http.MethodPut)
	router.HandleFunc("/question/{id}", routes.HandleDeleteQuestion).Methods(http.MethodDelete)
	router.HandleFunc("/question/{id}", routes.HandleGetQuestionAndAnswer).Methods(http.MethodGet)
	router.HandleFunc("/questions", routes.HandleGetAllQuestions).Methods(http.MethodGet)

	// Set answer related handlers
	router.HandleFunc("/question/{id}/answer", routes.HandleNewAnswer).Methods(http.MethodPost)
	router.HandleFunc("/question/{id}/answer", routes.HandleUpdateAnswer).Methods(http.MethodPut)

	// Set user related handlers
	router.HandleFunc("/user", routes.HandleNewUser).Methods(http.MethodPost)
	router.HandleFunc("/user/{id-username}", routes.HandleDeleteUser).Methods(http.MethodDelete)
	router.HandleFunc("/user/{id-username}", routes.HandleGetUser).Methods(http.MethodGet)
	router.HandleFunc("/user/{id-username}/questions", routes.HandleGetUserQuestions).Methods(http.MethodGet)
	router.HandleFunc("/users", routes.HandleGetAllUsers).Methods(http.MethodGet)

	router.NotFoundHandler = http.HandlerFunc(notFound)
	http.ListenAndServe(":8080", router)

}

func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Println("notFound")
	w.Header().Set("Content-Type", "application/json")
	routes.WriteError(w, http.StatusNotFound, "Page not found!", r.URL.Path)
}
