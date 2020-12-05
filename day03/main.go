package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Part I
	fmt.Println(countTrees("input.txt", 3, 1))

	// Part II
	r1d1 := countTrees("input.txt", 1, 1)
	r3d1 := countTrees("input.txt", 3, 1)
	r5d1 := countTrees("input.txt", 5, 1)
	r7d1 := countTrees("input.txt", 7, 1)
	r1d2 := countTrees("input.txt", 1, 2)
	fmt.Println(r1d1 * r3d1 * r5d1 * r7d1 * r1d2)

}

func countTrees(filename string, right, down int) int {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	var i, res, lineLen int
	firstLine := true
	var line string
	for {

		// skip first line and get the line length
		if firstLine {
			scanner.Scan()
			lineLen = len(scanner.Text())
			firstLine = false
		}

		// move right
		i = (i + right) % lineLen

		// move down, return when scan stops
		for j := 0; j < down; j++ {
			if b := scanner.Scan(); !b {
				return res
			}
			line = scanner.Text()
		}
		if string(line[i]) == "#" {
			res++
		}

	}
	return res
}
