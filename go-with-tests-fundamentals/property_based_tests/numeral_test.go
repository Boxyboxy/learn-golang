package numeral

import (
	"fmt"
	"log"
	"testing"
	"testing/quick"
)

// table driven tests
var cases = []struct {
	Arabic uint16
	Roman  string
}{
	{Arabic: 1, Roman: "I"},
	{Arabic: 2, Roman: "II"},
	{Arabic: 3, Roman: "III"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 6, Roman: "VI"},
	{Arabic: 7, Roman: "VII"},
	{Arabic: 8, Roman: "VIII"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 14, Roman: "XIV"},
	{Arabic: 18, Roman: "XVIII"},
	{Arabic: 20, Roman: "XX"},
	{Arabic: 39, Roman: "XXXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 47, Roman: "XLVII"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 798, Roman: "DCCXCVIII"},
}

func TestRomanNumerals(t *testing.T) {

	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Roman {
				t.Errorf("got %q, want %q", got, test.Roman)
			}
		})
	}

}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases[:1] {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {

			// can't do negative numbers with Roman Numerals
			// int has a much higher maximum value thatn 3999. int is not a great type, we need to try something appropriate

			return true
		}
		log.Println("testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}
	// a function that will run against a number of random inputs
	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}

/*
In this code, the property-based test is `TestPropertiesOfConversion`. Let's break down why this is a property-based test:

1. Use of `quick.Check`:
   The function uses `quick.Check` from the `testing/quick` package, which is Go's built-in tool for property-based testing.

2. Property definition:
   The `assertion` function defines a property that should hold true for all valid inputs:
   ```go
   assertion := func(arabic uint16) bool {
       if arabic > 3999 {
           return true
       }
       roman := ConvertToRoman(arabic)
       fromRoman := ConvertToArabic(roman)
       return fromRoman == arabic
   }
   ```
   This property states that converting an Arabic number to Roman and back should result in the original number.

3. Random input generation:
   `quick.Check` automatically generates random inputs (in this case, `uint16` values) to test the property.

4. Multiple test cases:
   The test runs the property check multiple times (up to 1000 times, as specified in the config) with different random inputs.

5. Failure reporting:
   If the property fails for any input, `quick.Check` will return an error, which is then reported by the test.

6. Input constraints:
   The test handles the constraint that Roman numerals typically don't go above 3999 by simply returning true for larger values.

This test is checking a fundamental property of the conversion functions: that they are inverse operations of each other. It's testing this property across a wide range of inputs, potentially uncovering edge cases or inconsistencies that might be missed in traditional unit tests.

The other test functions (`TestRomanNumerals` and `TestConvertingToArabic`) are examples of table-driven tests, which are different from property-based tests. They check specific input-output pairs rather than general properties of the functions.
*/
