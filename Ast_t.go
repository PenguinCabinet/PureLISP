package main

type ast_type_t int

const (
	ast_type_top ast_type_t = iota
	ast_type_call
	ast_type_symbol
)

type ast_t interface {
	Get_type() ast_type_t
}

type ast_top_t struct {
	data []ast_t
}

func (self *ast_top_t) Get_type() ast_type_t {
	return ast_type_top
}

type ast_call_t struct {
	data []ast_t
}

func (self *ast_call_t) Get_type() ast_type_t {
	return ast_type_call
}

type ast_symbol_t struct {
	raw_data string
}

func (self *ast_symbol_t) Get_type() ast_type_t {
	return ast_type_symbol
}
