package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuncs(t *testing.T) {
	assert.Equal(t, true, isNumber("1"))
	assert.Equal(t, false, isNumber("."))
	assert.Equal(t, false, isNumber("+"))

	assert.Equal(t, 4361, sum("test-input"))
}