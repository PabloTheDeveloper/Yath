package main

import "fmt"
import "testing"

func TestSameSlice(t *testing.T) {
	cases := []struct {
		desc     string
		a, b     []uint
		expected bool
	}{
		{"TestNils", nil, nil, true},
		{"TestNil and NonNilSlice", nil, []uint{1}, false},
		{"TestNonNilSlice and Nil", []uint{1}, nil, false},
		{"TestNoneNilSlices", []uint{1}, []uint{1}, true},
		{"TestNoneNilSlices: len(A) < len(B)", []uint{1}, []uint{1, 2}, false},
		{"TestNoneNilSlices: len(A) > len(B)", []uint{1}, []uint{1, 2}, false},
		{"TestNoneNilSlices: A[0] > B[0]", []uint{100}, []uint{1}, false},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			actual := sameSlice(tc.a, tc.b)
			if actual != tc.expected {
				t.Fatalf("expected: %v got: %v for a: %v and b: %v", tc.expected, actual, tc.a, tc.b)
			}
		})
	}
}
func TestIsSame(t *testing.T) {
	// TODO
}
func TestCreateNumber(t *testing.T) {
	expected1 := number{10, false, []uint{0, 1}, []uint{}}
	num1 := createNumber(10)

	if num1.isSame(expected1) != nil {
		fmt.Println("not same")
	} else {
		fmt.Println("same")
	}
}
