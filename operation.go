package main

import "math"

type operation interface {
	add()
}

type expression struct {
	operands []operation
	result   []operation
	operator rune
}

type numberOperation struct {
	left     number
	right    number
	operator rune
	result   number
}

func (op *numberOperation) add() {
	// TODO add lt1Digits - change overflow
	// lt1Overflow := 0
	// assuming same base
	leftLen := len(op.left.gte1Digits)
	rightLen := len(op.right.gte1Digits)

	maxLen := int(math.Max(float64(leftLen), float64(rightLen)))
	gte1Digits := make([]uint, maxLen)
	gte1Overflow := uint(0)

	for i := range gte1Digits {
		if gte1Overflow > 0 {
			gte1Digits[i] += gte1Overflow
			gte1Overflow = 0
		}
		if i < leftLen {
			gte1Digits[i] += op.left.gte1Digits[i]
		}
		if i < rightLen {
			gte1Digits[i] += op.right.gte1Digits[i]
		}
		if gte1Digits[i] >= op.left.base {
			gte1Digits[i] -= op.left.base
			gte1Overflow += 1
		}
	}
	if gte1Digits[len(gte1Digits)-1] == 0 {
		gte1Digits = gte1Digits[0 : len(gte1Digits)-1]
	}
	op.result = number{10, false, gte1Digits, nil}
}
