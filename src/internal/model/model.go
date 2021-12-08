package model

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

type Question struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int64  `json:"user_id"`
}

type Answer struct {
	ID         int64  `json:"id"`
	Body       string `json:"body"`
	UserId     int64  `json:"user_id"`
	QuestionId int64  `json:"question_id"`
}

type QnA struct {
	Q Question `json:"question"`
	A Answer   `json:"answer"`
}
