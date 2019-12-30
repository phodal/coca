package concept

import (
	"github.com/phodal/coca/config"
	languages2 "github.com/phodal/coca/core/domain/call_graph/stop_words/languages"
	"github.com/phodal/coca/core/models"
	"github.com/phodal/coca/core/infrastructure"
)

type ConceptAnalyser struct {
}

func NewConceptAnalyser() ConceptAnalyser {
	return *&ConceptAnalyser{}
}

func (c ConceptAnalyser) run() {

}

func (c ConceptAnalyser) Analysis(clzs *[]models.JClassNode) infrastructure.PairList {
	return buildMethodsFromDeps(*clzs)
}

func buildMethodsFromDeps(clzs []models.JClassNode) infrastructure.PairList {
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

	wordCounts := infrastructure.RankByWordCount(words)
	return wordCounts
}

func removeNormalWords(words map[string]int) map[string]int {
	var newWords = words
	var stopwords = languages2.ENGLISH_STOP_WORDS
	stopwords = append(stopwords, config.TechStopWords...)
	for _, normalWord := range stopwords {
		if newWords[normalWord] > 0 {
			delete(newWords, normalWord)
		}
	}

	return newWords
}
