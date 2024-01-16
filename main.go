package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Please input filename.")
	}
	src, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	Lexer_src := Lexer(string(src))
	Ast := Parser(Lexer_src)
	vm := NewVM()
	vm.Run(Ast)

}
