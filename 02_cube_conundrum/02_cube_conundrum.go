package cube_conundrum

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var colorLimit = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func Run() {
	var (
		lines        []string
		result       int = 0
		powerSetsSum int = 0
	)

	file, err := os.Open("./02_cube_conundrum/input.txt")
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

	for _, line := range lines {
		parts := strings.Split(line, ":")
		gamePart := strings.Split(parts[0], " ")
		gameId, _ := strconv.Atoi(gamePart[1])
		sets := strings.Split(parts[1], ";")
		isValidGame := true
		colorMaxCount := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		powerSet := 0

		for _, set := range sets {
			colors := strings.Split(set, ",")

			for _, color := range colors {
				colorParts := strings.Split(strings.TrimSpace(color), " ")

				count, _ := strconv.Atoi(colorParts[0])
				colorName := colorParts[1]

				if count > colorLimit[colorName] {
					isValidGame = false
				}

				if count > colorMaxCount[colorName] {
					colorMaxCount[colorName] = count
				}
			}

		}

		if isValidGame {
			result += gameId
		}
		powerSet = colorMaxCount["red"] * colorMaxCount["green"] * colorMaxCount["blue"]

		powerSetsSum += powerSet
	}

	fmt.Println("Part One - Sum of valid game ids:", result)
	fmt.Println("Part Two - Sum of power sets:", powerSetsSum)
}
