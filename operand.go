package main

type operandT int

const (
	integer operandT = iota
	float
	fraction
	uninitialized
)

type operand struct {
	operandType operandT
	quantify    []int
	variable    []string
	precision   uint
}
