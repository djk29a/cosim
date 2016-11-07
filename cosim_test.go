package cosim

import (
	"testing"
)

func range_test(t *testing.T, v float64) {
	if v < 0.0 {
		t.Error("Cosine similarity cannot be negative: ", v)
	}
	if v > 1.0 {
		t.Error("Cosine similarity cannot be greater than 1.0:", v)
	}
}

func TestCos(t *testing.T) {

	a, _ := Cosine("a", "a", false)
	range_test(t, a)
	if a != 1.0 {
		t.Error("cosine of string to itself is always exactly 1.0")
	}
	b, _ := Cosine("a", "b", false)
	range_test(t, b)

	_, err := Cosine("", "", true)
	if err == nil {
		t.Error("Empty string cosine should result in error")
	}
	c, _ := Cosine("aaaa", "AAAA", true)
	range_test(t, c)
	if c != 1.0 {
		t.Error("cosine of string to itself (case insensitive) is exactly 1.0", c)
	}
}
