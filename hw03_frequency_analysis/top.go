package hw03frequencyanalysis

import (
	"regexp"
	"sort"
)

var splitRegex = regexp.MustCompile(`\s+`) // regexp for split source string

func Top10(source string) []string {
	if len(source) == 0 {
		return nil
	}
	splitted := splitRegex.Split(source, -1) // split source
	countDict := map[string]int{}            // map for count of words
	var max = 0                              // variable for get max count
	for _, s := range splitted {
		count := countDict[s] + 1
		countDict[s] = count // counting this word
		if count > max {
			max = count
		}
	}
	wordsGroupedByCount := map[int][]string{} // map for grouped words by count and sorting them after
	for word, counter := range countDict {
		value, ok := wordsGroupedByCount[counter]
		if !ok { // if map dos't contain this words counter
			wordsGroupedByCount[counter] = append([]string{}, word)
		} else { // else append word to slice
			wordsGroupedByCount[counter] = append(value, word)
		}
	}
	res := make([]string, 0)   // result slice
	for i := max; i > 0; i-- { // iterate from max to zero count
		if wordsTheSameCount, ok := wordsGroupedByCount[i]; ok {
			sort.Strings(wordsTheSameCount)
			res = append(res, wordsTheSameCount...)
		}
	}
	return res[:10]
}
