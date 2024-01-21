package service

import "testing"

func TestFindPairForTarget(t *testing.T) {
	service := NewNumService()

	testCases := []struct {
		name     string
		a, b     []int
		target   int
		expected bool
	}{
		{
			name:     "Pair exists",
			a:        []int{1, 2, 3},
			b:        []int{1},
			target:   3,
			expected: true,
		},
		{
			name:     "Pair does nor exist",
			a:        []int{8, 9, 10},
			b:        []int{1},
			target:   3,
			expected: false,
		},
		{
			name:     "List are empty: result is false",
			a:        []int{},
			b:        []int{},
			target:   3,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := service.FindPairForTarget(tc.a, tc.b, tc.target)

			if tc.expected != actual {
				t.Fail()
			}
		})
	}
}
