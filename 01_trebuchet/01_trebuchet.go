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
		lines             []string
		sum               int = 0
		sumIncludingWords int = 0
	)

	file, err := os.Open("./01_trebuchet/input.txt")

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	for _, line := range lines {
		firstDigit, firstDigitIncludingWords := findFirstDigit(line)
		lastDigit, lastDigitIncludingWords := findLastDigit(line)

		sum += firstDigit*10 + lastDigit
		sumIncludingWords += firstDigitIncludingWords*10 + lastDigitIncludingWords
	}

	fmt.Println("Part One - Overall sum(including only digits): ", sum)
	fmt.Println("Part Two - Overall sum(including digits as words): ", sumIncludingWords)
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

func findFirstDigit(str string) (int, int) {
	var firstDigit int = 0
	var firstDigitAsWord, firstDigitAsWordIndex = findFirstDigitAsWord(str)

	for idx, ch := range str {
		if unicode.IsDigit(ch) {
			firstDigit = int(ch - '0')
			if idx < firstDigitAsWordIndex {
				firstDigitAsWord = firstDigit
			}
			break
		}
	}

	return firstDigit, firstDigitAsWord
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
			lastDigit = digit
		}
	}

	return lastDigit, lastIndex
}

func findLastDigit(str string) (int, int) {
	var lastDigit int
	var lastDigitAsWord, lastDigitAsWordIndex = findLastDigitAsWord(str)

	for i := len(str) - 1; i >= 0; i-- {
		ch := rune(str[i])
		if unicode.IsDigit(ch) {
			lastDigit = int(ch - '0')
			if i > lastDigitAsWordIndex {
				lastDigitAsWord = lastDigit
			}
			break
		}
	}

	return lastDigit, lastDigitAsWord
}
