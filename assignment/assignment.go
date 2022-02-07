package assignment

import (
	"math"
	"regexp"
	"sort"
	"strings"
)

func AddUint32(x, y uint32) (uint32, bool) {

	overflowFlag := false
	intResult := int(x) + int(y)
	uintResult := x + y

	// if there is overflow in the arithmetic
	if intResult > math.MaxUint32 {
		overflowFlag = true
	}

	return uintResult, overflowFlag
}

func CeilNumber(f float64) float64 {

	ceilByQuarter := math.Ceil(f*4) / 4
	return ceilByQuarter
}

func AlphabetSoup(s string) string {

	finalString := ""
	charCount := make(map[string]int)

	// collect string counts
	for _, v := range s {
		if charCount[string(v)] == 0 {
			charCount[string(v)] = 1
		} else {
			charCount[string(v)] += 1
		}
	}

	// sorting charCount hashmap by keys
	keys := make([]string, 0, len(charCount))

	for k := range charCount {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	// construct final string to be returned
	for _, k := range keys {
		// for every charCount there is, append it
		for i := 0; i < charCount[k]; i++ {
			finalString += k
		}
	}

	return finalString
}

func StringMask(s string, n uint) string {

	strLen := len(s)

	// for the case where there is an empty string
	// but given n is not 0
	if strLen == 0 && n != 0 {
		return "*"
	}

	// if given n is bigger than string length
	if n >= uint(strLen) {
		return strings.Repeat("*", strLen)
	}

	secondHalf := s[n:]
	secondHalfLen := len(secondHalf)

	// any char but new line
	re := regexp.MustCompile(`[\s\S]*`)
	secondHalf = re.ReplaceAllString(secondHalf, strings.Repeat("*", secondHalfLen))

	resultingStr := s[:n] + secondHalf

	return resultingStr
}

func WordSplit(arr [2]string) string {

	words := arr[1]
	doesWordExist := make(map[string]bool)
	words_ := strings.Split(words, ",")
	testStr := arr[0]
	tempStr := ""   // a string to load and erase found words
	resultStr := "" // a string to collect each word found

	// collect words and assign boolean accordingly
	for _, word := range words_ {
		if !doesWordExist[word] {
			doesWordExist[word] = true
		}
	}

	// iterate over test string to see
	// whether there is a match or not
	for _, letter := range testStr {
		tempStr += string(letter)

		// if there is a match occurred
		if _, ok := doesWordExist[tempStr]; ok {
			resultStr += tempStr + ","
			tempStr = ""
		}
	}

	if len(resultStr) == 0 {
		return "not possible"
	} else if len(resultStr)-2 != len(testStr) {
		return "not possible"
	}

	// to get rid of comma at the end
	return resultStr[:len(resultStr)-1]
}

func VariadicSet(i ...interface{}) []interface{} {
	finalInterface := make([]interface{}, 0)

TestLoop:
	for _, testElement := range i {

		if len(finalInterface) == 0 {
			finalInterface = append(finalInterface, testElement)
			continue
		}

		// checking whether there is any duplicate
		// in the interface that will be returned
		for _, setElement := range finalInterface {
			if setElement == testElement {
				continue TestLoop
			}
		}
		finalInterface = append(finalInterface, testElement)
	}

	return finalInterface
}
