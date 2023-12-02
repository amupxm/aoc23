package d1

import "testing"

func TestD1(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{`t1`, 142},
		{`t2`, 54719},
		{`t3`, 281 + 18},
	}

	for _, test := range tests {
		if got := Solution(test.input); got != test.want {
			t.Errorf("Solution(%q) = %v", test.input, got)
		}
	}
}
