package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuncs(t *testing.T) {
	assert.Equal(t, 1, extractFirstNumber("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"))
	assert.Equal(t, 3, extractFirstNumber(" 3 blue,"))
	assert.Equal(t, 4, extractFirstNumber("4 red;"))
	assert.Equal(t, 1, extractFirstNumber("; 1 red, 2 green, 6 blue; 2 green"))

	assert.Equal(t, 48, parseLine("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"))
	assert.Equal(t, 12, parseLine("Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"))
	assert.Equal(t, 1560, parseLine("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"))
	assert.Equal(t, 630, parseLine("Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"))
	assert.Equal(t, 36, parseLine("Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"))

	assert.Equal(t, 2286, sumPowers("test-input"))
}
