package main

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
				t.Fatalf("expected: %v got: %v for a: %v and b: %v",
					tc.expected, actual, tc.a, tc.b)
			}
		})
	}
}

func TestIsSame(t *testing.T) {
	cases := []struct {
		desc          string
		num, compared number
		expected      bool
	}{
		{"defaultNumberComparison", number{}, number{}, true},
		{
			"NumberComparisonDiffentBases",
			number{10, false, []uint{1}, []uint{0}},
			number{11, false, []uint{1}, []uint{0}},
			false,
		},
		{
			"NumberComparisonDiffentBasesSwitched",
			number{11, false, []uint{1}, []uint{0}},
			number{10, false, []uint{1}, []uint{0}},
			false,
		},
		{
			"NumberComparisonDiffentSigns",
			number{10, true, []uint{1}, []uint{0}},
			number{10, false, []uint{1}, []uint{0}},
			false,
		},
		{
			"NumberComparisonDiffentSignsSwitched",
			number{11, false, []uint{1}, []uint{0}},
			number{10, true, []uint{1}, []uint{0}},
			false,
		},
		{
			"SameNumber",
			number{10, false, []uint{1}, []uint{0}},
			number{10, true, []uint{1}, []uint{0}},
			false,
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			actual := tc.num.isSame(tc.compared)
			if actual != tc.expected {
				t.Fatalf("expected: %v got: %v for num: %v and compared: %v",
					tc.expected, actual, tc.num, tc.compared)
			}
		})
	}
}

func TestCreateInteger(t *testing.T) {
	cases := []struct {
		desc           string
		integer        int
		base           uint
		expected       number
		isSameExpected bool
	}{
		{"test zero", 0, 10, number{10, false, []uint{}, nil}, true},
		{"SameBaseZero", 0, 2, number{2, false, []uint{}, nil}, true},
		{"DifferentBaseZero", 0, 2, number{1, false, []uint{}, nil}, false},
		{"NegativeOne", -1, 10, number{10, true, []uint{1}, nil}, true},
		{"NegativeMultDigit", -201, 10, number{10, true, []uint{1, 0, 2}, nil}, true},
		{"PositiveMultDigit", 1020, 10, number{10, false, []uint{0, 2, 0, 1}, nil}, true},
		{"DifferentNegativeMultDigit", -1010, 10, number{10, true, []uint{1, 0, 1}, nil}, false},
		{"DifferentPositiveMultDigit", 1010, 10, number{10, false, []uint{1, 0, 1}, nil}, false},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			actual := createInteger(tc.integer, tc.base)
			if tc.isSameExpected != actual.isSame(tc.expected) {
				t.Fatalf("suppose to match: %v\n"+
					"expected: %v got: %v for integer: %v and base: %v",
					tc.isSameExpected, tc.expected, actual, tc.integer, tc.base)
			}
		})
	}
}
