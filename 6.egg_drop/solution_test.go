package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

type testCase struct{
	name string
	expected int
	k int
	n int
}
func TestSolution(t *testing.T){
	test := []testCase{
		{
			name: "10 floors 1 egg",
			expected: 10,
			n: 10,
			k: 1,
		},
		{
			name: "100 floors 1 egg",
			expected: 100,
			n: 100,
			k: 1,
		},
		{
			name: "10 floors 2 eggs",
			expected: 4,
			n: 10,
			k: 2,
		},
		{
			name: "100 floors 2 eggs",
			expected: 14,
			n: 100,
			k: 2,
		},
		{
			name: "10000 floors 100 eggs",
			expected: 14,
			n: 100,
			k: 2,
		},
	}
	for _,v :=range test{
		t.Run(v.name, func(t *testing.T) {
			actual := solution(v.n,v.k)
			assert.Equal(t,actual,v.expected)
		})
	}
}
