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
			t.Errorf("%s(%q) = %q, but expected %q\n", name, test.word, actual,
				test.exp)
		}
	}
}

func ToSlice(c <-chan string) []string {
	s := make([]string, 0)
	for i := range c {
		s = append(s, i)
	}
	return s
}

func LoopChanTests(t *testing.T, tests []StrToStrListTest,
	testedFunc func(w string) <-chan string, name string) {
	for _, test := range tests {
		actual := strings.Join(ToSlice(testedFunc(test.word)), "-")
		if !reflect.DeepEqual(actual, test.exp) {
			t.Errorf("%s(%q) = %q, but expected %q\n", name, test.word, actual,
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
	LoopChanTests(t, tests, splitSounds, "splitSounds")
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

		// ST
		// ?? {"stiklas", "stik-las"},
		// ?? {"spausti", "spaus-ti"},

		// SR
		// ?? {"slinkti", "slink-ti"},
		// ?? {"žvilgsnis", "žvilgs-nis"},

		// TR
		{"kratyti", "kra-ty-ti"},
		{"protas", "pro-tas"},

		{"aistra", "ai-stra"},
		{"aštrus", "a-štrus"},
		{"sėkla", "sė-kla"},
		{"mįslė", "mį-slė"},
		// ?? {"akti", "ak-ti"},

		// ?? {"išvaizda", "iš-vai-zda"},
		{"liūdnas", "liū-dnas"},

		{"medus", "me-dus"},
		{"siena", "sie-na"},
		// ?? {"kalnai", "kal-nai"},
		// ?? {"kalnas", "kal-nas"},
		{"ėmė", "ė-mė"},
		{"uodas", "uo-das"},
		// ?? {"arti", "ar-ti"},
		{"gėrio", "gė-rio"},
		{"keliu", "ke-liu"},
		{"važiavo", "va-žia-vo"},

		{"neša", "ne-ša"},
		{"ratas", "ra-tas"},
		{"stalas", "sta-las"},
		{"skylė", "sky-lė"},
		{"štai", "štai"},
		{"slogus", "slo-gus"},
		// ?? {"švelnus", "švel-nus"},
		{"žvėris", "žvė-ris"},
		{"skraidyti", "skrai-dy-ti"},
		// ?? {"sprogti", "sprog-ti"},
		{"strėlė", "strė-lė"},

		// ?? {"apsiašaroti", "ap-si-a-ša-ro-ti"},
		// ?? {"įsiamžinti", "į-si-am-žin-ti"},
		// ?? {"neilgas", "ne-il-gas"},
		// ?? {"paimti", "pa-im-ti"},
		// ?? {"paupys", "pa-u-pys"},
		// ?? {"suiro", "su-i-ro"},
		// ?? {"suošė", "su-o-šė"},
		// ?? {"neinicijuoti", "ne-i-ni-ci-juo-ti"},
		// ?? {"pailiustruoti", "pa-i-liu-struo-ti"},
		// ?? {"persiorientuoti", "per-si-o-rien-tuo-ti"},
		// ?? {"suintriguoti", "su-in-tri-guo-ti"},
		{"audra", "au-dra"},
		{"ąžuolas", "ą-žuo-las"},
		{"ėjo", "ė-jo"},
		{"ola", "o-la"},
		{"uosis", "uo-sis"},

		// ?? {"akvariumas", "ak-va-ri-u-mas"},
		// ?? {"biologija", "bi-o-lo-gi-ja"},
		// ?? {"cianas", "ci-a-nas"},
		// ?? {"čempionas", "čem-pi-o-nas"},
		// ?? {"abiturientas", "a-bi-tu-ri-en-tas"},
		// ?? {"audiencija", "au-di-en-ci-ja"},
		// ?? {"higiena", "hi-gi-e-na"},
		// ?? {"orientacija", "o-ri-en-ta-ci-ja"},
		// ?? {"induizmas", "in-du-iz-mas"},
		// ?? {"intuicija", "in-tu-i-ci-ja"},
		// ?? {"jėzuitai", "jė-zu-i-tai"},
		{"chemija", "che-mi-ja"},
		{"choras", "cho-ras"},
		{"skriauda", "skriau-da"},
		// ?? {"strigti", "strig-ti"},
		// ?? {"skirti", "skir-ti"},
		// ?? {"spinta", "spin-ta"},
		{"stogas", "sto-gas"},
		{"greitis", "grei-tis"},
		{"priekis", "prie-kis"},
		// ?? {"traukti", "trauk-ti"},

		// ?? {"vaiskrūmis", "vais-krū-mis"},
		// ?? {"atremti", "at-rem-ti"},

		{"laukas", "lau-kas"},
		{"nešiojo", "ne-šio-jo"},

		// Kai šaknis prasideda balsiu, ankstesnės morfemos galinis
		// priebalsis patenka į tolesnį skiemenį:
		// ?? {"antakis", "an-ta-kis"},
		{"pelėda", "pe-lė-da"},

		// afrikatomis nelaikomos tokio tipo samplaikos:
		// ?? {"juodžemis", "juod-že-mis"},
	}
	LoopTests(t, tests, syllabificate, "syllabificate")
}
