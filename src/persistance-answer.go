package main

import (
	"database/sql"
	"fmt"
)

func AddAnswer(answer Answer) (int64, error) {
	result, err := db.Exec("INSERT INTO Answers (body, user_fk, question_fk) VALUES (?, ?, ?)", answer.Body, answer.UserId, answer.QuestionId)
	if err != nil {
		return 0, fmt.Errorf("AddAnswer: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("AddAnswer: %v", err)
	}
	return id, nil
}

func RemoveAnswer(answerId int) error {
	_, err := db.Exec("DELETE FROM Answers WHERE id = ?", answerId)
	if err != nil {
		return fmt.Errorf("RemoveAnswer: %v", err)
	}
	return nil
}

// Updates one answer to the db. Cannot update question id
func UpdateAnswer(answer Answer) error {
	_, err := db.Exec("UPDATE Answers	SET body = ?, user_fk = ? WHERE id = ?", answer.Body, answer.UserId, answer.ID)
	if err != nil {
		return fmt.Errorf("UpdateAnswer: %v", err)
	}
	return nil
}

func AnswerByQuestionId(questionId int) (Answer, error) {
	var answer Answer

	row := db.QueryRow("SELECT * FROM Answers WHERE question_fk = ?", questionId)
	if err := row.Scan(&answer.ID, &answer.Body, &answer.UserId, &answer.QuestionId); err != nil {
		if err == sql.ErrNoRows {
			return answer, fmt.Errorf("AnswerByQuestionID %d: no such answer", questionId)
		}
		return answer, fmt.Errorf("AnswerByQuestionID %d: %v", questionId, err)
	}
	return answer, nil
}
