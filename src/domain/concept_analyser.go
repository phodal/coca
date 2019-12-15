package domain

import (
	"coca/src/adapter/models"
	"coca/src/domain/stop_words/languages"
	"coca/src/domain/support"
	"fmt"
)

type ConceptAnalyser struct {
}

func NewConceptAnalyser() ConceptAnalyser {
	return *&ConceptAnalyser{}
}

func (c ConceptAnalyser) run() {

}

func (c ConceptAnalyser) Analysis(clzs *[]models.JClassNode) {
	buildMethodsFromDeps(*clzs)
}

func buildMethodsFromDeps(clzs []models.JClassNode) {
	var methodsName []string
	var methodStr string
	for _, clz := range clzs {
		for _, method := range clz.Methods {
			methodName := method.Name
			methodsName = append(methodsName, methodName)
			methodStr = methodStr + " " + methodName
		}
	}

	words := support.SegmentConceptCamelcase(methodsName)

	words = removeNormalWords(words)

	wordCounts := support.RankByWordCount(words)
	for _, word := range wordCounts {
		if word.Value > 0 {
			fmt.Println(word.Key, word.Value)
		}
	}
}

var itStopWords = []string{
	"get",
	"create",
	"update",
	"delete",
	"save",

	"add",
	"insert",
	"select",

	"exist",
	"find",
	"new",
	"parse",

	"set",
	"get",
	"first",
	"last",

	"type",
	"key",
	"value",
	"equal",
	"greater",
	"greater",

	"all",
	"by",
	"id",
	"is",
	"of",
	"not",
	"with",
	"main",

	"status",
	"count",
	"equals",
	"start",
	"config",
	"sort",
}

func removeNormalWords(words map[string]int) map[string]int {
	var newWords = words
	var stopwords = languages.ENGLISH_STOP_WORDS
	stopwords = append(stopwords, itStopWords...)
	for _, normalWord := range stopwords {
		if newWords[normalWord] > 0 {
			delete(newWords, normalWord)
		}
	}

	fmt.Println(newWords)
	return newWords
}
