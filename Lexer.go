package main

/*
(atom (cons A B))
*/

func Lexer(raw_src_a string) []string {
	raw_src := []rune(raw_src_a)

	var A []string
	var temp_stack []rune

	flash_temp_stack := func() {
		if len(temp_stack) > 0 {
			A = append(A, string(temp_stack))
		}
		temp_stack = []rune{}
	}

	for i := 0; i < len(raw_src); i++ {
		switch raw_src[i] {
		case '(':
			flash_temp_stack()
			A = append(A, string(raw_src[i]))
		case ')':
			flash_temp_stack()
			A = append(A, string(raw_src[i]))
		case ' ', '\t', '\n':
			flash_temp_stack()
		default:
			temp_stack = append(temp_stack, raw_src[i])
		}
	}
	flash_temp_stack()

	return A
}
