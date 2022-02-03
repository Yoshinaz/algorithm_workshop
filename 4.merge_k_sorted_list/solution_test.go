package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

type testCase struct{
	name string
	expected []int
	a [][]int
}
func TestSolution(t *testing.T){
	test := []testCase{
		{
			name: "0 array",
			expected: []int{},
			a:[][]int{

			},
		},
		{
			name: "1 array",
			expected: []int{1, 4, 6, 8},
			a:[][]int{
				{
					1, 4, 6, 8,
				},
			},
		},
		{
			name: "2 array",
			expected: []int{1,2,4,5,6,8,100,101},
			a:[][]int{
				{
					1, 4, 6, 8,
				},
				{
					2, 5, 100, 101,
				},
			},
		},
		{
			name: "3 array",
			expected: []int{1,2,2,4,5,6,6,8,99,100,101,102},
			a:[][]int{
				{
					1, 4, 6, 8,
				},
				{
					2, 5, 100, 101,
				},
				{
					2, 6, 99, 102,
				},
			},
		},
	}
	for _,v :=range test{
		t.Run(v.name, func(t *testing.T) {
			actual := solution(v.a)
			assert.Equal(t,actual,v.expected)
		})
	}
}

