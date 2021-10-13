package logger

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// Global
var positionOfViolation int

/*
Prints issues to console for errors found in part1 query.
*/
func LogErrorToConsole(input string, inputTokenMap map[int]string, tokenIndex int, errorType string) {
	buildInputString(inputTokenMap, tokenIndex)
	fmt.Println(input)
	fmt.Println(createUnderCarriageCarrot(positionOfViolation))
	fmt.Println(createErrorLogMessage(errorType))
}

/*
Prints error message specific to missing right parentheses.
*/
func LogEndOfQueryErrorToConsole(input string, errorType string) {
	fmt.Println(input)
	fmt.Println(createUnderCarriageCarrot(len(input)))
	fmt.Println(createErrorLogMessage(errorType))
}

/*
Creates a mock input string to determine location of error in user input.
*/
func buildInputString(inputTokenMap map[int]string, errorKey int) string {
	result := ""

	keys := make([]int, 0)
	for k, _ := range inputTokenMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		currentSubstring := inputTokenMap[k]
		if k == errorKey {
			positionOfViolation = len(result)
		}

		leftGroupMatch, _ := regexp.MatchString(`[{[(]`, currentSubstring)
		rightGroupMatch, _ := regexp.MatchString(`[)}/\]]`, currentSubstring)

		if leftGroupMatch {
			result += currentSubstring
		} else if rightGroupMatch {
			result = strings.TrimSpace(result)
			result += currentSubstring + " "
		} else {
			result += currentSubstring + " "
		}
	}

	return strings.TrimSpace(result)
}

/*
Creates undercarriage string using carrot position information.
*/
func createUnderCarriageCarrot(numSpaceChars int) string {
	underCarriageCarrotString := ""
	for i := 0; i < numSpaceChars; i++ {
		underCarriageCarrotString += " "
	}
	underCarriageCarrotString += "^"
	return underCarriageCarrotString
}

/*
Creates specific error log message based on error type.
*/
func createErrorLogMessage(errorType string) string {
	specificMessage := ""
	switch errorType {
	case "lexical":
		specificMessage += "You offend Caesar with your sloppy lexical habits!"
	case "syntax":
		specificMessage += "True Romans would not understand your syntax!"
	case "zero":
		specificMessage += "Arab merchants haven't left for India yet!"
	case "negative":
		specificMessage += "Caesar demands positive thoughts!"
	default:
		// None...
	}
	return fmt.Sprintf("Quid dicis? %v", specificMessage)
}
