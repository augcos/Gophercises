# Gophexercises
## Introduction
Gophexercises are a group of tutorial Golang exercises created by Jon Calhoun. You can visit his website [here](https://courses.calhoun.io)

## Development
All exercises were developed using Go v1.17.3.

## 1. Quiz Game
This exercise consisted in the creation of a quiz game using a template of questions and answers provided in csv format. The program accepts two inputs: name of the csv file and time limit. The user will interact with the code using the terminal, and at the end of the quiz will get the final score. Some modifications were done on the base code, such as asking if the user wants to take part in the challenge before starting the quiz.
To compile the code and run the binary, run from the QuizGame directory the following commands:
```
go build quizgame.go
./quizgame --csv [filename.csv]  --limit [time limit in seconds]
```
If you prefer to directly run the code, you can use the command:
```
go run quizgame.go --csv [filename.csv]  --limit [time limit in seconds]
```
Remember to use the flags to overwrite the default settings (filename: problems.csv, time limit: 30 seconds).