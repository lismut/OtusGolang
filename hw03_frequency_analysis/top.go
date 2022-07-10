package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

// A data structure to hold a key/value pair.
type Pair struct {
	Key   string
	Value int
}

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9А-Яа-я\-]+`)

func clearString(str string) string {
	return strings.ToLower(nonAlphanumericRegex.ReplaceAllString(str, ""))
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int      { return len(p) }
func (p PairList) Less(i, j int) bool {
	return p[i].Value > p[j].Value || (p[i].Value == p[j].Value && p[i].Key < p[j].Key)
}

// A function to turn a map into a PairList, then sort and return it.
func sortMapByValue(m map[string]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	return p
}

func Top10(input string) []string {
	inputSlice := strings.Fields(input)
	counter := make(map[string]int)
	for _, r := range inputSlice {
		if r != "-" {
			counter[clearString(r)]++
		}
	}
	sortedCounter := sortMapByValue(counter)
	res := []string{}
	for i, value := range sortedCounter {
		res = append(res, value.Key)
		if i == 9 {
			break
		}
	}
	return res
}
