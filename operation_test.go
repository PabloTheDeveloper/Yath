package main

import "testing"

func TestAdd(t *testing.T) {
	numberOperationCases := []struct {
		desc     string
		a, b     number
		expected number
	}{
		{
			"12+13=25",
			createInteger(12, 10),
			createInteger(13, 10),
			createInteger(25, 10),
		},
		{
			"15+15=30",
			createInteger(15, 10),
			createInteger(15, 10),
			createInteger(30, 10),
		},
	}

	for _, tc := range numberOperationCases {
		t.Run(tc.desc, func(t *testing.T) {
			numOp := numberOperation{
				tc.a,
				tc.b,
				'+',
				number{},
			}
			numOp.add()
			if !numOp.result.isSame(tc.expected) {
				t.Fatalf("numbers don't match, %v, %v", numOp.result, tc.expected)
			}
		})
	}
}
