package utils

import (
	"strings"

	snowballhu "github.com/kljensen/snowball/hungarian"
)

func lowercaseFilter(tokens []string) []string {
	r := make([]string, len(tokens))

	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}

	return r
}

func stopWordFilter(tokens []string) []string {
	var stopWords = map[string]struct{}{
		"a": {}, "ahogy": {}, "ahol": {}, "aki": {}, "akik": {}, "akkor": {}, "alatt": {}, "által": {}, "általában": {}, "amely": {}, "amelyek": {}, "amelyekben": {}, "amelyeket": {}, "amelyet": {}, "amelynek": {}, "ami": {}, "amit": {}, "amolyan": {}, "amíg": {}, "amikor": {}, "át": {}, "abban": {}, "ahhoz": {}, "annak": {}, "arra": {}, "arról": {}, "az": {}, "azok": {}, "azon": {}, "azt": {}, "azzal": {}, "azért": {}, "aztán": {}, "azután": {}, "azonban": {}, "bár": {}, "be": {}, "belül": {}, "benne": {}, "cikk": {}, "cikkek": {}, "cikkeket": {}, "csak": {}, "de": {}, "e": {}, "eddig": {}, "egész": {}, "egy": {}, "egyes": {}, "egyetlen": {}, "egyéb": {}, "egyik": {}, "egyre": {}, "ekkor": {}, "el": {}, "elég": {}, "ellen": {}, "elő": {}, "először": {}, "előtt": {}, "első": {}, "én": {}, "éppen": {}, "ebben": {}, "ehhez": {}, "emilyen": {}, "ennek": {}, "erre": {}, "ez": {}, "ezt": {},
	}

	r := make([]string, 0, len(tokens))

	for _, token := range tokens {
		if _, ok := stopWords[token]; !ok {
			r = append(r, token)
		}
	}

	return r
}

func stemmerFilter(tokens []string) []string {
	r := make([]string, len(tokens))

	for i, token := range tokens {
		r[i] = snowballhu.Stem(token, false)
	}

	return r
}
