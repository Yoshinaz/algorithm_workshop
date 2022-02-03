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
			name: "not found",
			expected: false,
			a:[]int{1,2,3,4,5},
			k: 14,
		},
		{
			name: "found",
			expected: true,
			a:[]int{1,2,3,4,5},
			k: 10,
		},
	}
	for _,v :=range test{
		t.Run(v.name, func(t *testing.T) {
			actual := solution(v.a,v.k)
			assert.Equal(t,actual,v.expected)
		})
	}
}
