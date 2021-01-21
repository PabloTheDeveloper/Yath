package main

import "fmt"

type variable struct {
	letter   string
	exponent int
}

type number struct {
	signedness bool
	digits     []byte
	base       byte
	decimalPos byte
}

type term struct {
	coefficient number
	variables   []variable
}

type operand struct {
	constant  *number
	variable  *variable
	term      *term
	operation *operation
	kind      rune
}

type operation struct {
	left     *operand
	right    *operand
	operator rune
}

func main() {
	fmt.Println("vim-go")
}
