package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func decode(input string) (row, col int) {
	var rowL int = 0
	var rowH int = 127
	for _, l := range input[:7] {
		if l == 'F' {
			rowH -= int(math.Ceil((float64(rowH - rowL)) / float64(2)))
		} else if l == 'B' {
			rowL += int(math.Ceil((float64(rowH - rowL)) / float64(2)))
		} else {
			break
		}
	}
	row = rowL
	var colL int = 0
	var colH int = 7
	for _, l := range input[7:] {
		if l == 'L' {
			colH -= int(math.Ceil((float64(colH - colL)) / 2.0))
		} else if l == 'R' {
			colL += int(math.Ceil((float64(colH - colL)) / 2.0))
		} else {
			break
		}
	}
	col = colL
	return row, col
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Part I
	maxSeatId := 0
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		row, col := decode(line)
		maxSeatId = max(maxSeatId, row*8+col)
	}
	fmt.Println(maxSeatId)
	f.Close()

	// Part 2
	seats := make(map[int]bool)
	for i := 0; i <= maxSeatId; i++ {
		seats[i] = true
	}
	f, err = os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		row, col := decode(line)
		if _, ok := seats[row*8+col]; ok {
			delete(seats, row*8+col)
		}
	}
	f.Close()

	ss := []int{}
	for v, _ := range seats {
		ss = append(ss, v)
	}
	sort.Ints(ss)
	fmt.Println(ss)
}
