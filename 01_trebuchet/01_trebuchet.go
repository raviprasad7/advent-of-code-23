package trebuchet

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"unicode"
)

var digits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func Run() {
	var (
		lines []string
		sum   int = 0
	)

	fmt.Println("Enter the input(press Ctrl+D on Unix/Linux or Ctrl+Z on Windows to finish):")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	for _, line := range lines {
		firstDigit := findFirstDigit(line)
		lastDigit := findLastDigit(line)
		fmt.Println(firstDigit, lastDigit)
		sum += firstDigit*10 + lastDigit
	}

	fmt.Println("Overall sum: ", sum)
}

func findFirstDigitAsWord(str string) (int, int) {
	var (
		firstIndex int = math.MaxInt64
		firstDigit int = math.MaxInt64
	)

	for digitAsWord, digit := range digits {
		index := strings.Index(str, digitAsWord)

		if index > -1 && index < firstIndex {
			firstIndex = index
			firstDigit = digit
		}
	}

	return firstDigit, firstIndex
}

func findFirstDigit(str string) int {
	var firstDigitAsWord, firstDigitAsWordIndex = findFirstDigitAsWord(str)

	fmt.Println("output of findFirstDigitAsWord", firstDigitAsWord, firstDigitAsWordIndex)

	for idx, ch := range str {
		if unicode.IsDigit(ch) {
			if idx < firstDigitAsWordIndex {
				return int(ch - '0')
			} else {
				return firstDigitAsWord
			}
		}
	}
	if firstDigitAsWord < 10 {
		return firstDigitAsWord
	}
	return 0
}

func findLastDigitAsWord(str string) (int, int) {
	var (
		lastDigit int = -1
		lastIndex int = -1
	)

	for digitAsWord, digit := range digits {
		index := strings.LastIndex(str, digitAsWord)

		if index > lastIndex {
			lastIndex = index
			fmt.Println("Digit found", digitAsWord, digit)
			lastDigit = digit
		}
	}

	return lastDigit, lastIndex
}

func findLastDigit(str string) int {
	var lastDigitAsWord, lastDigitAsWordIndex = findLastDigitAsWord(str)

	fmt.Println("Output of findLastDigitAsWordIndex", lastDigitAsWord, lastDigitAsWordIndex)

	for i := len(str) - 1; i >= 0; i-- {
		ch := rune(str[i])
		if unicode.IsDigit(ch) {
			if i > lastDigitAsWordIndex {
				return int(ch - '0')
			} else {
				return lastDigitAsWord
			}
		}
	}
	if lastDigitAsWord > 0 {
		return lastDigitAsWord
	}
	return 0
}
