package main

import (
	"testing"
)

func TestMyFunctin(t *testing.T) {
	if test(1, 1) != 2 {
		t.Error("not euqals 2")
	} 
}

func Test(t *testing.T) {
	type a struct {
		desc     string
		value    int
		expected int
	}
	testCases := []a{
		{
			desc:     "value must be one",
			value:    test(1, 1),
			expected: 2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.value != tC.expected {
				t.Error(tC.desc)
			}
		})
	}
}
