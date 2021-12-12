package main

// Package imports
import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// we define the flags for the csv file and the time limit for answering the quiz
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")

	// we parse the flags and open the csv (exit if error)
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}

	// we read the data from the csv (exit if error)
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.\n")
	}

	// we ask if the user wants to accept the challenge
	fmt.Printf("This quiz will use the %s file and have a time limit of %d seconds\n", *csvFilename, *timeLimit)
	repeat := 0
	for {
		fmt.Println("Do you want to accept the challenge? [Yes/No]")
		var chAccepted string
		fmt.Scanf("%s\n", &chAccepted)
		if chAccepted == "Yes" {
			break
		} else if chAccepted == "No" {
			exit("Maybe next time...\n")
		} else {
			if repeat == 1 {
				exit("Please introduce a valid answer... Next time.\n")
			}
			fmt.Println("Please introduce a valid answer.")
			repeat++
		}

	}

	// we parse the data from the csv to problem type structs and define the timer channel
	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0

	// we define problemLoop to break out of the loop when time expires
problemLoop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q) // We print the question
		answerCh := make(chan string)                // We start a new channel to capture the answer
		go func() {                                  // We use an anonymous function to scan for the
			var answer string // answer and pass it to the answerCh channel
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		// we start a select structure
		select {
		case <-timer.C: // if the time expires we break out of the loop
			break problemLoop
		case answer := <-answerCh: // if the user writes an answer, we evaluate if
			if answer == p.a { // it is correct and update the number of correct
				correct++ // answers accordingly
			}
		}
	}

	// We print out the final score. Have fun!
	fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))

}

// Problem struct: contains a question (q) and an answer (a) field
type problem struct {
	q string
	a string
}

// parseLines function: takes the lines from the csv and parses them into an array of problem type structs
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

// exit function: prints an exit message en finished the execution
func exit(msg string) {
	fmt.Print(msg)
	os.Exit(1)
}
