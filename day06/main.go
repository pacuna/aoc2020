package main

import (
	"bufio"
	"fmt"
	"os"
)

func solveI(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return -1, err
	}
	scanner := bufio.NewScanner(f)
	var line string
	curGroup := make(map[rune]int)
	var res int
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			// process line
			for _, c := range line {
				curGroup[c]++
			}
			res += len(curGroup)

			// reset
			line = ""
			curGroup = make(map[rune]int)
		} else {
			line += scanner.Text()
		}
	}

	// process last line
	for _, c := range line {
		curGroup[c]++
	}
	res += len(curGroup)

	return res, nil
}

func solveII(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return -1, err
	}
	scanner := bufio.NewScanner(f)
	var line string
	curGroup := make(map[rune]int)
	var res int
	var numPeople int
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			// process line
			for _, c := range line {
				curGroup[c]++
			}

			for _, v := range curGroup {
				if v == numPeople {
					res++
				}
			}

			// reset
			line = ""
			curGroup = make(map[rune]int)
			numPeople = 0
		} else {
			line += scanner.Text()
			numPeople++
		}
	}

	// process last line
	for _, c := range line {
		curGroup[c]++
	}
	for _, v := range curGroup {
		if v == numPeople {
			res++
		}
	}

	return res, nil
}
func main() {
	fmt.Println(solveI("input.txt"))
	fmt.Println(solveII("input.txt"))
}
