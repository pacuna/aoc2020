package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

func solveI(input io.Reader) (int, int) {
	scanner := bufio.NewScanner(input)
	nums := []int{}
	for scanner.Scan() {
		var n int
		fmt.Sscanf(scanner.Text(), "%d", &n)
		nums = append(nums, n)
	}
	sort.Ints(nums)
	var diff1, diff3 int
	switch nums[0] {
	case 1:
		diff1++
	case 3:
		diff3++
	}
	for i := 1; i < len(nums); i++ {
		switch nums[i] - nums[i-1] {
		case 1:
			diff1++
		case 3:
			diff3++
		}
	}
	return diff1, diff3 + 1
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	d1, d3 := solveI(f)
	fmt.Println(d1 * d3)

}
