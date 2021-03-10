package main

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

func (num number) isSame(compared number) bool {
	return num.base == compared.base &&
		num.isNegative == compared.isNegative &&
		sameSlice(num.gte1Digits, compared.gte1Digits) &&
		sameSlice(num.lt1Digits, compared.lt1Digits)
}

func createInteger(integer int, base uint) (num number) {
	num.base = base
	if integer < 0 {
		num.isNegative = integer < 0
		integer *= -1
	}
	digits := make([]uint, 0)
	for integer != 0 {
		digits = append(digits, uint(integer%10))
		integer /= 10
	}
	num.gte1Digits = digits
	return
}
