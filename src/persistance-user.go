package main

import (
	"database/sql"
	"fmt"
)

func AddUser(user User) (int64, error) {
	result, err := db.Exec("INSERT INTO Users (username) VALUES (?)", user.Username)
	if err != nil {
		return 0, fmt.Errorf("addUser: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addUser: %v", err)
	}
	return id, nil
}

func RemoveUser(userId int) error {
	_, err := db.Exec("DELETE FROM Users WHERE id = ?", userId)
	if err != nil {
		return fmt.Errorf("RemoveUser: %v", err)
	}
	return nil
}

func GetAllUsers() ([]User, error) {

	var users []User

	rows, err := db.Query("SELECT * FROM Users")
	if err != nil {
		return nil, fmt.Errorf("GetAllUsers: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username); err != nil {
			return nil, fmt.Errorf("GetAllUsers: %v", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAllUsers: %v", err)
	}
	return users, nil
}

func UserByID(userId int) (User, error) {
	var user User

	row := db.QueryRow("SELECT * FROM Users WHERE id = ?", userId)
	if err := row.Scan(&user.ID, &user.Username); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("userByID %d: no such user", userId)
		}
		return user, fmt.Errorf("userByID %d: %v", userId, err)
	}
	return user, nil
}

func UserByUsername(username string) (User, error) {
	var user User

	row := db.QueryRow("SELECT * FROM Users WHERE username = ?", username)
	if err := row.Scan(&user.ID, &user.Username); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("UserByUsername %v: no such user", username)
		}
		return user, fmt.Errorf("UserByUsername %v: %v", username, err)
	}
	return user, nil
}

func QuestionsByUserId(userId int) ([]Question, error) {

	var questions []Question

	rows, err := db.Query("SELECT * FROM Questions WHERE user_fk = ?", userId)
	if err != nil {
		return nil, fmt.Errorf("questionsByUserId %q: %v", userId, err)
	}
	defer rows.Close()

	for rows.Next() {
		var question Question
		if err := rows.Scan(&question.ID, &question.Title, &question.Body, &question.UserId); err != nil {
			return nil, fmt.Errorf("questionsByUserId %q: %v", userId, err)
		}
		questions = append(questions, question)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("questionsByUserId %q: %v", userId, err)
	}
	return questions, nil
}
