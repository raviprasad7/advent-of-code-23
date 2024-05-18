package haunted_wasteland

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Name      string
	LeftNode  string
	RightNode string
}

func Run() {
	fmt.Println("Part One - Steps to reach ZZZ:", PartOne())
	fmt.Println("Part Two - Steps to reach all xxZ:", PartTwo())
}

func PartOne() int {
	var (
		instructions string
		nodes        []Node
		nodeMap      map[string]Node = make(map[string]Node)
		steps        int
	)

	file, err := os.Open("./08_haunted_wasteland/input.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}

	scanner := bufio.NewScanner(file)
	instructions = scanner.Text()
	scanner.Text()

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if instructions == "" {
			instructions = line
			continue
		}
		parts := strings.Split(line, "=")
		nodeName := strings.TrimSpace(parts[0])
		nodeDirectionParts := strings.Split(parts[1], ",")
		nodeLeft := nodeDirectionParts[0][2:]
		nodeRight := nodeDirectionParts[1][1:4]

		newNode := Node{
			Name:      nodeName,
			LeftNode:  nodeLeft,
			RightNode: nodeRight,
		}
		nodes = append(nodes, newNode)
		nodeMap[nodeName] = newNode
	}

	ptr := 0
	currNode := "AAA"

	for {
		if currNode == "ZZZ" {
			break
		}
		if ptr == len(instructions) {
			ptr = 0
		}
		currDirection := string(instructions[ptr])
		if currDirection == "R" {
			currNode = nodeMap[currNode].RightNode
		} else if currDirection == "L" {
			currNode = nodeMap[currNode].LeftNode
		}
		ptr++
		steps++
	}

	return steps
}

func PartTwo() int {
	var (
		instructions string
		nodeMap      map[string]Node = make(map[string]Node)
		steps        int
		startNodes   []Node
		nodeCache    map[string]string = make(map[string]string)
	)

	file, err := os.Open("./08_haunted_wasteland/input.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}

	scanner := bufio.NewScanner(file)
	instructions = scanner.Text()
	scanner.Text()

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if instructions == "" {
			instructions = line
			continue
		}
		parts := strings.Split(line, "=")
		nodeName := strings.TrimSpace(parts[0])
		nodeDirectionParts := strings.Split(parts[1], ",")
		nodeLeft := nodeDirectionParts[0][2:]
		nodeRight := nodeDirectionParts[1][1:4]

		newNode := Node{
			Name:      nodeName,
			LeftNode:  nodeLeft,
			RightNode: nodeRight,
		}
		nodeMap[nodeName] = newNode

		if nodeName[2:3] == "A" {
			startNodes = append(startNodes, newNode)
		}
	}
	fmt.Println("Length of nodeMap", len(nodeMap))

	ptr := 0
	currNodes := make([]string, len(startNodes))

	for idx, node := range startNodes {
		currNodes[idx] = node.Name
	}

	for {
		hasAllReachedEnd := true
		// fmt.Println("Start nodes", currNodes, currNodes[0][2:3])
		concatenatedString := ""
		concatenatedEndString := ""
		for _, currNode := range currNodes {
			concatenatedString += currNode
			concatenatedEndString += currNode[2:3]
			if currNode[2:3] != "Z" {
				hasAllReachedEnd = false
				// break
			}
		}
		if nodeCache[concatenatedString] != "" {
			fmt.Println("EXISTSSSSSSS", concatenatedString)
		}
		nodeCache[concatenatedString] = concatenatedEndString
		if hasAllReachedEnd {
			break
		}
		if ptr == len(instructions) {
			ptr = 0
		}
		currDirection := string(instructions[ptr])
		if currDirection == "R" {
			for i := 0; i < len(currNodes); i++ {
				currNodes[i] = nodeMap[currNodes[i]].RightNode
			}
		} else if currDirection == "L" {
			for i := 0; i < len(currNodes); i++ {
				currNodes[i] = nodeMap[currNodes[i]].LeftNode
			}
		}
		ptr++
		steps++
	}

	return steps
}
