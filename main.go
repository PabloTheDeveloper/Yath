package main

import "fmt"

type variable struct {
	letter   string
	exponent int
}

type number struct {
	base       uint
	isNegative bool
	gte1Digits []uint
	lt1Digits  []uint
}

func sameSlice(a []uint, b []uint) bool {
	if len(a) != len(b) {
		return false
	}
	for i, item := range a {
		if item != b[i] {
			return false
		}
	}
	return true
}

func (num number) isSame(compared number) error {
	switch {
	case num.base != compared.base:
		return nil
	case num.isNegative != compared.isNegative:
		return nil
	case !sameSlice(num.gte1Digits, compared.gte1Digits):
		return nil
	case !sameSlice(num.lt1Digits, compared.lt1Digits):
		return nil
	default:
		return nil
	}
}

func createNumber(integer int) (num number) {
	num.base = 10
	num.isNegative = integer < 0
	digits := make([]uint, 0)

	for integer != 0 {
		digits = append(digits, uint(integer%10))
		integer /= 10
	}
	num.gte1Digits = digits
	return
}

type term struct {
	coefficient number
	variables   []variable
}

type operand struct {
	constant  number
	variable  variable
	term      term
	operation operation
	kind      rune
}

func createConstant(num number) (oprd operand) {
	oprd.kind = 'c'
	oprd.constant = num
	return
}

type operation struct {
	left     *operand
	right    *operand
	operator rune
}

func main() {
	fmt.Println("a")
}
