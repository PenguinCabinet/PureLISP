package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVM(t *testing.T) {
	assert.Equal(
		t,
		&data_atom_t{type_id: data_t_T, raw_data: ""},
		(NewVM()).Run(
			Parser(
				Lexer("( atom AA )"),
			),
		),
	)

	assert.Equal(
		t,
		&data_pair_t{
			left:  &data_atom_t{type_id: data_t_val, raw_data: "A"},
			right: &data_atom_t{type_id: data_t_val, raw_data: "B"},
		},
		(NewVM()).Run(
			Parser(
				Lexer("( cons A B )"),
			),
		),
	)

	assert.Equal(
		t,
		&data_pair_t{
			left: &data_atom_t{type_id: data_t_val, raw_data: "A"},
			right: &data_pair_t{
				left:  &data_atom_t{type_id: data_t_val, raw_data: "B"},
				right: &data_atom_t{type_id: data_t_val, raw_data: "C"},
			},
		},
		(NewVM()).Run(
			Parser(
				Lexer("( cons A (cons B C) )"),
			),
		),
	)

	assert.Equal(
		t,
		&data_atom_t{type_id: data_t_nil, raw_data: ""},
		(NewVM()).Run(
			Parser(
				Lexer("( atom (cons B C) )"),
			),
		),
	)

	assert.Equal(
		t,
		&data_atom_t{type_id: data_t_T, raw_data: ""},
		(NewVM()).Run(
			Parser(
				Lexer("( eq (cons B C) (cons B C) )"),
			),
		),
	)

	assert.Equal(
		t,
		&data_pair_t{
			left:  &data_atom_t{type_id: data_t_val, raw_data: "A"},
			right: &data_atom_t{type_id: data_t_val, raw_data: "B"},
		},
		(NewVM()).Run(
			Parser(
				Lexer("(car (cons (cons A B) (cons C D)) )"),
			),
		),
	)

	assert.Equal(
		t,
		&data_pair_t{
			left:  &data_atom_t{type_id: data_t_val, raw_data: "C"},
			right: &data_atom_t{type_id: data_t_val, raw_data: "D"},
		},
		(NewVM()).Run(
			Parser(
				Lexer("(cdr (cons (cons A B) (cons C D)) )"),
			),
		),
	)

	assert.Equal(
		t,
		&data_atom_t{type_id: data_t_val, raw_data: "A"},
		(NewVM()).Run(
			Parser(
				Lexer("(if (atom nil) A B)"),
			),
		),
	)

	assert.Equal(
		t,
		&data_atom_t{type_id: data_t_val, raw_data: "B"},
		(NewVM()).Run(
			Parser(
				Lexer("(if (atom (cons X Y)) A B)"),
			),
		),
	)

	assert.Equal(
		t,
		&data_atom_t{type_id: data_t_val, raw_data: "B"},
		(NewVM()).Run(
			Parser(
				Lexer("(if nil A B)"),
			),
		),
	)

	assert.Equal(
		t,
		&data_pair_t{
			left: &data_atom_t{type_id: data_t_val, raw_data: "A"},
			right: &data_pair_t{
				left: &data_atom_t{type_id: data_t_val, raw_data: "B"},
				right: &data_pair_t{
					left: &data_atom_t{type_id: data_t_val, raw_data: "C"},
					right: &data_pair_t{
						left:  &data_atom_t{type_id: data_t_val, raw_data: "D"},
						right: &data_atom_t{type_id: data_t_nil},
					},
				},
			},
		},
		(NewVM()).Run(
			Parser(
				Lexer("(quote (A B C D))"),
			),
		),
	)
	assert.Equal(
		t,
		&data_pair_t{
			left: &data_atom_t{type_id: data_t_val, raw_data: "A"},
			right: &data_pair_t{
				left: &data_atom_t{type_id: data_t_val, raw_data: "B"},
				right: &data_pair_t{
					left: &data_pair_t{
						left: &data_atom_t{type_id: data_t_val, raw_data: "C"},
						right: &data_pair_t{
							left:  &data_atom_t{type_id: data_t_val, raw_data: "D"},
							right: &data_atom_t{type_id: data_t_nil},
						},
					},
					right: &data_atom_t{type_id: data_t_nil},
				},
			},
		},
		(NewVM()).Run(
			Parser(
				Lexer("(quote (A B (C D)))"),
			),
		),
	)

	assert.Equal(
		t,
		&data_atom_t{type_id: data_t_nil},
		(NewVM()).Run(
			Parser(
				Lexer("(define A 1)"),
			),
		),
	)

	assert.Equal(
		t,
		&data_lambda_t{
			ast: &ast_call_t{
				data: []ast_t{
					&ast_symbol_t{raw_data: "lambda"},
					&ast_call_t{
						data: []ast_t{
							&ast_symbol_t{raw_data: "arg1"},
						},
					},
					&ast_call_t{
						data: []ast_t{
							&ast_symbol_t{raw_data: "cons"},
							&ast_symbol_t{raw_data: "arg1"},
							&ast_symbol_t{raw_data: "A"},
						},
					},
				},
			},
		},
		(NewVM()).Run(
			Parser(
				Lexer("(lambda (arg1) (cons arg1 A))"),
			),
		),
	)

	assert.Equal(
		t,
		&data_atom_t{type_id: data_t_nil},
		(NewVM()).Run(
			Parser(
				Lexer("(print HelloWorld)"),
			),
		),
	)

	assert.Equal(
		t,
		&data_atom_t{type_id: data_t_nil},
		(NewVM()).Run(
			Parser(
				Lexer("(define str HelloWorld)(print str)"),
			),
		),
	)

	assert.Equal(
		t,
		&data_atom_t{type_id: data_t_T},
		(NewVM()).Run(
			Parser(
				Lexer("(define f (lambda (arg1 arg2) (eq arg1 arg2)))(f A A)"),
			),
		),
	)

	assert.Equal(
		t,
		&data_atom_t{type_id: data_t_nil},
		(NewVM()).Run(
			Parser(
				Lexer("(define f (lambda (arg1 arg2) (eq arg1 arg2)))(f A B)"),
			),
		),
	)

	assert.Equal(
		t,
		&data_atom_t{type_id: data_t_nil},
		(NewVM()).Run(
			Parser(
				Lexer("(define f (lambda (arg1) (print arg1)))(f HelloWorld)"),
			),
		),
	)

	assert.Equal(
		t,
		&data_atom_t{type_id: data_t_nil},
		(NewVM()).Run(
			Parser(
				Lexer("(define f2 (lambda (arg2) (print arg2)))(define f1 (lambda (arg1) (f2 arg1)))(f1 HelloWorld)"),
			),
		),
	)

	assert.Equal(
		t,
		&data_atom_t{type_id: data_t_nil},
		(NewVM()).Run(
			Parser(
				Lexer("(define f2 (lambda (arg2) (print arg1)))(define f1 (lambda (arg1) (f2 arg1)))(f1 HelloWorld)"),
			),
		),
	)

}
