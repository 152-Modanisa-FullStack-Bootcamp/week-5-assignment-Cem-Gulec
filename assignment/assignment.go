package assignment

import (
	"math"
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
	return ""
}

func StringMask(s string, n uint) string {
	return ""
}

func WordSplit(arr [2]string) string {
	return ""
}

func VariadicSet(i ...interface{}) []interface{} {
	return nil
}
