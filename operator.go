package main

type operatorT rune

const (
	nop operatorT = iota
	add
	subtract
	multiply
	expontiate
	factorialize
)

type operator struct {
	operatorType operatorT
	result       operand
}
