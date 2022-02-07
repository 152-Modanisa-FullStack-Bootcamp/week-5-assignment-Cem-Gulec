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
	return ""
}

func VariadicSet(i ...interface{}) []interface{} {
	return nil
}
