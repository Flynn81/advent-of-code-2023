package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeLine(t *testing.T) {
	assert.Equal(t, 12, decodeLine("1abc2"))
	assert.Equal(t, 38, decodeLine("pqr3stu8vwx"))
	assert.Equal(t, 15, decodeLine("a1b2c3d4e5f"))
	assert.Equal(t, 77, decodeLine("treb7uchet"))
	assert.Equal(t, 44, decodeLine("4271344"))
	assert.Equal(t, 44, decodeLine("42seven13four4"))
	assert.Equal(t, 38, decodeLine("3three2dfourfour58"))
	assert.Equal(t, 66, decodeLine("seven6"))
	assert.Equal(t, 44, decodeLine("eightthree44vxbjmvbpfleightvgbjxcgrjonesix"))
}

func TestDecodeSumValues(t *testing.T) {
	assert.Equal(t, 142, sumValues([]string{"1abc2","pqr3stu8vwx","a1b2c3d4e5f","treb7uchet"}))
}