package main

import (
	"testing"
)

func TestCorrectMostElf(t *testing.T) {
	tests := []struct {
		input         string
		expected, cal int
	}{
		{
			`1

2

1
1
1`, 3, 3,
		}, {
			`1

2
1
4
5

4

4
2
1

3

`, 2, 12,
		}, {
			`4000
30000
2

2
3
4
5
6

2

123
`, 1, 34002,
		},
	}

	for _, test := range tests {
		x, c := findMostElf(test.input)
		if x != test.expected && c != test.cal {
			t.Logf("Expected %d, %d got %d, %d", test.expected, test.cal, x, c)
			t.Fail()
		}
	}
}
