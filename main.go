package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/raviprasad7/advent-of-code-23/01_trebuchet"
	"github.com/raviprasad7/advent-of-code-23/02_cube_conundrum"
)

func main() {
	var problemNumber int
	packageMapping := map[int]string{
		1: "Trebuchet",
		2: "Cube Conundrum",
	}

	args := os.Args

	if len(args) > 1 {
		problemNumber, _ = strconv.Atoi(args[1])
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Which problem to run?[1-2] ")
		input, _ := reader.ReadString('\n')
		problemNumber, _ = strconv.Atoi(strings.TrimSpace(input))
	}

	fmt.Printf("Running the problem: %s\n", packageMapping[problemNumber])

	switch problemNumber {
	case 1:
		trebuchet.Run()
	case 2:
		cube_conundrum.Run()
	}
}
