package games

import (
	"math/rand"
	"strconv"
	"strings"
)

type GeometricProgressionQuestion struct {
	question      string
	correctAnswer string
}

func (q *GeometricProgressionQuestion) GetText() string {
	return q.question
}

func (q *GeometricProgressionQuestion) GetAnswer() string {
	return q.correctAnswer
}

func CreateProgressionQuestion() Question {
	length := rand.Intn(6) + 5
	start := rand.Intn(10) + 1
	ratio := rand.Intn(3) + 2
	progression := make([]string, length)

	for i := 0; i < length; i++ {
		progression = append(progression, strconv.Itoa(start*pow(ratio, i+1)))
	}

	return &GeometricProgressionQuestion{
		question:      strings.Join(progression[:len(progression)-1], " "),
		correctAnswer: progression[len(progression)-1],
	}
}

func pow(start int, pow int) int {
	base := start
	for i := 1; i <= pow; i++ {
		base *= start
	}
	return base
}
