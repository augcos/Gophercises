# Gophexercises
## Introduction
Gophexercises are a group of tutorial Golang exercises created by Jon Calhoun. You can visit his website [here](https://courses.calhoun.io). This repository consists in my implementation of the Gophexercises with additional modifications to improve the code.

## Development
All exercises were developed using Go v1.17.3.

## 1. Quiz Game
This exercise consists in the creation of a quiz game using a template of questions and answers provided in csv format. The program accepts two inputs: name of the csv file and time limit. The user will interact with the code using the terminal, and at the end of the quiz will get the final score. Some modifications were done on the base code, such as asking if the user wants to take part in the challenge before starting the quiz.
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

## 2. URL Shortener
This exercise consists in the creation of a URL shortener app. The app has a set of default (path, url) combinations, but you cand load more from a YAML or JSON file. To compile the code and run the binary, run from the URLShortener directory the following commands:
```
go build urlShortener.go
./urlShortener --filename [name of the .yaml or .json file]
```
If you prefer to directly run the code, you can use the command:
```
go run urlShortener.go --filename [name of the .yaml or .json file]
```
Remember to use the flags to overwrite the default settings (filename: url.yaml). The app runs on port 8080, so to access a shortened url go in your browser to localhost:8080/url.

## 3. Choose your own adventure
This extercises consisted in the creation of a choose your own adventure (cyoa) web app. Once you run the app, you'll be able to access it from your browser and advance throught different paths in the story depending on your decisions. To compile the code and run the binary, run from the ChooseAdventure directory the following commands:
```
go build storyApp.go
./storyApp --port [port in local host to access app] --filename [.json file with the story script]
```
If you prefer to directly run the code, you can use the command:
```
go run storyApp.go --port [port in local host to access app] --filename [.json file with the story script]
```
Remember to use the flags to overwrite the default settings (port: 3000, filename: adventureScript.json). The app is compatible with two HTML templates, so you can access the initial page either through localhost:port or localhost:port/story.

## 4. Link parser
This extercises consisted in the creation of a HTML link parser. The program takes some HTML code and prints all links it finds inside it. To compile the code and run the binary, run from the LinkParser directory the following commands:
```
go build ex1.go
./ex1
```
If you prefer to directly run the code, you can use the command:
```
go run ex1.go
```
You have three more examples in that same directory, you can run them changing the go filename in the command.