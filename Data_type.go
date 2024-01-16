package main

type data_t interface {
	is_pair() bool
}

type atom_type_t int

const (
	data_t_nil atom_type_t = iota
	data_t_val
	data_t_T
	data_t_pair
	data_t_lambda
)

type data_pair_t struct {
	type_id atom_type_t
	left    data_t
	right   data_t
}

func (self *data_pair_t) is_pair() bool {
	return true
}

type data_atom_t struct {
	type_id  atom_type_t
	raw_data string
}

func (self *data_atom_t) is_pair() bool {
	return false
}

type data_lambda_t struct {
	ast *ast_call_t
	//arg_names []string
}

func (self *data_lambda_t) is_pair() bool {
	return false
}

func same_data_t(v1, v2 data_t) bool {
	if v1.is_pair() == v2.is_pair() {
		if v1.is_pair() {
			v1_temp := v1.(*data_pair_t)
			v2_temp := v2.(*data_pair_t)
			return same_data_t(v1_temp.left, v2_temp.left) && same_data_t(v1_temp.right, v2_temp.right)
		} else {
			v1_temp := v1.(*data_atom_t)
			v2_temp := v2.(*data_atom_t)
			return v1_temp.type_id == v2_temp.type_id && v1_temp.raw_data == v2_temp.raw_data
		}
	} else {
		return false
	}
}
