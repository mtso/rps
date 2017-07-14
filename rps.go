// Package rps is a random choice picker.
//
// Defaults to {"rock", "paper", "scissors"},
// but can supply a slice of strings to choose
// from in Roll() or use Set() to change the
// choices for all subsequent calls to Roll().
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

// Random number generator.
var source = rand.NewSource(time.Now().Unix())
var rng = rand.New(source)

func randomInt(end int) int {
	num := rng.Int()
	return num % end
}

// Roll returns a random choice from an array.
// (Defaults to {"rock", "paper", "scissors"}.
// Optionally pass a string slice to pick from.
func Roll(custom ...[]string) string {
	if len(custom) > 0 {
		choices := custom[0]
		choice := randomInt(len(choices))
		return choices[choice]
	}

	choice := randomInt(choicelen)
	return choices[choice]
}

// Set changes the choices for subsequent Roll() calls.
func Set(newChoices []string) {
	choices = newChoices
	choicelen = len(newChoices)
}
