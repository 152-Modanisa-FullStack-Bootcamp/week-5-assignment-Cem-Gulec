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

	// collect char counts from string
	for _, v := range s {
		// if char does not exist before
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
		// append every char multiplied by its charcount
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

	// second half of the string from nth character
	secondHalf := s[n:]
	secondHalfLen := len(secondHalf)

	// using regex to find any character but new line
	re := regexp.MustCompile(`[\s\S]*`)

	// replacing them with * character
	secondHalf = re.ReplaceAllString(secondHalf, strings.Repeat("*", secondHalfLen))

	// appending them to each other for constructing final string
	resultingStr := s[:n] + secondHalf

	return resultingStr
}

func WordSplit(arr [2]string) string {

	words := arr[1]                        // words that will be used for comparison
	doesWordExist := make(map[string]bool) // hashmap to check whether word already collected
	words_ := strings.Split(words, ",")    // splitted version of words string
	testStr := arr[0]                      // given string at the test case
	tempStr := ""                          // a string to load and erase found words
	resultStr := ""                        // a string to collect each word found

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
		// append it to resultStr
		if _, ok := doesWordExist[tempStr]; ok {
			resultStr += tempStr + ","
			tempStr = ""
		}
	}

	// if there is no matching
	if len(resultStr) == 0 {
		return "not possible"
	} else if len(resultStr)-2 != len(testStr) {
		// example case where this may happen:
		// words: cat bat ....
		// string: catbatt
		// it will not be seperated into exactly two but there will be one letter left unmatched
		return "not possible"
	}

	// to get rid of comma at the end
	// in this case we seperated string into exactly two
	return resultStr[:len(resultStr)-1]
}

func VariadicSet(i ...interface{}) []interface{} {
	finalInterface := make([]interface{}, 0)

TestLoop:
	for _, testElement := range i {

		// initially start by appending first element
		if len(finalInterface) == 0 {
			finalInterface = append(finalInterface, testElement)
			continue
		}

		// checking whether there is any duplicate
		// if there is, continue from next element
		for _, setElement := range finalInterface {
			if setElement == testElement {
				continue TestLoop
			}
		}

		// since there is no duplicate of the element, append it
		finalInterface = append(finalInterface, testElement)
	}

	return finalInterface
}
