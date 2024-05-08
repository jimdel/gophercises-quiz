package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type quiz []question
type question struct {
	prompt string
	answer string
}

/**
- Read the quiz file
- Check where to read file (if arg present)
- Format the file to [] of T(question)
*/

func makeQuestion(str string) question {
	parts := strings.Split(str, ",")
	return question{
		prompt: parts[0],
		answer: parts[1],
	}
}

func makeQuiz() quiz {
	filepath := "problems.csv"
	contents, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	trimmed := strings.TrimSpace(string(contents))
	quizStr := strings.Split(trimmed, "\n")

	var quiz quiz
	for _, q := range quizStr {
		question := makeQuestion(q)
		quiz = append(quiz, question)
	}
	return quiz
}

func main() {
	fmt.Println("Starting...")
	q := makeQuiz()
	var score int
	for _, question := range q {
		fmt.Println(question.prompt)
		var resp string
		_, err := fmt.Scanln(&resp)
		if err != nil {
			log.Fatal(err)
		}
		if resp != question.answer {
			fmt.Println("Incorrect!")
		} else {
			score += 1
			fmt.Println("Correct!")
		}
	}

	// fmt.Sprint

	// for _, question := range q {
	// 	fmt.Println(question)
	// }

}
