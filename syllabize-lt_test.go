package main

import (
	"reflect"
	"strings"
	"testing"
)

type StrToStrListTest struct {
	word string
	exp  string
}

func LoopTests(t *testing.T, tests []StrToStrListTest,
	testedFunc func(w string) []string, name string) {
	for _, test := range tests {
		actual := strings.Join(testedFunc(test.word), "-")
		if !reflect.DeepEqual(actual, test.exp) {
			t.Errorf("%s(%q) = %v, but expected %v\n", name, test.word, actual,
				test.exp)
		}
	}
}

func TestSplitSounds(t *testing.T) {
	var tests = []StrToStrListTest{
		{"labas", "l-a-b-a-s"},
		{"rytas", "r-y-t-a-s"},
		{"malonu", "m-a-l-o-n-u"},
		{"jus", "j-u-s"},
		{"matyti", "m-a-t-y-t-i"},
		{"džipas", "dž-i-p-a-s"},
		{"dzūkas", "dz-ū-k-a-s"},
	}
	LoopTests(t, tests, splitSounds, "splitSounds")
}

func TestSyllabificate(t *testing.T) {
	var tests = []StrToStrListTest{
		{"labas", "la-bas"},
		{"rytas", "ry-tas"},
		{"malonu", "ma-lo-nu"},
		{"jus", "jus"},
		{"matyti", "ma-ty-ti"},
		{"džipas", "dži-pas"},
		{"dzūkas", "dzū-kas"},
	}
	LoopTests(t, tests, syllabificate, "syllabificate")
}
