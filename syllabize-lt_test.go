package main

import (
	"reflect"
	"testing"
)

type StrToStrListTest struct {
	word string
	exp  []string
}

func LoopTests(t *testing.T, tests []StrToStrListTest,
	testedFunc func(w string) []string, name string) {
	for _, test := range tests {
		actual := testedFunc(test.word)
		if !reflect.DeepEqual(actual, test.exp) {
			t.Errorf("%s(%q) = %v, but expected %v\n", name, test.word, actual,
				test.exp)
		}
	}
}

func TestSplitSounds(t *testing.T) {
	var tests = []StrToStrListTest{
		{"labas", []string{"l", "a", "b", "a", "s"}},
		{"rytas", []string{"r", "y", "t", "a", "s"}},
		{"malonu", []string{"m", "a", "l", "o", "n", "u"}},
		{"jus", []string{"j", "u", "s"}},
		{"matyti", []string{"m", "a", "t", "y", "t", "i"}},
		{"džipas", []string{"dž", "i", "p", "a", "s"}},
		{"dzūkas", []string{"dz", "ū", "k", "a", "s"}},
	}
	LoopTests(t, tests, splitSounds, "splitSounds")
}

func TestSyllabificate(t *testing.T) {
	var tests = []StrToStrListTest{
		{"labas", []string{"la", "bas"}},
		{"rytas", []string{"ry", "tas"}},
		{"malonu", []string{"ma", "lo", "nu"}},
		{"jus", []string{"jus"}},
		{"matyti", []string{"ma", "ty", "ti"}},
		{"džipas", []string{"dži", "pas"}},
		{"dzūkas", []string{"dzū", "kas"}},
	}
	LoopTests(t, tests, syllabificate, "syllabificate")
}
