package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var tests = []struct {
		name   string
		a      int
		b      int
		result int
	}{
		{"1 plus 1", 1, 1, 2},
		{"2 plus 2", 2, 2, 5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if result := Add(test.a, test.b); result != test.result {
				t.Errorf("expected %d + %d to be %d got %d", test.a, test.b, test.result, result)
			}
		})
	}

}
