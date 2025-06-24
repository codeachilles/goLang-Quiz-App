package main

import (
	"encoding/json"
	"fmt"
	"goLang_quiz_app/quiz"
	"os"
)

func main() {
	fmt.Println("Welcome to the Go Quiz Application!")

	// Read and parse questions from JSON file
	questions, err := loadQuestions("questions.json")
	if err != nil {
		fmt.Printf("Error loading questions: %v\n", err)
		return
	}
	quiz.AskQuestions(questions)
	score := quiz.CalculateScore(questions)

	fmt.Printf("\n--- Quiz Summary ---\n")
	for i, q := range questions {
		fmt.Printf("Q%d: %s\n", i+1, q.Text)
		fmt.Printf("Your Answer: %s\n", q.UserAnswer)
		fmt.Printf("Correct Answer: %s\n", q.Answer)
		fmt.Println("--------------------")
	}
	fmt.Printf("Your final score is: %d out of %d\n", score, len(questions))
}

func loadQuestions(file string) ([]quiz.Question, error) {
	var questions []quiz.Question

	// Read the file content
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Unmarshal the JSON to the questions slice
	if err := json.Unmarshal(content, &questions); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return questions, nil
}
