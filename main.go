package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/raviprasad7/advent-of-code-23/01_trebuchet"
)

func main() {
	packageMapping := map[int]string{
		1: "Trebuchet",
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Which problem to run?[1] ")
	input, _ := reader.ReadString('\n')
	problemNumber, _ := strconv.Atoi(strings.TrimSpace(input))

	fmt.Printf("Running the problem: %s\n", packageMapping[problemNumber])

	switch problemNumber {
	case 1:
		trebuchet.Run()
	}
}
