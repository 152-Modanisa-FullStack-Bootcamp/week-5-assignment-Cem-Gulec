package assignment

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUint32(t *testing.T) {
	/*
		Sum uint32 numbers, return uint32 sum value and boolean overflow flag
		cases need to pass:
			math.MaxUint32, 1 => 0, true
			1, 1 => 2, false
			42, 2701 => 2743, false
			42, math.MaxUint32 => 41, true
			4294967290, 5 => 4294967295, false
			4294967290, 6 => 0, true
			4294967290, 10 => 4, true
	*/
	testCases := []struct {
		x        uint32
		y        uint32
		want     uint32
		overflow bool
	}{
		{
			math.MaxUint32, 1,
			0, true,
		},
		{
			1, 1,
			2, false,
		},
		{
			42, 2701,
			2743, false,
		},
		{
			42, math.MaxUint32,
			41, true,
		},
		{
			4294967290, 5,
			4294967295, false,
		},
		{
			4294967290, 6,
			0, true,
		},
		{
			4294967290, 10,
			4, true,
		},
	}

	for _, testCase := range testCases {
		sum, overflow := AddUint32(testCase.x, testCase.y)
		assert.Equal(t, testCase.want, sum)
		assert.Equal(t, testCase.overflow, overflow)
	}
}

func TestCeilNumber(t *testing.T) {

	/*	Ceil the number within 0.25
		cases need to pass:
			42.42 => 42.50
			42 => 42
			42.01 => 42.25
			42.24 => 42.25
			42.25 => 42.25
			42.26 => 42.50
			42.55 => 42.75
			42.75 => 42.75
			42.76 => 43
			42.99 => 43
			43.13 => 43.25*/

	testCases := []struct {
		value float64
		want  float64
	}{
		{
			42.42,
			42.50,
		},
		{
			42,
			42,
		},
		{
			42.01,
			42.25,
		},
		{
			42.24,
			42.25,
		},
		{
			42.25,
			42.25,
		},
		{
			42.26,
			42.50,
		},
		{
			42.55,
			42.75,
		},
		{
			42.75,
			42.75,
		},
		{
			42.76,
			43,
		},
		{
			42.99,
			43,
		},
		{
			43.13,
			43.25,
		},
	}

	for _, testCase := range testCases {
		point := CeilNumber(testCase.value)
		assert.Equal(t, testCase.want, point)
	}
}

func TestAlphabetSoup(t *testing.T) {

	/*String with the letters in alphabetical order.
	cases need to pass:
	 	"hello" => "ehllo"
		"" => ""
		"h" => "h"
		"ab" => "ab"
		"ba" => "ab"
		"bac" => "abc"
		"cba" => "abc"*/

	testCases := []struct {
		text string
		want string
	}{
		{
			"hello",
			"ehllo",
		},
		{
			"",
			"",
		},
		{
			"h",
			"h",
		},
		{
			"ab",
			"ab",
		},
		{
			"ba",
			"ab",
		},
		{
			"bac",
			"abc",
		},
		{
			"cba",
			"abc",
		},
	}

	for _, testCase := range testCases {
		result := AlphabetSoup(testCase.text)
		assert.Equal(t, testCase.want, result)
	}
}

func TestStringMask(t *testing.T) {

	/*Replace after n(uint) character of string with '*' character.
	cases need to pass:
		"!mysecret*", 2 => "!m********"
		"", n(any positive number) => "*"
		"a", 1 => "*"
		"string", 0 => "******"
		"string", 3 => "str***"
		"string", 5 => "strin*"
		"string", 6 => "******"
		"string", 7(bigger than len of "string") => "******"
		"s*r*n*", 3 => "s*r***" */

	testCases := []struct {
		text string
		num  uint
		want string
	}{
		{
			"!mysecret*", 2,
			"!m********",
		},
		{
			"", 5,
			"*",
		},
		{
			"a", 1,
			"*",
		},
		{
			"string", 0,
			"******",
		},
		{
			"string", 3,
			"str***",
		},
		{
			"string", 5,
			"strin*",
		},
		{
			"string", 6,
			"******",
		},
		{
			"string", 7,
			"******",
		},
		{
			"s*r*n*", 3,
			"s*r***",
		},
	}

	for _, testCase := range testCases {
		result := StringMask(testCase.text, testCase.num)
		assert.Equal(t, testCase.want, result)
	}
}

func TestWordSplit(t *testing.T) {
	/*Your goal is to determine if the first element in the array can be split into two words,
	where both words exist in the dictionary(words variable) that is provided in the second element of array.

	cases need to pass:
		[2]string{"hellocat",words} => hello,cat
		[2]string{"catbat",words} => cat,bat
		[2]string{"yellowapple",words} => yellow,apple
		[2]string{"",words} => not possible
		[2]string{"notcat",words} => not possible
		[2]string{"bootcamprocks!",words} => not possible*/

	words := "apple,bat,cat,goodbye,hello,yellow,why"

	testCases := []struct {
		test [2]string
		want string
	}{
		{
			[2]string{"hellocat", words},
			"hello,cat",
		},
		{
			[2]string{"catbat", words},
			"cat,bat",
		},
		{
			[2]string{"yellowapple", words},
			"yellow,apple",
		},
		{
			[2]string{"", words},
			"not possible",
		},
		{
			[2]string{"notcat", words},
			"not possible",
		},
		{
			[2]string{"bootcamprocks!", words},
			"not possible",
		},
	}

	for _, testCase := range testCases {
		result := WordSplit(testCase.test)
		assert.Equal(t, testCase.want, result)
	}
}

func TestVariadicSet(t *testing.T) {

	/*FINAL BOSS ALERT :)
	Tip: Learn and apply golang variadic functions(search engine -> "golang variadic function" -> WOW You can really dance! )

	Convert inputs to set(no duplicate element)
	cases need to pass:
		4,2,5,4,2,4 => []interface{4,2,5}
		"bootcamp","rocks!","really","rocks! => []interface{"bootcamp","rocks!","really"}
		1,uint32(1),"first",2,uint32(2),"second",1,uint32(2),"first" => []interface{1,uint32(1),"first",2,uint32(2),"second"}
	*/

	testCases := []struct {
		test []interface{}
		want []interface{}
	}{
		{
			[]interface{}{4, 2, 5, 4, 2, 4},
			[]interface{}{4, 2, 5},
		},
		{
			[]interface{}{"bootcamp", "rocks!", "really", "rocks!"},
			[]interface{}{"bootcamp", "rocks!", "really"},
		},
		{
			[]interface{}{1, uint32(1), "first", 2, uint32(2), "second", 1, uint32(2), "first"},
			[]interface{}{1, uint32(1), "first", 2, uint32(2), "second"},
		},
	}

	for _, testCase := range testCases {
		set := VariadicSet(testCase.test...)
		assert.Equal(t, testCase.want, set)
	}
}
