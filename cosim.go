// Package cosim calculates cosine similarity of strings, slices of strings, or
// two channels that will read strings from each channel
//
package cosim

import (
	"errors"
	"math"
	"unicode"
)

// "Cosine" similarity of two strings. It's in quotes because
// we don't actually use the cosine function
func Cosine(a, b string, caseSensitive bool) (cosine float64, err error) {
	a_len := len(a)
	b_len := len(a)
	rank := a_len

	if rank < b_len {
		rank = len(b)
	}
	if a_len == 0 || b_len == 0 {
		return 0.0, errors.New("Cosine similarity undefined on empty strings")
	}

	var sum, s0, s1 uint64 = 0, 0, 0

	for i := 0; i < rank; i++ {
		a_i := uint64(rune(a[i]))
		b_i := uint64(rune(b[i]))
		if !caseSensitive {
			a_i = uint64(unicode.ToLower(rune(a[i])))
			b_i = uint64(unicode.ToLower(rune(b[i])))
		}
		if i >= a_len {
			a_i = uint64(0)
		}
		if i >= b_len {
			b_i = uint64(0)
		}

		sum += a_i * b_i
		s0 += a_i * a_i
		s1 += b_i * b_i
	}
	return float64(sum) / (math.Sqrt(float64(s0)) * math.Sqrt(float64(s1))), nil
}

// Cosine similarity of two string arrays, usually obtained after performing
// a split on an original string
func CosineWord(a, b []string, caseSensitive bool) (cosine float64, err error) {
	// construct intersection set of strings
	return 0.0, nil
}

type CosAccum struct {
	count                                   uint64
	a_len, b_len                            uint64
	numerator, denominator_a, denominator_b uint64
	//	cache to hold onto set, trie, huffman tree
}

// Accumulating function variant of Cosine

// Accumulating function variant of CosineWordString
