package d2

import (
	"testing"
)

func TestD1(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{`t1`, 2286},
		{`t2`, 70265},
	}

	for _, test := range tests {
		if got := Solution(test.input); got != test.want {
			t.Errorf("Solution(%q) = %v", test.input, got)
		}
	}
}
