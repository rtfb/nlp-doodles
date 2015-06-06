package main

import (
	"fmt"
	"strings"

	"gopkg.in/fatih/set.v0"
)

const (
	vowels = "aąeęėiįyouųū"
)

var (
	S       = set.New("s", "z", "š", "ž")
	T       = set.New("p", "b", "t", "d", "k", "g", "c", "č", "dz", "dž")
	R       = set.New("l", "m", "n", "r", "v", "j")
	STRULES = set.New("STR", "ST", "SR", "TR")
)

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

func SoundToSTR(sound string) string {
	if S.Has(sound) {
		return "S"
	}
	if T.Has(sound) {
		return "T"
	}
	if R.Has(sound) {
		return "R"
	}
	return "?"
}

func isVowel(sound string) bool {
	return strings.Contains(vowels, sound)
}

func filter(s []string, fn func(string) bool) []string {
	var p []string // == nil
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}

func syllabificate(word string) []string {
	var syllables []string
	var consonants []string
	syllable := ""
	STR := ""
	wasVowel := false
	for _, sound := range splitSounds(word) {
		if isVowel(sound) && wasVowel {
			syllable += sound
		} else if isVowel(sound) {
			carryConsonants := (len(consonants) > 0 && len(syllables) == 0) ||
				(len(consonants) == 1) ||
				STRULES.Has(STR)
			if carryConsonants {
				syllables = append(syllables, syllable)
				syllable = strings.Join(consonants, "")
			} else if len(STR) > 2 && STRULES.Has(STR[len(STR)-2:]) {
				syllable += strings.Join(consonants[:len(consonants)-2], "")
				syllables = append(syllables, syllable)
				syllable = strings.Join(consonants[len(consonants)-2:], "")
			} else if len(STR) > 1 {
				syllable += strings.Join(consonants[:len(consonants)-1], "")
				syllables = append(syllables, syllable)
				syllable = strings.Join(consonants[len(consonants)-1:], "")
			} else {
				syllable = strings.Join(consonants, "")
				syllables = append(syllables, syllable)
				syllable = ""
			}
			STR = ""
			consonants = nil
			syllable += sound
		} else {
			STR += SoundToSTR(sound)
			consonants = append(consonants, sound)
			if len(STR) > 3 {
				syllable += consonants[0]
				consonants = consonants[1:]
				STR = STR[1:]
			}
		}
		wasVowel = isVowel(sound)
	}
	if syllable != "" {
		syllables = append(syllables, syllable+strings.Join(consonants, ""))
	} else if len(consonants) > 0 && len(syllables) > 0 {
		syllables[len(syllables)-1] += strings.Join(consonants, "")
	}
	// TODO: try to avoid this filtering, make it return syllables
	return filter(syllables, func(s string) bool {
		if s == "" {
			return false
		}
		return true
	})
}

func main() {
	s := "labas rytas malonu jus matyti"
	l := strings.Split(s, " ")
	for _, w := range l {
		fmt.Printf("%s: %#v\n", w, syllabificate(w))
	}
}
