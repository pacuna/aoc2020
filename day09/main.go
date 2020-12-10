package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

// true if there's two numbers that add up to the target
func hasSum(si []int, target int) bool {
	d := make(map[int]int)
	for idx, i := range si {
		d[i] = idx
	}
	for _, el := range si {
		if _, ok := d[target-el]; ok {
			return true
		}
	}
	return false

}

func PartI() int {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	si := []int{}

	c := -1
	for scanner.Scan() {
		var n int
		fmt.Sscanf(scanner.Text(), "%d", &n)
		si = append(si, n)
		c++

		// check if we should count
		if c > 25 {
			if !hasSum(si[c-25:c], n) {
				return n
			}
		}
	}
	return -1
}

func PartII() int {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	si := []int{}

	for scanner.Scan() {
		var n int
		fmt.Sscanf(scanner.Text(), "%d", &n)
		si = append(si, n)
	}

	// simple sliding window
	var start, end, sum int
	target := PartI()
	for end < len(si) {
		sum += si[end]
		for sum > target {
			sum -= si[start]
			start++
		}
		if sum == target {
			min, max := minMax(si[start : end+1])
			return min + max
		}
		end++
	}
	return -1
}

func minMax(si []int) (int, int) {
	min := math.MaxInt64
	max := math.MinInt64
	for _, el := range si {
		if el < min {
			min = el
		}
		if el > max {
			max = el
		}
	}
	return min, max
}

func main() {

	fmt.Println(PartI())
	fmt.Println(PartII())
}
