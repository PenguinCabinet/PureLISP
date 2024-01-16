package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexer(t *testing.T) {
	assert.Equal(
		t,
		[]string{"(", "A", "B", ")"},
		Lexer("( A B )"),
	)

	assert.Equal(
		t,
		[]string{"(", "A", "B", ")"},
		Lexer("(A B)"),
	)

	assert.Equal(
		t,
		[]string{"(", "A", "B", ")"},
		Lexer("(A\nB)"),
	)

	assert.Equal(
		t,
		[]string{"(", "A", "(", "(", "C", "A", ")", "D", "E", ")", ")"},
		Lexer("(A ((C A) D E))"),
	)
}
