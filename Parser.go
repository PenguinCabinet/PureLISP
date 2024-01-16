package main

func Parser2(src []string) ast_t {
	if src[0] != "(" {
		return &ast_symbol_t{
			raw_data: src[0],
		}
	} else {

		Ast := &ast_call_t{}

		I := 1
		for ; I < len(src)-1; I++ {
			if src[I] == "(" {
				si := I
				nest := 0
				for ; I < len(src); I++ {
					if src[I] == "(" {
						nest += 1
					}
					if src[I] == ")" {
						nest -= 1
					}
					if nest == 0 {
						break
					}
				}
				Ast.data = append(Ast.data, Parser2(src[si:I+1]))
			} else {
				Ast.data = append(Ast.data, Parser2([]string{src[I]}))
			}
		}
		return Ast
	}
}

func Parser(src []string) ast_t {
	Ast := &ast_top_t{}

	I := 0
	for ; I < len(src); I++ {
		if src[I] == "(" {
			si := I
			nest := 0
			for ; I < len(src); I++ {
				if src[I] == "(" {
					nest += 1
				}
				if src[I] == ")" {
					nest -= 1
				}
				if nest == 0 {
					break
				}
			}
			Ast.data = append(Ast.data, Parser2(src[si:I+1]))
		} else {
			Ast.data = append(Ast.data, Parser2([]string{src[I]}))
		}
	}
	return Ast
}
