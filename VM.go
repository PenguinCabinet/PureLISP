package main

import "log"

type VM struct {
	Stack        []data_t
	Defined_data map[string]data_t
}

func str_to_atom(v string) *data_atom_t {
	if v == "nil" {
		return &data_atom_t{
			type_id:  data_t_nil,
			raw_data: "",
		}
	} else {
		return &data_atom_t{
			type_id:  data_t_val,
			raw_data: v,
		}
	}

}

func (self *VM) Run_func_atom(ast *ast_call_t) data_t {
	r := self.Run(ast.data[1])
	if !r.is_pair() {
		return &data_atom_t{
			type_id:  data_t_T,
			raw_data: "",
		}
	} else {
		return &data_atom_t{
			type_id:  data_t_nil,
			raw_data: "",
		}
	}
}

func (self *VM) Run_func_eq(ast *ast_call_t) data_t {
	a1 := self.Run(ast.data[1])
	a2 := self.Run(ast.data[2])
	if same_data_t(a1, a2) {
		return &data_atom_t{
			type_id:  data_t_T,
			raw_data: "",
		}
	} else {
		return &data_atom_t{
			type_id:  data_t_nil,
			raw_data: "",
		}
	}
}

func (self *VM) Run_func_car(ast *ast_call_t) data_t {
	a1 := self.Run(ast.data[1])
	if a1.is_pair() {
		temp := a1.(*data_pair_t)
		return temp.left
	} else {
		//ERROR
		return &data_atom_t{type_id: data_t_nil}
	}
}

func (self *VM) Run_func_cdr(ast *ast_call_t) data_t {
	a1 := self.Run(ast.data[1])
	if a1.is_pair() {
		temp := a1.(*data_pair_t)
		return temp.right
	} else {
		//ERROR
		return &data_atom_t{type_id: data_t_nil}
	}
}

func (self *VM) Run_func_cons(ast *ast_call_t) data_t {
	a1 := self.Run(ast.data[1])
	a2 := self.Run(ast.data[2])
	return &data_pair_t{
		left:  a1,
		right: a2,
	}
}

func (self *VM) Run_if(ast *ast_call_t) data_t {
	a1 := self.Run(ast.data[1])
	exp := a1.(*data_atom_t)

	if exp.type_id == data_t_T {
		return self.Run(ast.data[2])
	} else {
		return self.Run(ast.data[3])
	}
}

func Ast_to_Pair(Ast ast_t) data_t {
	if Ast.Get_type() == ast_type_call {
		ast := Ast.(*ast_call_t)
		A := &data_pair_t{
			left:  &data_atom_t{type_id: data_t_nil},
			right: &data_pair_t{},
		}
		temp := A
		for i, e := range ast.data {
			temp.left = Ast_to_Pair(e)
			if i != len(ast.data)-1 {
				temp.right = &data_pair_t{
					left:  &data_atom_t{type_id: data_t_nil},
					right: &data_pair_t{},
				}
				temp = temp.right.(*data_pair_t)
			} else {
				temp.right = &data_atom_t{type_id: data_t_nil}
			}
		}

		return A
	} else {
		ast := Ast.(*ast_symbol_t)
		return str_to_atom(ast.raw_data)
	}
}

func (self *VM) Run_quote(Ast *ast_call_t) data_t {

	return Ast_to_Pair(
		Ast.data[1],
	)
}

func (self *VM) Run_define(Ast *ast_call_t) data_t {
	(self.Defined_data)[Ast.data[1].(*ast_symbol_t).raw_data] = self.Run(Ast.data[2])
	return &data_atom_t{type_id: data_t_nil}
}

func (self *VM) Run_lambda(Ast *ast_call_t) data_t {
	return &data_lambda_t{
		ast: Ast,
	}
}

func (self *VM) Run(Ast ast_t) data_t {
	if Ast.Get_type() == ast_type_call {
		ast := Ast.(*ast_call_t)
		if ast.data[0].Get_type() != ast_type_call {
			call_func_name := ast.data[0].(*ast_symbol_t).raw_data

			if call_func_name == "atom" {
				return self.Run_func_atom(ast)
			} else if call_func_name == "eq" {
				return self.Run_func_eq(ast)
			} else if call_func_name == "car" {
				return self.Run_func_car(ast)
			} else if call_func_name == "cdr" {
				return self.Run_func_cdr(ast)
			} else if call_func_name == "cons" {
				return self.Run_func_cons(ast)
			} else if call_func_name == "if" {
				return self.Run_if(ast)
			} else if call_func_name == "quote" {
				return self.Run_quote(ast)
			} else if call_func_name == "define" {
				return self.Run_define(ast)
			} else if call_func_name == "lambda" {
				return self.Run_lambda(ast)
			} else if call_func_name == "print" {
				res := self.Run(ast.data[1])
				log.Printf("%v\n", res)
				return &data_atom_t{}
			} else {
				if Func_ast, ok := self.Defined_data[call_func_name]; ok {
					for i, e := range ast.data[1:] {
						self.Defined_data[Func_ast.(*data_lambda_t).ast.data[1].(*ast_call_t).data[i].(*ast_symbol_t).raw_data] = self.Run(e)
					}
					res := self.Run(Func_ast.(*data_lambda_t).ast.data[2])
					for i, _ := range ast.data[1:] {
						delete(self.Defined_data, Func_ast.(*data_lambda_t).ast.data[1].(*ast_call_t).data[i].(*ast_symbol_t).raw_data)
					}
					return res
				} else {
					//ERROR
					//未定義の関数を呼び出そうとしている
				}
			}

		} else {
			//ラムダ式を使用した即席関数
		}
	} else if Ast.Get_type() == ast_type_top {
		var res data_t
		ast := Ast.(*ast_top_t)
		for _, e := range ast.data {
			res = self.Run(e)
		}
		return res
	} else {
		ast := Ast.(*ast_symbol_t)
		if e, ok := self.Defined_data[ast.raw_data]; ok {
			return e
		}
		return str_to_atom(ast.raw_data)
	}

	return &data_atom_t{}
}

func NewVM() *VM {
	res := &VM{}
	res.Defined_data = map[string]data_t{}

	return res
}
