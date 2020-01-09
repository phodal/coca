// https://github.com/eMAGTechLabs/go-apriori

/*
MIT License

Copyright (c) 2019 eMAG IT Research

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package apriori

import (
	"errors"
	"sort"
)

const combinationStringChannelLastElement = "STOP"
const combinationIntChannelLastElement = -1
const minLengthNeededForNextCandidates = 3

type SupportRecord struct {
	items   []string
	support float64
}

func (sr SupportRecord) GetItems() []string {
	return sr.items
}

func (sr SupportRecord) GetSupport() float64 {
	return sr.support
}

type OrderedStatistic struct {
	base       []string
	add        []string
	confidence float64
	lift       float64
}

func (os OrderedStatistic) GetBase() []string {
	return os.base
}

func (os OrderedStatistic) GetAdd() []string {
	return os.add
}

func (os OrderedStatistic) GetConfidence() float64 {
	return os.confidence
}

func (os OrderedStatistic) GetLift() float64 {
	return os.lift
}

type RelationRecord struct {
	supportRecord    SupportRecord
	orderedStatistic []OrderedStatistic
}

func (r RelationRecord) GetSupportRecord() SupportRecord {
	return r.supportRecord
}

func (r RelationRecord) GetOrderedStatistic() []OrderedStatistic {
	return r.orderedStatistic
}

type Options struct {
	MinSupport    float64 `json:"minSupport"`    // The minimum support of relations (float).
	MinConfidence float64 `json:"minConfidence"` // The minimum confidence of relations (float).
	MinLift       float64 `json:"minLift"`       // The minimum lift of relations (float).
	MaxLength     int     `json:"maxLength"`     // The maximum length of the relation (integer).
}

func (options Options) check() error {
	// Check Options
	if options.MinSupport <= 0 {
		return errors.New("minimum support must be > 0")
	}

	return nil
}

type Apriori struct {
	transactionNo       int64
	items               []string
	transactionIndexMap map[interface{}][]int64
}

func NewOptions(minSupport float64, minConfidence float64, minLift float64, maxLength int) Options {
	return Options{MinSupport: minSupport, MinConfidence: minConfidence, MinLift: minLift, MaxLength: maxLength}
}

func NewApriori(transactions [][]string) *Apriori {
	var a Apriori
	a.transactionIndexMap = make(map[interface{}][]int64)
	for _, transaction := range transactions {
		a.addTransaction(transaction)
	}

	return &a
}

func (a *Apriori) Calculate(options Options) []RelationRecord {
	if err := options.check(); err != nil {
		panic(err)
	}

	// Calculate supports
	supportRecords := make(chan SupportRecord)
	go a.generateSupportRecords(supportRecords, options.MinSupport, options.MaxLength)

	var relationRecords []RelationRecord
	// Calculate ordered stats
	for {
		supportRecord := <-supportRecords
		if supportRecord.support == -1 {
			break
		}

		filteredOrderedStatistics := a.filterOrderedStatistics(
			a.generateOrderedStatistics(supportRecord),
			options.MinConfidence,
			options.MinLift)

		if len(filteredOrderedStatistics) == 0 {
			continue
		}

		relationRecords = append(relationRecords, RelationRecord{supportRecord, filteredOrderedStatistics})
	}

	return relationRecords
}

func (a *Apriori) addTransaction(transaction []string) {
	for _, item := range transaction {
		if _, ok := a.transactionIndexMap[item]; !ok {
			a.items = append(a.items, item)
			a.transactionIndexMap[item] = []int64{}
		}
		a.transactionIndexMap[item] = append(a.transactionIndexMap[item], a.transactionNo)
	}
	a.transactionNo++
}

// Returns a support for items.
func (a *Apriori) calculateSupport(items []string) float64 {
	// Empty items are supported by all transactions.
	if len(items) == 0 {
		return 1.0
	}

	// Empty transactions supports no items.
	if a.transactionNo == 0 {
		return 0.0
	}

	// Create the transaction index intersection.
	var sumIndexes []int64
	for _, item := range items {
		indexes := a.transactionIndexMap[item]
		// No support for any set that contains a not existing item.
		if len(indexes) == 0 {
			return 0.0
		}
		if len(sumIndexes) == 0 {
			// Assign the indexes on the first time.
			sumIndexes = indexes
		} else {
			// Calculate the intersection on not the first time.
			sumIndexes = a.transactionIntersection(sumIndexes, indexes)
		}
	}

	// Calculate and return the support.
	return float64(len(sumIndexes)) / float64(a.transactionNo)
}

// Returns the initial candidates.
func (a *Apriori) initialCandidates() [][]string {
	var initialCandidates [][]string
	for _, item := range a.getItems() {
		initialCandidates = append(initialCandidates, []string{item})
	}

	return initialCandidates
}

// Returns the item list that the transaction is consisted of.
func (a *Apriori) getItems() []string {
	sort.Strings(a.items)

	return a.items
}

// Returns a generator of ordered statistics as OrderedStatistic instances.
func (a *Apriori) generateOrderedStatistics(record SupportRecord) []OrderedStatistic {
	items := record.items
	sort.Strings(items)

	var ch = make(chan []string)
	defer close(ch)
	go combinations(ch, items, len(items)-1)

	var orderedStatistics []OrderedStatistic
	for combination := range ch {
		if checkIfLastInStringChan(combination) {
			break
		}
		orderedStatistics = append(orderedStatistics, a.generateOrderedStatistic(combination, items, record.support))
	}

	return orderedStatistics
}

func (a *Apriori) generateOrderedStatistic(base []string, items []string, recordSupport float64) OrderedStatistic {
	add := a.itemDifference(items, base)
	supportForBase := a.calculateSupport(base)
	confidence := recordSupport / supportForBase
	supportForAdd := a.calculateSupport(add)
	lift := confidence / supportForAdd

	return OrderedStatistic{base, add, confidence, lift}
}

// Filter OrderedStatistic objects
func (a *Apriori) filterOrderedStatistics(orderedStatistics []OrderedStatistic, minConfidence float64, minLift float64) []OrderedStatistic {
	var filteredOrderedStatistic []OrderedStatistic
	for _, orderedStatistic := range orderedStatistics {
		if orderedStatistic.confidence < minConfidence || orderedStatistic.lift < minLift {
			continue
		}
		filteredOrderedStatistic = append(filteredOrderedStatistic, orderedStatistic)
	}

	return filteredOrderedStatistic
}

// Returns a generator of support records with given transactions.
func (a *Apriori) generateSupportRecords(supportRecordChan chan SupportRecord, minSupport float64, maxLength int) {
	// Process
	candidates := a.initialCandidates()
	var length = 1
	for len(candidates) > 0 {
		var relations [][]string
		for _, relationCandidate := range candidates {
			support := a.calculateSupport(relationCandidate)
			if support < minSupport {
				continue
			}
			relations = append(relations, relationCandidate)
			supportRecordChan <- SupportRecord{relationCandidate, support}
		}
		length++
		if maxLength != 0 && length > maxLength {
			break
		}
		candidates = a.createNextCandidates(relations, length)
	}
	supportRecordChan <- SupportRecord{[]string{}, -1}
}

//func (a *Apriori) generateRelationRecords(relationRecords chan RelationRecord, supportRecord SupportRecord, minConfidence float64, minLift float64) {
//	// Calculate ordered stats
//	filteredOrderedStatistics := a.filterOrderedStatistics(
//		a.generateOrderedStatistics(supportRecord),
//		minConfidence,
//		minLift)
//
//	if len(filteredOrderedStatistics) != 0 {
//		relationRecords <- RelationRecord{supportRecord, filteredOrderedStatistics}
//	}
//}

// Returns the Apriori candidates as a list.
func (a *Apriori) createNextCandidates(prevCandidates [][]string, length int) [][]string {
	var items []string
	for _, candidate := range prevCandidates {
		items = append(items, candidate...)
	}
	sort.Strings(items)
	items = a.uniqueItems(items)

	// Create the temporary candidates. These will be filtered below.
	tmpNextCandidates := a.generateCandidateCombinations(items, length)

	// Return all the candidates if the length of the next candidates is 2
	// because their subsets are the same as items.
	if length < minLengthNeededForNextCandidates {
		return tmpNextCandidates
	}

	// Filter candidates that all of their subsets are
	// in the previous candidates.
	var nextCandidates [][]string
	for _, candidate := range tmpNextCandidates {
		candidateCombinations := a.generateCandidateCombinations(candidate, length-1)

		allAreInPrev := 0
		for _, candidates := range candidateCombinations {
			if a.isSubset(candidates, prevCandidates) {
				allAreInPrev++
			}
		}
		if allAreInPrev == len(candidateCombinations) {
			nextCandidates = append(nextCandidates, candidate)
		}
	}

	return nextCandidates
}

func (a *Apriori) generateCandidateCombinations(items []string, length int) [][]string {
	var tmpNextCandidates [][]string
	if len(items) >= length {
		var ch = make(chan []string)
		defer close(ch)
		go combinations(ch, items, length)

		for candidate := range ch {
			if checkIfLastInStringChan(candidate) {
				break
			}
			tmpNextCandidates = append(tmpNextCandidates, candidate)
		}
	}

	return tmpNextCandidates
}

func (a *Apriori) isSubset(needle []string, haystack [][]string) bool {
	needleLen := len(needle)
	for _, value := range haystack {
		found := 0
		for _, i := range needle {
			if a.inSlice(i, value) {
				found++
			}
		}

		if needleLen > found {
			continue
		}
		return true
	}

	return false
}

func (a *Apriori) inSlice(needle string, haystack []string) bool {
	for _, str := range haystack {
		if str == needle {
			return true
		}
	}

	return false
}

func (a *Apriori) uniqueItems(items []string) []string {
	keys := make(map[string]bool)
	var uniqueItems []string
	for _, entry := range items {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			uniqueItems = append(uniqueItems, entry)
		}
	}

	return uniqueItems
}

func (a *Apriori) transactionIntersection(first, second []int64) []int64 {
	m := make(map[int64]bool)
	var intersection []int64

	for _, item := range first {
		m[item] = true
	}

	for _, item := range second {
		if _, ok := m[item]; ok {
			intersection = append(intersection, item)
		}
	}

	return intersection
}

func (a *Apriori) itemDifference(first []string, second []string) []string {
	var diff []string
	// Loop two times, first to find first strings not in second,
	// second loop to find second strings not in first
	for i := 0; i < 2; i++ {
		for _, firstString := range first {
			found := false
			for _, secondString := range second {
				if firstString == secondString {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, firstString)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			first, second = second, first
		}
	}

	return diff
}

func combinations(ch chan []string, iterable []string, r int) {
	if r != 0 {
		length := len(iterable)

		if r > length {
			panic("Invalid arguments")
		}

		intCh := make(chan []int)
		defer close(intCh)
		go genCombinations(intCh, length, r)

		for comb := range intCh {
			if checkIfLastInIntChan(comb) {
				break
			}
			result := make([]string, r)
			for i, val := range comb {
				result[i] = iterable[val]
			}
			ch <- result
		}
	} else {
		result := make([]string, r)
		ch <- result
	}
	ch <- []string{combinationStringChannelLastElement}
}

func genCombinations(ch chan []int, n, r int) {
	result := make([]int, r)
	for i := range result {
		result[i] = i
	}

	temp := make([]int, r)
	copy(temp, result) // avoid overwriting of result
	ch <- temp

	for {
		for i := r - 1; i >= 0; i-- {
			if result[i] < i+n-r {
				result[i]++
				for j := 1; j < r-i; j++ {
					result[i+j] = result[i] + j
				}
				temp := make([]int, r)
				copy(temp, result) // avoid overwriting of result
				ch <- temp
				break
			}
		}
		if result[0] >= n-r {
			break
		}
	}
	ch <- []int{combinationIntChannelLastElement}
}

func checkIfLastInStringChan(strings []string) bool {
	return len(strings) > 0 && strings[0] == combinationStringChannelLastElement
}

func checkIfLastInIntChan(ints []int) bool {
	return len(ints) > 0 && ints[0] == combinationIntChannelLastElement
}
