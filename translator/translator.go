package translator

import (
	"fmt"
	"math"
	"part1/logger"
	"part1/roman"
	"part1/scanner"
	"regexp"
	"strconv"
)

// Globals
var currentToken token
var nextToken token
var currentSlice []token
var todoMap = make(map[int][]token)
var groupLevel int

// Constants
const ERR string = "error"
const LEX string = "lexical"
const SYX string = "syntax"
const TER string = "terminal"

// Token structure
type token struct {
	splitIndex  int
	data        string
	translation string
	priority    int
	errType     string
}

// Operator maps
var termOperators = map[string]string{
	"plus":  "+",
	"minus": "-"}

var factorOperators = map[string]string{
	"times":  "*",
	"divide": "/"}

var expOperators = map[string]string{
	"power":  "^",
	"modulo": "%"}

/*
Creates input slice using scanner, sets current/next token, starts parsing process.
*/
func Start(input string) {
	scanner.CreateInputSlice(input)

	// Initial setup for current/next token (required on setup).
	currentToken = constructTokenFromScanner()
	nextToken = constructTokenFromScanner()
	groupLevel = 0

	for currentToken.data != TER {
		checkExpression()
		updateTokens()
	}

	logResultToConsole(input)
}

/*
Determines error message to display or prints operation result.
*/
func logResultToConsole(input string) {
	if !errorsInSlice(input) && !errorInGroupLevel(input) {
		result := findHighestPriorityOpAndTokens(currentSlice)
		if !errorsInSlice(input) {
			aToI, _ := strconv.Atoi(result.translation)
			fmt.Println(roman.ConvertToRomanNumeral(aToI))
		}
	}
}

/*
Determines if errors exist in currentSlice tokens. If they do, adds to errorSlice.
*/
func errorsInSlice(input string) bool {
	result := false
	for index := 0; index < len(currentSlice); index++ {
		if currentSlice[index].translation == ERR {
			er := currentSlice[index]
			logger.LogErrorToConsole(input, scanner.GetInputTokenSlice(), er.splitIndex, er.errType)
			result = true
			break
		}
	}
	return result
}

/*
Adds syntax error to end of part1 query if group level is not zero.
*/
func errorInGroupLevel(input string) bool {
	result := false
	if groupLevel != 0 {
		result = true
		logger.LogEndOfQueryErrorToConsole(input, SYX)
	}
	return result
}

/*
Sets current token to next, gets next token from scanner.
*/
func updateTokens() {
	currentToken = nextToken
	nextToken = constructTokenFromScanner()
}

/*
Operations for slices/maps based on grouping level increment.
*/
func incrGroupLevel() {
	addToTodoMapSlice()
	resetCurrentSlice()
	groupLevel += 1
}

/*
Operations for slices/maps based on grouping level decrement.
*/
func decrGroupLevel() {
	resultingToken := findHighestPriorityOpAndTokens(currentSlice)
	resetCurrentSlice()
	todoMap[groupLevel] = currentSlice

	groupLevel -= 1
	setCurrentSlice()
	currentSlice = append(currentSlice, resultingToken)
}

/*
Appends current token to current slice.
*/
func updateCurrentSlice() {
	currentSlice = append(currentSlice, currentToken)
}

/*
Resets current slice by setting to 'nil'.
*/
func resetCurrentSlice() {
	currentSlice = nil
}

/*
Sets current slice to todo map slice at the group level key.
*/
func setCurrentSlice() {
	currentSlice = todoMap[groupLevel]
}

/*
Set current slice in todoMap at the group level key.
*/
func addToTodoMapSlice() {
	if todoMap[groupLevel] == nil {
		todoMap[groupLevel] = currentSlice
	}
}

/*
Creates token from scanner data.
*/
func constructTokenFromScanner() token {
	si, d := scanner.GetNextToken()
	return token{si, d, "tbd", -1, "none"}
}

/*
Checks base expressions of grammar tree, branches according to current token being left grouping or Roman numeral.
*/
func checkExpression() {
	// L from grammar
	if checkLeftGrouping() {
		incrGroupLevel()
	} else if checkRomanNumeral() {
		romanToNum := roman.ConvertToInteger(currentToken.data)
		currentToken.translation = strconv.Itoa(romanToNum)
		updateCurrentSlice()
		updateTokens()
		if !checkTerm() {
			checkRightGrouping()
		}
	} else {
		currentToken.translation = ERR
		currentToken.errType = LEX
		updateCurrentSlice()
	}
}

/*
Determine if token is left grouping character.
*/
func checkLeftGrouping() bool {
	isLeftGrouping, _ := regexp.MatchString("[{([]", currentToken.data)
	return isLeftGrouping
}

/*
Determine if token is valid Roman numeral.
*/
func checkRomanNumeral() bool {
	isRomanNumeral, _ := regexp.MatchString("^M{0,4}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$", currentToken.data)
	return isRomanNumeral
}

/*
Determine if token is right grouping character.
*/
func checkRightGrouping() {
	isRightGrouping, _ := regexp.MatchString(`[/\]/\)/\}]`, currentToken.data)
	if isRightGrouping {
		decrGroupLevel()
		updateTokens()
		checkRightGrouping()
	} else if checkTerm() {
	} else if currentToken.data != TER {
		currentToken.translation = ERR
		currentToken.errType = LEX
		updateCurrentSlice()
	}
}

/*
Determine if token is term operator.
*/
func checkTerm() bool {
	if val, ok := termOperators[currentToken.data]; ok {
		currentToken.translation = val
		currentToken.priority = 1
		updateCurrentSlice()
		return true
	} else {
		return checkFactor()
	}
}

/*
Determine if token is factor operator.
*/
func checkFactor() bool {
	if val, ok := factorOperators[currentToken.data]; ok {
		currentToken.translation = val
		currentToken.priority = 2
		updateCurrentSlice()
		return true
	} else {
		return checkExpo()
	}
}

/*
Determine if token is exponential or modulo operator.
*/
func checkExpo() bool {
	if val, ok := expOperators[currentToken.data]; ok {
		currentToken.translation = val
		currentToken.priority = 3
		updateCurrentSlice()
		return true
	}
	return false
}

/*
Executes grouping ops per priority.
*/
func findHighestPriorityOpAndTokens(group []token) token {
	for len(group) > 2 {
		highestPriority := -1
		highestPriorityIndex := -1

		for index := 0; index < len(group); index++ {
			if group[index].priority > highestPriority || group[index].priority == 3 {
				highestPriority = group[index].priority
				highestPriorityIndex = index
			}
		}

		leftToken := group[highestPriorityIndex-1]
		rightToken := group[highestPriorityIndex+1]
		newToken := executeOps(group[highestPriorityIndex], leftToken, rightToken)
		group = updateSlice(group, highestPriorityIndex, newToken)
	}
	return group[0]
}

/*
Updates the slice for simplified token.
*/
func updateSlice(slice []token, index int, newSliceToken token) []token {
	updatedSlice := slice[:index-1]
	updatedSlice = append(updatedSlice, newSliceToken)
	if len(slice) > index+2 {
		rightSlice := slice[index+2:]
		updatedSlice = append(updatedSlice, rightSlice...)
	}

	return updatedSlice
}

/*
Used in todo token iteration to perform part1 operations against translated values.
*/
func executeOps(op token, leftToken token, rightToken token) token {
	result := 0
	leftInt, _ := strconv.Atoi(leftToken.translation)
	rightInt, _ := strconv.Atoi(rightToken.translation)

	switch op.translation {
	case "+":
		result = leftInt + rightInt
	case "-":
		result = leftInt - rightInt
	case "*":
		result = leftInt * rightInt
	case "/":
		result = leftInt / rightInt
	case "^":
		result = int(math.Pow(float64(leftInt), float64(rightInt)))
	case "%":
		result = int(math.Mod(float64(leftInt), float64(rightInt)))
	default:
		message := fmt.Sprintf("Failed to execute op (%v) for %v, %v", op.translation, leftToken, rightToken)
		fmt.Println(message)
	}

	return determineOpResultToken(op, result)
}

/*
Creates resulting ops token for successful roman numeral operations and failures.
*/
func determineOpResultToken(op token, result int) token {
	var resultToken token

	if result < 0 {
		resultToken = token{op.splitIndex, op.data, ERR, -1, "negative"}
	} else if result == 0 {
		resultToken = token{op.splitIndex, op.data, ERR, -1, "zero"}
	} else {
		resultToken = token{op.splitIndex, "done", strconv.Itoa(result), -1, "none"}
	}

	return resultToken
}
