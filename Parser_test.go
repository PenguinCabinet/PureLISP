package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	assert.Equal(
		t,
		&ast_top_t{
			data: []ast_t{
				&ast_call_t{
					data: []ast_t{
						&ast_symbol_t{raw_data: "A"},
						&ast_symbol_t{raw_data: "B"},
					},
				},
			},
		},
		Parser([]string{"(", "A", "B", ")"}),
	)

	assert.Equal(
		t,
		&ast_top_t{
			data: []ast_t{
				&ast_call_t{

					data: []ast_t{
						&ast_symbol_t{raw_data: "A"},
						&ast_symbol_t{raw_data: "B"},
						&ast_symbol_t{raw_data: "C"},
					},
				},
			},
		},
		Parser([]string{"(", "A", "B", "C", ")"}),
	)
	//log.Println("OKKKKKKK!!!!")

	assert.Equal(
		t,
		&ast_top_t{
			data: []ast_t{
				&ast_call_t{
					data: []ast_t{
						&ast_symbol_t{raw_data: "A"},
						&ast_call_t{
							data: []ast_t{
								&ast_call_t{
									data: []ast_t{
										&ast_symbol_t{raw_data: "B"},
										&ast_symbol_t{raw_data: "C"},
									},
								},
								&ast_symbol_t{raw_data: "D"},
							},
						},
					},
				},
			},
		},
		Parser([]string{"(", "A", "(", "(", "B", "C", ")", "D", ")", ")"}),
	)

	assert.Equal(
		t,
		&ast_top_t{
			data: []ast_t{
				&ast_call_t{
					data: []ast_t{
						&ast_symbol_t{raw_data: "A"},
						&ast_call_t{
							data: []ast_t{
								&ast_symbol_t{raw_data: "B"},
								&ast_call_t{
									data: []ast_t{
										&ast_symbol_t{raw_data: "C"},
										&ast_symbol_t{raw_data: "D"},
									},
								},
								&ast_symbol_t{raw_data: "E"},
							},
						},
					},
				},
			},
		},
		Parser([]string{"(", "A", "(", "B", "(", "C", "D", ")", "E", ")", ")"}),
	)

	assert.Equal(
		t,
		&ast_top_t{
			data: []ast_t{
				&ast_call_t{
					data: []ast_t{
						&ast_call_t{
							data: []ast_t{
								&ast_symbol_t{raw_data: "B"},
								&ast_call_t{
									data: []ast_t{
										&ast_symbol_t{raw_data: "C"},
										&ast_symbol_t{raw_data: "D"},
									},
								},
								&ast_symbol_t{raw_data: "E"},
							},
						},
					},
				},
			},
		},
		Parser([]string{"(", "(", "B", "(", "C", "D", ")", "E", ")", ")"}),
	)

	assert.Equal(
		t,
		&ast_top_t{
			data: []ast_t{
				&ast_call_t{
					data: []ast_t{
						&ast_symbol_t{raw_data: "A"},
						&ast_symbol_t{raw_data: "B"},
					},
				},
				&ast_call_t{
					data: []ast_t{
						&ast_symbol_t{raw_data: "A"},
						&ast_symbol_t{raw_data: "B"},
					},
				},
			},
		},
		Parser([]string{"(", "A", "B", ")", "(", "A", "B", ")"}),
	)

}
