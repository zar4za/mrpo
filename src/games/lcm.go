package games

import (
	"fmt"
	"math/rand"
	"strconv"
)

type LcmQuestion struct {
	a   int
	b   int
	c   int
	lcm int
}

func (q *LcmQuestion) GetText() string {
	return fmt.Sprintf("%d %d %d", q.a, q.b, q.c)
}

func (q *LcmQuestion) GetAnswer() string {
	return strconv.Itoa(q.lcm)
}

func CreateLcmQuestion() Question {
	question := &LcmQuestion{
		a: rand.Intn(10) + 1,
		b: rand.Intn(10) + 1,
		c: rand.Intn(10) + 1,
	}

	question.lcm = lcm(question.a, question.b, question.c)
	return question
}

func lcm(a int, b int, c int) int {
	gcd := func(x, y int) int {
		for y != 0 {
			x, y = y, x%y
		}
		return x
	}

	lcmTwo := func(x, y int) int {
		return (x * y) / gcd(x, y)
	}

	return lcmTwo(lcmTwo(a, b), c)
}
