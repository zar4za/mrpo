package games

import "fmt"

type Question interface {
	GetText() string
	GetAnswer() string
}

func Run(name string, maxTries int, createQuestion func() Question) {
	for i := 0; i < maxTries; i++ {
		question := createQuestion()

		fmt.Printf("Question %d: %s\n", i+1, question.GetText())
		var answer string
		fmt.Print("Your answer: ")
		fmt.Scanln(&answer)

		if answer == question.GetAnswer() {
			fmt.Println("Correct!")
		} else {
			fmt.Printf("%s is wrong answer. The correct was %s\n", answer, question.GetAnswer())
			fmt.Printf("Let's try again, %s\n", name)
		}
	}

	fmt.Printf("Congratulations, %s", name)
}
