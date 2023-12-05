package scratchcards

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

func Run() {
	var (
		lines []string
	)

	file, err := os.Open("./04_scratchcards/input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	totalPointsWon, totalScratchcardsWon := findPointsWon(lines)

	fmt.Println("Part One - Total points won:", totalPointsWon)
	fmt.Println("Part Two - Total scratchcards won:", totalScratchcardsWon)
}

func findPointsWon(cards []string) (int, int) {
	totalPoints := 0
	totalScratchcards := 0
	digitPattern := `\d+`
	digitRegex := regexp.MustCompile(digitPattern)
	scratchcardsWonCache := map[int]int{}

	for idx, card := range cards {
		pointsWon := []string{}

		parts := strings.Split(card, ":")
		points := strings.Split(parts[1], "|")
		winningPoints := digitRegex.FindAllString(points[0], -1)
		pointsGotten := digitRegex.FindAllString(points[1], -1)

		for _, winningPoint := range winningPoints {
			for _, pointGotten := range pointsGotten {
				if winningPoint == pointGotten {
					pointsWon = append(pointsWon, winningPoint)
					break
				}
			}
		}

		cardPoints := len(pointsWon)
		scratchcardsWonCache[idx] = cardPoints

		if len(pointsWon) == 0 {
			continue
		}

		totalPoints += int(math.Pow(2, float64(len(pointsWon)-1)))
	}

	for i := len(cards) - 1; i >= 0; i-- {
		currCardPoints := scratchcardsWonCache[i]
		totalScratchcards++
		currScratchcards := 1

		for j := 1; j <= currCardPoints; j++ {
			totalScratchcards += scratchcardsWonCache[i+j]
			currScratchcards += scratchcardsWonCache[i+j]
		}
		scratchcardsWonCache[i] = currScratchcards
	}

	return totalPoints, totalScratchcards
}
