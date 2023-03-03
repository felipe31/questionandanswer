# Question And Answer ![GitHub](https://img.shields.io/github/license/Felipe31/questionandanswer?style=flat-square)

This is an exercise. This is a backend side of a system called QuestionsAndAnswers. It is similar to Quora/Stackoverflow and others with 1 major difference. It only allow 1 answer per question. If someone thinks they have a better answer, they will have to update the existing answer for that particular question instead of adding another answer.

## Install
```sh
docker-compose up -d
```
Once running, the application will be available on `localhost:8080`.
## API
### Organization
The data used in this API is organized as follows:

#### Question
```json
{
  "id": 1,
  "title": "How to exit from vim?",
  "body": "I hear Vim is a very respected editor. So I was willing to try it out, but not I am not able to exit from it. And if I kill the process, the content of the file gets lost. What should I do?",
  "user_id": 1
}
```
#### Answer
```json
{
  "id": 1,
  "body": "In order to exit Vim, type ':q'. If you want to save before exiting, type ':wq'.!",
  "user_id": 1,
  "question_id": 1
}
```
#### User
```json
{
    "id": 1,
    "username": "user"
}
```

#### QnA
```json
{
  "question": {
    "id": 1,
    "title": "How to exit from vim??",
    "body": "I hear Vim is a very respected editor. So I was willing to try it out, but not I am not able to exit from it. And if I kill the process, the content of the file gets lost. What should I do?",
    "user_id": 1
  },
  "answer": {
    "id": 1,
    "body": "In order to exit Vim, type ':q'. If you want to save before exiting, type ':wq'.",
    "user_id": 1,
    "question_id": 1
  }
}
```

### Endpoints
There are endpoints availables for questions, answers and users as follows:

##### Question endpoints:
```
POST /question
PUT /question/{id}
DELETE /question/{id}
GET /question/{id}
GET /questions
```

##### Answer endpoints:
```
POST /question/{id}/answer
PUT /question/{id}/answer
```

##### User endpoints:
```
POST /user
DELETE /user/{id-username}
GET /user/{id-username}
GET /user/{id-username}/questions
GET /users
```
## Testing
This API can be tested automatically using `go test` or manually, by using the [Postman](https://github.com/Felipe31/questionandanswer/blob/main/Go%20QuestionAndAnswer.com.postman_collection.json) file (which contains one test case for each route).
Note: `go test` does **not** test the route handlers yet. 

## Things to add
- Get answers by user: this feature could help checking who is more active in the community
- Test each route handler.

---
Note: This project is for learning/testing. It is not taking any security measure. All the important data (as db username and password) is in plain text. 
