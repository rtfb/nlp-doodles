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

func TestSyllabificate(t *testing.T) {
	var tests = []struct {
		word string
		exp  []string
	}{
		{"labas", []string{"la", "bas"}},
		{"rytas", []string{"ry", "tas"}},
		{"malonu", []string{"ma", "lo", "nu"}},
		{"jus", []string{"jus"}},
		{"matyti", []string{"ma", "ty", "ti"}},
		{"džipas", []string{"dži", "pas"}},
		{"dzūkas", []string{"dzū", "kas"}},
	}
	for _, test := range tests {
		actual := syllabificate(test.word)
		if !reflect.DeepEqual(actual, test.exp) {
			t.Errorf("syllabificate(%q) = %v, but expected %v\n", test.word,
				actual, test.exp)
		}
	}
}
