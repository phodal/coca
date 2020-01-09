package concept

import (
	languages2 "github.com/phodal/coca/pkg/application/call/stop_words/languages"
	"github.com/phodal/coca/pkg/domain"
	"github.com/phodal/coca/pkg/infrastructure/constants"
	"github.com/phodal/coca/pkg/infrastructure/string_helper"
)

type ConceptAnalyser struct {
}

func NewConceptAnalyser() ConceptAnalyser {
	return *&ConceptAnalyser{}
}

func (c ConceptAnalyser) Analysis(clzs *[]domain.JClassNode) string_helper.PairList {
	return buildMethodsFromDeps(*clzs)
}

func buildMethodsFromDeps(clzs []domain.JClassNode) string_helper.PairList {
	var methodsName []string
	var methodStr string
	for _, clz := range clzs {
		for _, method := range clz.Methods {
			methodName := method.Name
			methodsName = append(methodsName, methodName)
			methodStr = methodStr + " " + methodName
		}
	}

	words := SegmentConceptCamelcase(methodsName)

	words = removeNormalWords(words)

	wordCounts := string_helper.RankByWordCount(words)
	return wordCounts
}

func removeNormalWords(words map[string]int) map[string]int {
	var newWords = words
	var stopwords = languages2.ENGLISH_STOP_WORDS
	stopwords = append(stopwords, constants.TechStopWords...)
	for _, normalWord := range stopwords {
		if newWords[normalWord] > 0 {
			delete(newWords, normalWord)
		}
	}

	return newWords
}
