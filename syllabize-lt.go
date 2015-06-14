package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"gopkg.in/fatih/set.v0"
)

type templateEntry struct {
	template string
	tags     []string
}

const (
	vowels = "aąeęėiįyouųū"
)

var (
	S       = set.New("s", "z", "š", "ž")
	T       = set.New("p", "b", "t", "d", "k", "g", "c", "č", "dz", "dž")
	R       = set.New("l", "m", "n", "r", "v", "j")
	STRULES = set.New("STR", "ST", "SR", "TR")
)

func templateToRegexp(tmpl string) *regexp.Regexp {
	reStr := strings.Replace(tmpl, "-", "", -1)
	if reStr[0] == '*' {
		reStr = strings.Replace(reStr, "*", ".*", 1)
	}
	if strings.HasSuffix(reStr, "#") {
		reStr = strings.Replace(reStr, "#", ".*", 1)
	}
	if strings.HasSuffix(reStr, "*") {
		tmp := strings.TrimRight(reStr, "*")
		reStr = tmp + ".+"
	}
	return regexp.MustCompile(reStr)
}

func scanTemplate(text string) templateEntry {
	var result templateEntry
	parts := strings.Split(text, " ")
	if len(parts) < 1 {
		panic("can't parse template: " + text)
	}
	result.template = parts[0]
	for _, t := range parts[1:] {
		result.tags = append(result.tags, strings.Trim(t, "<>"))
	}
	return result
}

func loadTemplates(dict string) []templateEntry {
	file, err := os.Open(dict)
	if err != nil {
		panic("can't load templates dict")
	}
	defer file.Close()
	var result []templateEntry
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanTemplate(scanner.Text()))
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	return result
}

func splitSounds(word string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		last := ""
		for _, char := range word {
			if (char == 'z' || char == 'ž') && last == "d" {
				out <- last + string(char)
				last = ""
				continue
			}
			if last != "" {
				out <- last
			}
			last = string(char)
		}
		if last != "" {
			out <- last
		}
	}()
	return out
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

func syllabificate(word string) []string {
	var syllables []string
	var consonants []string
	syllable := ""
	STR := ""
	wasVowel := false
	for sound := range splitSounds(word) {
		if isVowel(sound) && wasVowel {
			syllable += sound
		} else if isVowel(sound) {
			carryConsonants := (len(consonants) > 0 && len(syllables) == 0) ||
				(len(consonants) == 1) ||
				STRULES.Has(STR)
			if carryConsonants {
				if syllable != "" {
					syllables = append(syllables, syllable)
				}
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
				if syllable != "" {
					syllables = append(syllables, syllable)
				}
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
	return syllables
}

func main() {
	s := "labas rytas malonu jus matyti"
	l := strings.Split(s, " ")
	for _, w := range l {
		fmt.Printf("%s: %#v\n", w, syllabificate(w))
	}
	tmpl := loadTemplates("templates.dict")
	fmt.Printf("%+v\n", tmpl)
}
