package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

type testCase struct{
	name string
	expected int
	a []int
}
func TestSolutionFindLength(t *testing.T){
	test := []testCase{
		{
			name: "empty",
			expected: 0,
			a:[]int{},
		},
		{
			name: "1",
			expected: 1,
			a:[]int{1},
		},
		{
			name: "2",
			expected: 1,
			a:[]int{2,1},
		},
		{
			name: "exist",
			expected: 6,
			a:[]int{10, 22, 9, 33, 21, 50, 41, 60, 80},
		},
		{
			name: "minus",
			expected: 7,
			a:[]int{10, -22, -9, 33, -1,21, 50, 41, 60, 80},
		},
	}
	for _,v :=range test{
		t.Run(v.name, func(t *testing.T) {
			actual := solutionFindLength(v.a)
			assert.Equal(t,actual,v.expected)
		})
	}
}
