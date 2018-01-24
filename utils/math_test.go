package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	fmt.Println("Tokens =", Tokens)
	fmt.Println("length of Tokens is", Length)
}
func TestIdToString(t *testing.T) {
	id := 72
	expectvalue := "1a"
	assert.Equal(t, expectvalue, IdToString(id))
}
func TestStringToId(t *testing.T) {
	s := "a"
	expectvalue := 10
	assert.Equal(t, expectvalue, StringToId(s))
}
