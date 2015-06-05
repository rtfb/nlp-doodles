package main

import (
	"fmt"
	"strings"
)

//const (
//	vowels       = "aąeęėiįyouųū"
//	front_vowels = "iįyeęė"
//	back_vowels  = "aąouųū"
//)

// XXX: convert to chans instead of accumulating array?
func splitSounds(word string) []string {
	var result []string
	last := ""
	for _, char := range word {
		if (char == 'z' || char == 'ž') && last == "d" {
			result = append(result, last+string(char))
			last = ""
			continue
		}
		if last != "" {
			result = append(result, last)
		}
		last = string(char)
	}
	if last != "" {
		result = append(result, last)
	}
	return result
}

func syllabificate(word string) []string {
	var syllables []string
	//var consonants []string
	//syllable := ""
	//STR := ""
	//isVowel := false
	//for _, sound := range splitSounds(word) {
	//}
	return syllables
}

func main() {
	s := "labas rytas malonu jus matyti"
	l := strings.Split(s, " ")
	for _, w := range l {
		fmt.Printf("%s: %v\n", w, syllabificate(w))
	}
}
