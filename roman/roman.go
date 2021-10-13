/*
This module is modeled after https://github.com/chonla/roman-number-go,
but has been modified to be more readable and not require a constructor.
Instead, it acts more 'statically'.
*/
package roman

// Roman numeral/number mappings
var romanNumeralToNumberMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var numberToRomanNumeralMap = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

var descendingIntegersSlice = []int{
	1000,
	900,
	500,
	400,
	100,
	90,
	50,
	40,
	10,
	9,
	5,
	4,
	1,
}

/*
Convert roman numeral to integer.
*/
func ConvertToInteger(rn string) int {
	result := 0
	lengthOfRomanNumeral := len(rn)
	for i := 0; i < lengthOfRomanNumeral; i++ {
		rnCharacter := string(rn[i])
		rnValue := romanNumeralToNumberMap[rnCharacter]
		if i < lengthOfRomanNumeral-1 {
			rnNextCharacter := string(rn[i+1])
			rnNextValue := romanNumeralToNumberMap[rnNextCharacter]
			if rnValue < rnNextValue {
				result += rnNextValue - rnValue
				i++
			} else {
				result += rnValue
			}
		} else {
			result += rnValue
		}
	}
	return result
}

/*
Convert integer to roman numeral.
*/
func ConvertToRomanNumeral(anInt int) string {
	result := ""
	for anInt > 0 {
		intReductionAmount := largestInteger(anInt)
		result += numberToRomanNumeralMap[intReductionAmount]
		anInt -= intReductionAmount
	}
	return result
}

/*
Determines largest integer in current int to return Roman numeral int reduce amount.
*/
func largestInteger(currentInt int) int {
	for _, value := range descendingIntegersSlice {
		if value <= currentInt {
			return value
		}
	}
	return 1
}
