package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

type testCase struct{
	name string
	expected bool
	a []int
	k int
}
func TestSolution(t *testing.T){
	test := []testCase{
		{
			name: "empty",
			expected: false,
			a:[]int{},
			k: 100,
		},
		{
			name: "1",
			expected: false,
			a:[]int{1},
			k: 100,
		},
		{
			name: "found 1",
			expected: true,
			a:[]int{1,2,3,4,5},
			k: 14,
		},
		{
			name: "found",
			expected: true,
			a:[]int{1,2,3,4,5},
			k: 10,
		},
		{
			name: "not found",
			expected: false,
			a:[]int{28,32,87,65,98},
			k: 66,
		},
		{
			name: "found 28",
			expected: true,
			a:[]int{28,32,87,65,98},
			k: 28,
		},
		{
			name: "found 93",
			expected: true,
			a:[]int{28,32,87,65,98},
			k: 93,
		},
		{
			name: "found 180",
			expected: true,
			a:[]int{28,32,87,65,98},
			k: 180,
		},
		{
			name: "not found 208",
			expected: false,
			a:[]int{28,32,87,65,98},
			k: 208,
		},
		{
			name: "found 98",
			expected: true,
			a:[]int{28,32,87,65,98},
			k: 98,
		},
	}
	for _,v :=range test{
		t.Run(v.name, func(t *testing.T) {
			actual := solution(v.a,v.k)
			assert.Equal(t,actual,v.expected)
		})
	}
}
