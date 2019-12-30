package support

import "github.com/phodal/coca/config"

func IsTechStopWords(firstWord string) bool {
	for _, word := range config.TechStopWords {
		if word == firstWord {
			return true;
		}
	}

	return false;
}

