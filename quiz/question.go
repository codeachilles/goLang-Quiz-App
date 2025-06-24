package quiz

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Question struct to hold question and answer
type Question struct {
	Text       string   `json:"text"`
	Options    []string `json:"options"`
	Answer     string   `json:"answer"`
	UserAnswer string
}

func AskQuestions(questions []Question) {
	reader := bufio.NewReader(os.Stdin)
	for i := range questions {
		fmt.Printf("Question #%d: %s\n", i+1, questions[i].Text)
		for j, option := range questions[i].Options {
			fmt.Printf("%c) %s\n", 'A'+j, option)
		}

		for {
			fmt.Print("Your answer (A, B, C, D): ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(strings.ToUpper(input))

			if len(input) == 1 && input[0] >= 'A' && input[0] < 'A'+byte(len(questions[i].Options)) {
				selectedIndex := int(input[0] - 'A')
				questions[i].UserAnswer = questions[i].Options[selectedIndex]
				break
			} else {
				fmt.Println("Invalid option. Please choose one of the available options.")
			}
		}
	}
}

func CalculateScore(questions []Question) int {
	score := 0
	for _, q := range questions {
		if strings.EqualFold(q.UserAnswer, q.Answer) {
			score++
		}
	}
	return score
}
