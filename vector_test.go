package solarsim

import (
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		a, b, expected Vector
	}{
		{Vector{0, 0}, Vector{1, 0}, Vector{1, 0}},
	}
	for _, c := range cases {
		c.a.add(c.b)
		if c.a != c.expected {
			t.Errorf("")
		}
	}
}