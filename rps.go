package rps

import (
	"math/rand"
	"time"
)

var defaultChoices = []string{
	"rock",
	"paper",
	"scissors",
}

// Default choices {"rock", "paper", "scissors"}.
var choices = defaultChoices

var choicelen = len(choices)

var source = rand.NewSource(time.Now().Unix())

// Random number generator.
var rng = rand.New(source)

func randomInt(end int) int {
	num := rng.Int()
	return num % end
}

func Roll(custom ...[]string) string {
	if len(custom) > 0 {
		choices := custom[0]
		choice := randomInt(len(choices))
		return choices[choice]
	}

	choice := randomInt(choicelen)
	return choices[choice]
}

func Set(inputchoices []string) {
	choices = inputchoices
	choicelen = len(inputchoices)
}
