package main

import (
	"reflect"
	"testing"
)

func TestSplitSounds(t *testing.T) {
	var tests = []struct {
		word string
		exp  []string
	}{
		{"labas", []string{"l", "a", "b", "a", "s"}},
		{"rytas", []string{"r", "y", "t", "a", "s"}},
		{"malonu", []string{"m", "a", "l", "o", "n", "u"}},
		{"jus", []string{"j", "u", "s"}},
		{"matyti", []string{"m", "a", "t", "y", "t", "i"}},
		{"džipas", []string{"dž", "i", "p", "a", "s"}},
		{"dzūkas", []string{"dz", "ū", "k", "a", "s"}},
	}
	for _, test := range tests {
		actual := splitSounds(test.word)
		if !reflect.DeepEqual(actual, test.exp) {
			t.Errorf("splitSounds(%q) = %v, but expected %v\n", test.word,
				actual, test.exp)
		}
	}
}
