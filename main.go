package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/raviprasad7/advent-of-code-23/01_trebuchet"
	"github.com/raviprasad7/advent-of-code-23/02_cube_conundrum"
	"github.com/raviprasad7/advent-of-code-23/03_gear_ratios"
	"github.com/raviprasad7/advent-of-code-23/04_scratchcards"
)

func main() {
	var problemNumber int
	problemMapping := map[int]string{
		1: "Trebuchet",
		2: "Cube Conundrum",
		3: "Gear Ratios",
		4: "Scratchcards",
	}

	args := os.Args

	if len(args) > 1 {
		problemNumber, _ = strconv.Atoi(args[1])
	} else {
		reader := bufio.NewReader(os.Stdin)
		for idx, problem := range problemMapping {
			fmt.Printf("Day %d - %s\n", idx, problem)
		}
		fmt.Print("Which problem to run?[1-4] ")
		input, _ := reader.ReadString('\n')
		problemNumber, _ = strconv.Atoi(strings.TrimSpace(input))
	}

	fmt.Printf("\nRunning the problem: %s\n", problemMapping[problemNumber])

	switch problemNumber {
	case 1:
		trebuchet.Run()
	case 2:
		cube_conundrum.Run()
	case 3:
		gear_ratios.Run()
	case 4:
		scratchcards.Run()
	}
}
