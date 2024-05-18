package wait_for_it

import (
	"bufio"
	"fmt"
	"os"
)

func Run() {
	var (
		lines []string
	)

	file, err := os.Open("./06_wait_for_it/input.txt")
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

	fmt.Println("Input", lines)
}
