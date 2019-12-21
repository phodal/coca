package concept

import (
	languages2 "coca/core/domain/call_graph/stop_words/languages"
	"coca/core/models"
	"coca/core/support"
)

type ConceptAnalyser struct {
}

func NewConceptAnalyser() ConceptAnalyser {
	return *&ConceptAnalyser{}
}

func (c ConceptAnalyser) run() {

}

func (c ConceptAnalyser) Analysis(clzs *[]models.JClassNode) support.PairList {
	return buildMethodsFromDeps(*clzs)
}

func buildMethodsFromDeps(clzs []models.JClassNode) support.PairList {
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

	wordCounts := support.RankByWordCount(words)
	return wordCounts
}

var itStopWords = []string{
	"get",
	"create",
	"update",
	"delete",
	"save",

	"add",
	"remove",
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
	"handle",
	"handler",
	"internal",
	"cache",
	"request",
	"process",

	"parameter",
	"method",
	"class",
	"default",
	"object",
	"annotation",

	"read",
	"write",

	"bean",
	"message",
	"factory",
	"error",
	"error",
	"exception",
}

func removeNormalWords(words map[string]int) map[string]int {
	var newWords = words
	var stopwords = languages2.ENGLISH_STOP_WORDS
	stopwords = append(stopwords, itStopWords...)
	for _, normalWord := range stopwords {
		if newWords[normalWord] > 0 {
			delete(newWords, normalWord)
		}
	}

	return newWords
}
