package gear_ratios

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Number struct {
	Value      string
	Row        int
	StartIndex int
	EndIndex   int
}

type Symbol struct {
	Value string
	Row   int
	Index int
}

func Run() {
	var (
		lines      []string
		numberList []Number
		symbolList []Symbol
	)

	file, err := os.Open("./03_gear_ratios/input.txt")

	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	numberList = buildNumberList(lines)
	symbolList = buildSymbolList(lines)

	fmt.Println("Part One - Sum of part numbers:", findPartNumbersSum(numberList, symbolList))
	fmt.Println("Part Two - Sum of gear ratios:", findGearRatiosSum(numberList, symbolList))
}

func buildNumberList(input []string) []Number {
	numberList := []Number{}
	digitPattern := `\d+`
	digitRegex := regexp.MustCompile(digitPattern)

	for idx, line := range input {
		matches := digitRegex.FindAllStringIndex(line, -1)

		for _, match := range matches {
			startIndex := match[0]
			endIndex := match[1]
			matchedNumber := line[startIndex:endIndex]
			numberList = append(numberList, Number{
				Value:      matchedNumber,
				StartIndex: startIndex,
				EndIndex:   endIndex,
				Row:        idx,
			})
		}
	}

	return numberList
}

func buildSymbolList(input []string) []Symbol {
	symbolList := []Symbol{}
	symbolPattern := `[^\d|.]`
	symbolRegex := regexp.MustCompile(symbolPattern)

	for idx, line := range input {
		matches := symbolRegex.FindAllStringIndex(line, -1)

		for _, match := range matches {
			index := match[0]
			matchedSymbol := line[index]
			symbolList = append(symbolList, Symbol{
				Value: string(matchedSymbol),
				Index: index,
				Row:   idx,
			})
		}
	}

	return symbolList
}

func findPartNumbersSum(numberList []Number, symbolList []Symbol) int {
	partNumbersSum := 0

	for _, number := range numberList {
		for _, symbol := range symbolList {
			inAdjacentRow := symbol.Row >= number.Row-1 && symbol.Row <= number.Row+1
			inAdjacentColumn := symbol.Index >= number.StartIndex-1 && symbol.Index <= number.EndIndex

			if inAdjacentRow && inAdjacentColumn {
				numberValue, _ := strconv.Atoi(number.Value)
				partNumbersSum += numberValue
				break
			}
		}
	}

	return partNumbersSum
}

func findGearRatiosSum(numberList []Number, symbolList []Symbol) int {
	gearRatiosSum := 0

	for _, symbol := range symbolList {
		adjacentNumbers := []Number{}

		if symbol.Value != "*" {
			continue
		}

		for _, number := range numberList {
			numberLength := len(number.Value)
			inAdjacentRow := number.Row >= symbol.Row-1 && number.Row <= symbol.Row+1
			inAdjacentColumn := number.StartIndex >= symbol.Index-numberLength && number.EndIndex <= symbol.Index+numberLength+1

			if inAdjacentRow && inAdjacentColumn {
				adjacentNumbers = append(adjacentNumbers, number)
			}
		}

		if len(adjacentNumbers) == 2 {
			firstNumber, _ := strconv.Atoi(adjacentNumbers[0].Value)
			secondNumber, _ := strconv.Atoi(adjacentNumbers[1].Value)
			gearRatio := firstNumber * secondNumber

			gearRatiosSum += gearRatio
		}
	}

	return gearRatiosSum
}
