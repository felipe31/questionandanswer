package main

import (
	"database/sql"
	"fmt"
)

// Adds one question to the db. Id field is not expected
func AddQuestion(question Question) (int64, error) {
	result, err := db.Exec("INSERT INTO Questions (title, body, user_fk) VALUES (?, ?, ?)", question.Title, question.Body, question.UserId)
	if err != nil {
		return 0, fmt.Errorf("addQuestion: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addQuestion: %v", err)
	}
	return id, nil
}

// Updates one question to the db
func UpdateQuestion(question Question) error {
	_, err := db.Exec("UPDATE Questions	SET title = ?, body = ?, user_fk = ? WHERE id = ?", question.Title, question.Body, question.UserId, question.ID)
	if err != nil {
		return fmt.Errorf("UpdateQuestion: %v", err)
	}
	return nil
}

func RemoveQuestion(questionId int) error {
	_, err := db.Exec("DELETE FROM Questions WHERE id = ?", questionId)
	if err != nil {
		return fmt.Errorf("RemoveQuestion: %v", err)
	}
	return nil
}

// Retrives all existing questions from the db
func GetAllQuestions() ([]Question, error) {
	var questions []Question

	rows, err := db.Query("SELECT * FROM Questions")
	if err != nil {
		return nil, fmt.Errorf("GetAllQuestions: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var question Question
		if err := rows.Scan(&question.ID, &question.Title, &question.Body, &question.UserId); err != nil {
			return nil, fmt.Errorf("GetAllQuestions: %v", err)
		}
		questions = append(questions, question)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAllQuestions: %v", err)
	}
	return questions, nil
}

// Retrives exactly one question from the db, by the title parameter
func QuestionByTitle(title string) (Question, error) {
	var question Question

	row := db.QueryRow("SELECT * FROM Questions WHERE title = ?", title)
	if err := row.Scan(&question.ID, &question.Title, &question.Body, &question.UserId); err != nil {
		if err == sql.ErrNoRows {
			return question, fmt.Errorf("QuestionByTitle %v: no such question", title)
		}
		return question, fmt.Errorf("QuestionByTitle %v: %v", title, err)
	}
	return question, nil
}

// Retrives exactly one question from the db, by the id parameter
func QuestionById(questionId int) (Question, error) {
	var question Question

	row := db.QueryRow("SELECT * FROM Questions WHERE id = ?", questionId)
	if err := row.Scan(&question.ID, &question.Title, &question.Body, &question.UserId); err != nil {
		if err == sql.ErrNoRows {
			return question, fmt.Errorf("QuestionById %d: no such question", questionId)
		}
		return question, fmt.Errorf("QuestionById %d: %v", questionId, err)
	}
	return question, nil
}
