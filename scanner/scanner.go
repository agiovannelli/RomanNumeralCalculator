package scanner

import (
	"regexp"
	"strings"
)

// Globals
var inputTokens []token
var splitIndexIncr int = 0
var totalTokens int
var currentTokenIndex int = -1

// Token structure
type token struct {
	splitIndex int
	data       string
}

/*
Creates input slice from console args and sets global length value of input array.
*/
func CreateInputSlice(input string) {
	indicesNeedingTrailSpace := regexp.MustCompile(`[/\{+/\(+/\)+/\}+/\[+/\]+]`).FindAllStringIndex(input, -1)
	for i := len(indicesNeedingTrailSpace) - 1; i >= 0; i-- {
		substr := " " + input[indicesNeedingTrailSpace[i][0]:indicesNeedingTrailSpace[i][1]] + " "
		input = input[:indicesNeedingTrailSpace[i][0]] + substr + input[indicesNeedingTrailSpace[i][1]:]
	}

	input = strings.TrimSpace(input)
	splitString := strings.Split(input, " ")

	for i := 0; i < len(splitString); i++ {
		if splitString[i] != "" && splitString[i] != " " {
			inputTokens = append(inputTokens, token{splitIndexIncr, splitString[i]})
			splitIndexIncr++
		}
	}

	totalTokens = len(inputTokens)
}

/*
Sets the global currentToken and nextToken variables.
*/
func GetNextToken() (int, string) {
	currentTokenIndex += 1
	if currentTokenIndex < totalTokens {
		return inputTokens[currentTokenIndex].splitIndex, inputTokens[currentTokenIndex].data
	}
	return -1, "terminal"
}

/*
Creates a map of input token slice data.
*/
func GetInputTokenSlice() map[int]string {
	inputTokenMap := make(map[int]string)

	for i := 0; i < len(inputTokens); i++ {
		inputTokenMap[inputTokens[i].splitIndex] = inputTokens[i].data
	}

	return inputTokenMap
}
